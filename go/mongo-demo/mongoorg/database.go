package mongoorg

import "go.mongodb.org/mongo-driver/mongo"

// Database is an agent of mongo.Database.
type Database struct {
	mdb *mongo.Database
}

// C returns a Collection base on mongo.Collection.
func (db *Database) C(collection string) *Collection {
	c := db.mdb.Collection(collection)
	return &Collection{
		mcoll: c,
	}
}
