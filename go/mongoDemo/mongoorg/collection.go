package mongoorg

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection is an agent of mongo.Collection.
type Collection struct {
	mcoll *mongo.Collection
}

var upsert = true

// Insert insert a document.
func (c *Collection) Insert(document interface{}) (*mongo.InsertOneResult, error) {
	return c.mcoll.InsertOne(context.Background(), document)
}

// InsertMany executes an insert command to insert multiple documents into the collection.
func (c *Collection) InsertMany(documents []interface{}) (*mongo.InsertManyResult, error) {
	return c.mcoll.InsertMany(context.TODO(), documents)
}

// Upsert a new document will be inserted if the filter does not match any documents in the collection.
func (c *Collection) Upsert(filter, update interface{}) (*mongo.UpdateResult, error) {
	return c.mcoll.UpdateOne(context.Background(), update, options.Update().SetUpsert(true))
}

// BulkUpsert base on mongo.Colltion.BulkWrite.
func (c *Collection) BulkUpsert(filters, updates []interface{}) (*mongo.BulkWriteResult, error) {
	opts := options.BulkWrite().SetOrdered(false)
	return c.mcoll.BulkWrite(context.TODO(), buildBulkUpsert(filters, updates), opts)
}

func buildBulkUpsert(filter, update []interface{}) []mongo.WriteModel {
	wms := make([]mongo.WriteModel, 0, len(filter))
	for i := 0; i < len(filter); i++ {
		wms = append(wms, mongo.NewUpdateOneModel().SetFilter(filter[i]).SetUpdate(update).SetUpsert(true))
	}

	return wms
}
