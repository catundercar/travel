package main

import (
	"fmt"
	"log"

	"go.study.org/mongoDemo/mongoorg"
	"go.study.org/mongoDemo/student"
)

func main() {
	//	c, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	//	if err != nil {
	//		log.Fatal(err)
	//		return
	//	}
	//
	//	collection := c.Database("mongo_org").Collection("tests")
	//	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//	c.Connect(ctx)
	//	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	//	if err != nil {
	//		log.Fatalf("insert error: %s", err.Error())
	//	}
	//
	//	fmt.Println(res.InsertedID)

	s := &student.Student{
		Name: "Jack",
		Age:  18,
		Sex:  "man",
	}

	e := &mongoorg.Engine{
		Host:   "127.0.0.1",
		Port:   "27017",
		DbName: "mongoorg",
	}
	if err := e.Init(); err != nil {
		log.Fatalf("Init error: %s", err.Error())
		return
	}

	err := s.Add(e)
	if err != nil {
		fmt.Println("xxx")
		log.Fatal(err)
		return
	}
}
