package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

// Helpers to build bsoncore.Document instances. This file is not build-flagged because it can be used when testing both
// the driver and mgo.

func buildRawDoc(n int) bsoncore.Document {
	idx, doc := bsoncore.AppendDocumentStart(nil)
	for i := 0; i < n; i++ {
		key, val := getStringElement(i)
		doc = bsoncore.AppendStringElement(doc, key, val)
	}

	doc, _ = bsoncore.AppendDocumentEnd(doc, idx)
	return doc
}

func buildNestedRawDoc(n int) bsoncore.Document {
	idx, doc := bsoncore.AppendDocumentStart(nil)
	key, val := getStringElement(n)
	doc = bsoncore.AppendStringElement(doc, key, val)
	doc, _ = bsoncore.AppendDocumentEnd(doc, idx)

	for i := 0; i < n; i++ {
		newIdx, newDoc := bsoncore.AppendDocumentStart(nil)
		newDoc = bsoncore.AppendDocumentElement(newDoc, "x", doc)
		newDoc, _ = bsoncore.AppendDocumentEnd(newDoc, newIdx)

		doc = newDoc
	}

	return doc
}

func getStringElement(idx int) (key, value string) {
	return fmt.Sprintf("field%v", idx), "blabla"
}
