package mongoorg

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Engine contains a handle of mongodb and configures.
type Engine struct {
	Conf    *Conf
	c       *mongo.Client
	session *session
}

type session struct {
	db *mongo.Database
}

// Database is an agent of mongo.Client.
type Database struct {
}

// Conf contains options to configure a Engine instance.
type Conf struct {
	Host           string
	Port           string
	DbName         string
	ConnectTimeout time.Duration
	SocketTimeout  time.Duration
	MaxPoolSize    uint64
	MinPoolSize    uint64
	Direct         bool
}

// EngineOpt 可选方法
type EngineOpt func(*Engine)

// WithConf 指定配置
func WithConf(conf *Conf) EngineOpt {
	return func(e *Engine) {
		e.Conf = conf
	}
}

// DefaultTimeout 默认超时时间
const DefaultTimeout = 10 * time.Second

// Init initializes the mongodb client.
func (e *Engine) Init() error {
	opts := conf2options(e.Conf)
	c, err := mongo.Connect(NewCtx(40*time.Second),
		opts.ApplyURI(fmt.Sprintf("mongodb://%s:%s/%s", e.Conf.Host, e.Conf.Port, e.Conf.DbName)).SetConnectTimeout(time.Second*10))
	if err != nil {
		return err
	}

	if err := c.Ping(NewCtx(40*time.Second), readpref.Primary()); err != nil {
		return err
	}

	e.c = c
	e.session = &session{
		db: c.Database(e.Conf.DbName),
	}
	return nil
}

func conf2options(conf *Conf) *options.ClientOptions {
	return &options.ClientOptions{
		ConnectTimeout: &conf.ConnectTimeout,
		SocketTimeout:  &conf.SocketTimeout,
		MaxPoolSize:    &conf.MaxPoolSize,
		MinPoolSize:    &conf.MinPoolSize,
		Direct:         &conf.Direct,
	}
}

// NewEngine returns a engine instance.
func NewEngine(opts ...EngineOpt) *Engine {
	e := &Engine{
		Conf: &Conf{
			Host:           "127.0.0.1",
			Port:           "27017",
			DbName:         "mongoorg",
			ConnectTimeout: DefaultTimeout,
			SocketTimeout:  DefaultTimeout,
			MaxPoolSize:    10,
			MinPoolSize:    1,
			Direct:         true,
		},
	}

	for _, fn := range opts {
		fn(e)
	}

	return e
}

// DefaultCtx returns a default timeout ctx, it's 10 second.
func DefaultCtx() context.Context {
	return NewCtx(DefaultTimeout)
}

// NewCtx returns a ctx with a custom timeout ctx.
func NewCtx(timeout time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), timeout)

	return ctx
}

func (e *Engine) getDB() *mongo.Database {
	return e.session.db
}

// Insert insert a document.
func (e *Engine) Insert(collection string, document interface{}) error {
	db := e.getDB()

	_, err := db.Collection(collection).InsertOne(NewCtx(e.Conf.SocketTimeout), document)
	if err != nil {
		return err
	}

	return nil
}

// InsertMany insert many documents.
func (e *Engine) InsertMany(collection string, documents []interface{}) error {
	db := e.getDB()

	_, err := db.Collection(collection).InsertMany(NewCtx(e.Conf.SocketTimeout), documents)
	if err != nil {
		return err
	}

	return nil
}

// Upsert find one document and replace it.
func (e *Engine) Upsert(collection string, filter interface{}, document interface{}) error {
	db := e.getDB()
	upsert := true
	result := db.Collection(collection).FindOneAndReplace(NewCtx(e.Conf.SocketTimeout),
		filter, document, &options.FindOneAndReplaceOptions{
			Upsert: &upsert,
		})
	if result.Err() != nil {
		return result.Err()
	}

	return nil
}
