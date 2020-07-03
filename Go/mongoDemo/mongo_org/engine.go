package mongo_org

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Student struct {
	Name string
	Age  int
	Sex  string
}

func New() (*mongo.Client, error) {
	return mongo.NewClient()
}

func (s *Student) Add(c *mongo.Client) error {
	collection := c.Database("mongo_org").Collection("students")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, s)
	if err != nil {
		return err
	}
	fmt.Println(res.InsertedID)
	return nil
}
