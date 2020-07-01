// +build mgo

package main

import (
	"gopkg.in/mgo.v2/bson"
)

// Helpers to wrap mgo's bson.Marshal and bson.Unmarshal functions.

func marshal(doc interface{}) ([]byte, error) {
	return bson.Marshal(doc)
}

func unmarshal(docBytes []byte) (bson.D, error) {
	var newDoc bson.D
	err := bson.Unmarshal(docBytes, &newDoc)
	return newDoc, err
}
