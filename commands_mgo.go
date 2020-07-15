// +build mgo

package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Helper functions to build command request/response documents using mgo's BSON libray.

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
		{"hostsBsonD", []bson.D{
			bson.D{{"host", "blabla1"}},
			bson.D{{"host", "blabla2"}},
			bson.D{{"host", "blabla3"}},
		}},
		{"hostsIf", []interface{}{
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
			{"id", bson.Binary{
				Kind: uint8(4),
				Data: []byte("blalblalbalblablalabl"),
			}},
		}},
		{"$clusterTime", bson.D{
			{"clusterTime", bson.MongoTimestamp(1593340459)},
			{"signature", bson.D{
				{"hash", bson.Binary{
					Kind: uint8(4),
					Data: []byte("blalblalbalblablalablibibibibibibibi"),
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
			{"firstBatch", []bson.D{
				bson.D{
					{"_id", bson.NewObjectId()},
					{"a", 1},
				},
			}},
		}},
		{"$db", "test"},
		{"ok", 1},
		{"lsid", bson.D{
			{"id", bson.Binary{
				Kind: uint8(4),
				Data: []byte("blalblalbalblablalabl"),
			}},
		}},
		{"$clusterTime", bson.D{
			{"clusterTime", bson.MongoTimestamp(1593340459)},
			{"signature", bson.D{
				{"hash", bson.Binary{
					Kind: uint8(4),
					Data: []byte("blalblalbalblablalablibibibibibibibi"),
				}},
				{"keyId", int64(6843344346754842627)},
			}},
		}},
		{"operationTime", bson.MongoTimestamp(1593340459)},
	}
}

func getLargeInsertRequest() bson.D {
	isMasterResponse := getIsMasterResponse()
	for i := 0; i < 50; i++ {
		isMasterResponse = bson.D{{"nest", isMasterResponse}}
	}

	var insertDocs []bson.D
	for i := 0; i < 200; i++ {
		insertDocs = append(insertDocs, isMasterResponse)
	}

	return bson.D{
		{"insert", "collection"},
		{"documents", insertDocs},
		{"$db", "test"},
		{"lsid", bson.D{
			{"id", bson.Binary{
				Kind: uint8(4),
				Data: []byte("blalblalbalblablalabl"),
			}},
		}},
		{"$clusterTime", bson.D{
			{"clusterTime", bson.MongoTimestamp(1593340459)},
			{"signature", bson.D{
				{"hash", bson.Binary{
					Kind: uint8(4),
					Data: []byte("blalblalbalblablalablibibibibibibibi"),
				}},
				{"keyId", int64(6843344346754842627)},
			}},
		}},
	}
}
