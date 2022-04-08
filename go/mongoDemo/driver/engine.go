package driver

type Engine interface {
	InsertOne(collection string, document interface{}) error
	InsertMany(collection string, documents []interface{}) error
}
