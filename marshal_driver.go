// +build !mgo

package main

import (
	"go.mongodb.org/mongo-driver/bson"
)

var (
	bsonRegistry = bson.DefaultRegistry
)

// Helpers to wrap the driver's bson.Marshal and bson.Unmarshal functions.

func marshal(doc interface{}) ([]byte, error) {
	return bson.MarshalWithRegistry(bsonRegistry, doc)
}

func unmarshal(docBytes []byte) (bson.D, error) {
	var newDoc bson.D
	err := bson.UnmarshalWithRegistry(bsonRegistry, docBytes, &newDoc)
	return newDoc, err
}
