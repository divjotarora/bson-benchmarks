// +build mgo

package main

import "gopkg.in/mgo.v2/bson"

func marshal(doc bson.D) ([]byte, error) {
	return bson.Marshal(doc)
}

func unmarshal(docBytes []byte) (bson.D, error) {
	var newDoc bson.D
	err := bson.Unmarshal(docBytes, &newDoc)
	return newDoc, err
}
