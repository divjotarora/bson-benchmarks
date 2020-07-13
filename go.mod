module github.com/divjotarora/bson-benchmarks

go 1.13

replace gopkg.in/mgo.v2 => github.com/divjotarora/mgo modules
replace go.mongodb.org/mongo-driver master => github.com/divjotarora/mongo-go-driver godriver1682-recursive-typedecoder

require (
	go.mongodb.org/mongo-driver master
	gopkg.in/mgo.v2 master
)
