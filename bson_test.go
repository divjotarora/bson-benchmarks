package main

import (
	"fmt"
	"testing"
	"time"

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
			doc := buildDoc(tc.docSize)

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
			doc := buildRawDoc(tc.docSize)

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

func BenchmarkUnmarshalNested(b *testing.B) {
	testCases := []struct {
		name  string
		depth int
	}{
		// No "empty" test because a doc without any nesting is already covered by BenchmarkUnmarshal
		{"small", 1},
		{"size 10", 10},
		{"size 50", 50},
		{"size 100", 100},
		{"size 500", 500},
		{"size 1000", 1000},
	}
	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			doc := buildNestedRawDoc(tc.depth)

			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				var unmarshalled bson.D
				err := bson.Unmarshal(doc, &unmarshalled)
				if err != nil {
					b.Error(err)
				}
			}
		})
	}
}

func BenchmarkUnmarshalIsMaster(b *testing.B) {
	isMasterBytes := buildIsMasterResponse(b)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var doc bson.D
		err := bson.Unmarshal(isMasterBytes, &doc)
		if err != nil {
			b.Error(err)
		}
	}
}

func buildIsMasterResponse(b *testing.B) []byte {
	response := bson.D{
		{"ismaster", true},
		{"maxBsonObjectSize", 16777216},
		{"maxMessageSizeBytes", 48000000},
		{"maxWriteBatchSize", 100000},
		{"localTime", time.Now()},
		{"logicalSessionTimeoutMinutes", 30},
		{"minWireVersion", 0},
		{"maxWireVersion", 6},
		{"readOnly", false},
		{"hostsBsonD", []interface{}{
			bson.E{"host", "blabla1"},
			bson.E{"host", "blabla2"},
			bson.E{"host", "blabla3"},
		}},
		{"hostsIf", []interface{}{
			bson.D{{"host", "blabla1"}},
			bson.D{{"host", "blabla2"}},
			bson.D{{"host", "blabla3"}},
		}},
	}

	res, err := bson.Marshal(response)
	if err != nil {
		b.Fatal(err)
	}
	return res
}

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
