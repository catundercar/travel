package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.study.org/mongoDemo/mongo_org"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	c, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
		return
	}

	collection := c.Database("mongo_org").Collection("tests")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	c.Connect(ctx)
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		log.Fatalf("insert error: %s", err.Error())
	}

	fmt.Println(res.InsertedID)

	s := &mongo_org.Student{
		Name: "Jack",
		Age:  18,
		Sex:  "man",
	}

	err = s.Add(c)
	if err != nil {
		fmt.Println("xxx")
		log.Fatal(err)
		return
	}
}
