package main

import (
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

// Benchmarks ported from https://github.com/mongodb/mongonet/blob/Benchmarks/bsonutil_test.go

func BenchmarkMarshal(b *testing.B) {
	testCases := []struct {
		name    string
		docSize int
	}{
		{"empty", 0},
		{"small", 1},
		{"size 10", 10},
		{"size 50", 50},
		{"size 100", 100},
		{"size 500", 500},
		{"size 1000", 1000},
	}
	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			doc := getDocOfSize(tc.docSize)

			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, err := bson.Marshal(doc)
				if err != nil {
					b.Error(err)
				}
			}
		})
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	testCases := []struct {
		name    string
		docSize int
	}{
		{"empty", 0},
		{"small", 1},
		{"size 10", 10},
		{"size 50", 50},
		{"size 100", 100},
		{"size 500", 500},
		{"size 1000", 1000},
	}
	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			doc := getRawDocOfSize(tc.docSize)

			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				unmarshalled := bson.D{}
				err := bson.Unmarshal(doc, &unmarshalled)
				if err != nil {
					b.Error(err)
				}
			}
		})
	}
}

func getDocOfSize(n int) bson.D {
	doc := bson.D{}
	for i := 0; i < n; i++ {
		key, val := getStringElement(i)
		doc = append(doc, bson.E{key, val})
	}
	return doc
}

func getRawDocOfSize(n int) bsoncore.Document {
	idx, doc := bsoncore.AppendDocumentStart(nil)
	for i := 0; i < n; i++ {
		key, val := getStringElement(i)
		doc = bsoncore.AppendStringElement(doc, key, val)
	}

	doc, _ = bsoncore.AppendDocumentEnd(doc, idx)
	return doc
}

func getStringElement(idx int) (key, value string) {
	return fmt.Sprintf("field%v", idx), "blabla"
}
