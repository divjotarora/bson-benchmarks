// +build mgo

package main

import "gopkg.in/mgo.v2/bson"

func buildDoc(n int) bson.D {
	doc := bson.D{}
	for i := 0; i < n; i++ {
		key, val := getStringElement(i)
		doc = append(doc, bson.DocElem{key, val})
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
