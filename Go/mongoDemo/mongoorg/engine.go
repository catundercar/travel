package mongoorg

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Engine struct {
	Host           string
	Port           string
	ConnectTimeout time.Duration

	DbName string
	db     *mongo.Database
}

func (e *Engine) Init() error {
	c, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", e.Host, e.Port)))
	if err != nil {
		return err
	}

	return nil
}

func (e *Engine) getDB() (*mongo.Database, error) {
	// TODO here can set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.session.Client().Connect(ctx); err != nil {
		return nil, err
	}
	return e.session.Client().Database(e.DbName), nil
}

func (e *Engine) InsertOne(collection string, document interface{}) error {
	db, err := e.getDB()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = db.Collection(collection).InsertOne(ctx, document)
	if err != nil {
		return err
	}
	return nil
}

func (e *Engine) InsertMany(collection string, documents []interface{}) error {
	db, err := e.getDB()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = db.Collection(collection).InsertMany(ctx, documents)
	if err != nil {
		return err
	}

	return nil
}
