// +build !mgo

package main

import "go.mongodb.org/mongo-driver/bson"

var (
	bsonRegistry = bson.DefaultRegistry
)

func marshal(doc bson.D) ([]byte, error) {
	return bson.MarshalWithRegistry(bsonRegistry, doc)
}

func unmarshal(docBytes []byte) (bson.D, error) {
	var newDoc bson.D
	err := bson.UnmarshalWithRegistry(bsonRegistry, docBytes, &newDoc)
	return newDoc, err
}
