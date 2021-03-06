// +build !mgo

package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Helper functions to build command request/response documents using the driver's BSON libray.

func getIsMasterResponse() bson.D {
	return bson.D{
		{"ismaster", true},
		{"maxBsonObjectSize", 16777216},
		{"maxMessageSizeBytes", 48000000},
		{"maxWriteBatchSize", 100000},
		{"localTime", time.Now()},
		{"logicalSessionTimeoutMinutes", 30},
		{"minWireVersion", 0},
		{"maxWireVersion", 6},
		{"readOnly", false},
		{"hostsBsonD", bson.A{
			bson.D{{"host", "blabla1"}},
			bson.D{{"host", "blabla2"}},
			bson.D{{"host", "blabla3"}},
		}},
		{"hostsIf", bson.A{
			bson.D{{"host", "blabla1"}},
			bson.D{{"host", "blabla2"}},
			bson.D{{"host", "blabla3"}},
		}},
	}
}

func getFindOneRequest() bson.D {
	return bson.D{
		{"find", "bla"},
		{"$db", "test"},
		{"filter", bson.D{{"b", 1}}},
		{"limit", float64(1)},
		{"singleBatch", true},
		{"lsid", bson.D{
			{"id", primitive.Binary{
				Subtype: uint8(4),
				Data:    []byte("blalblalbalblablalabl"),
			}},
		}},
		{"$clusterTime", bson.D{
			{"clusterTime", primitive.Timestamp{
				T: uint32(1593340459),
				I: uint32(1),
			}},
			{"signature", bson.D{
				{"hash", primitive.Binary{
					Subtype: uint8(4),
					Data:    []byte("blalblalbalblablalablibibibibibibibi"),
				}},
				{"keyId", int64(6843344346754842627)},
			}},
		}},
	}
}

func getFindOneResponse() bson.D {
	return bson.D{
		{"cursor", bson.D{
			{"id", int64(0)},
			{"ns", "eliot1-bla.test"},
			{"firstBatch", bson.A{
				bson.D{
					{"_id", primitive.NewObjectID()},
					{"a", 1},
				},
			}},
		}},
		{"$db", "test"},
		{"ok", 1},
		{"lsid", bson.D{
			{"id", primitive.Binary{
				Subtype: uint8(4),
				Data:    []byte("blalblalbalblablalabl"),
			}},
		}},
		{"$clusterTime", bson.D{
			{"clusterTime", primitive.Timestamp{
				T: uint32(1593340459),
				I: uint32(1),
			}},
			{"signature", bson.D{
				{"hash", primitive.Binary{
					Subtype: uint8(4),
					Data:    []byte("blalblalbalblablalablibibibibibibibi"),
				}},
				{"keyId", int64(6843344346754842627)},
			}},
		}},
		{"operationTime", primitive.Timestamp{
			T: uint32(1593340459),
			I: uint32(1),
		}},
	}
}

func getLargeInsertRequest() bson.D {
	isMasterResponse := getIsMasterResponse()
	for i := 0; i < 50; i++ {
		isMasterResponse = bson.D{{"nest", isMasterResponse}}
	}

	var insertDocs bson.A
	for i := 0; i < 200; i++ {
		insertDocs = append(insertDocs, isMasterResponse)
	}

	return bson.D{
		{"insert", "collection"},
		{"documents", insertDocs},
		{"$db", "test"},
		{"lsid", bson.D{
			{"id", primitive.Binary{
				Subtype: uint8(4),
				Data:    []byte("blalblalbalblablalabl"),
			}},
		}},
		{"$clusterTime", bson.D{
			{"clusterTime", primitive.Timestamp{
				T: uint32(1593340459),
				I: uint32(1),
			}},
			{"signature", bson.D{
				{"hash", primitive.Binary{
					Subtype: uint8(4),
					Data:    []byte("blalblalbalblablalablibibibibibibibi"),
				}},
				{"keyId", int64(6843344346754842627)},
			}},
		}},
	}
}
