// +build !mgo

package main

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var (
	tD           = reflect.TypeOf(primitive.D{})
	bsonRegistry = bson.NewRegistryBuilder().
			RegisterCodec(tD, bsonx.ReflectionFreeDCodec).
			Build()
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
