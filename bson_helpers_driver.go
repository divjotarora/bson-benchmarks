// +build !mgo

package main

import (
	"go.mongodb.org/mongo-driver/bson"
)

// Helpers to build bson.D instances using the driver's BSON library.

func buildDoc(n int) bson.D {
	doc := bson.D{}
	for i := 0; i < n; i++ {
		key, val := getStringElement(i)
		doc = append(doc, bson.E{key, val})
	}
	return doc
}

func buildNestedDoc(n int) bson.D {
	doc := bson.D{}
	for i := 0; i < n; i++ {
		doc = bson.D{{"x", doc}}
	}
	return doc
}
