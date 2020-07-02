module github.com/divjotarora/bson-benchmarks

go 1.13

replace gopkg.in/mgo.v2 => github.com/divjotarora/mgo modules

require (
	go.mongodb.org/mongo-driver 09ccd6fca3f93c26d6b2c74c40ed4ba4705eb163
	gopkg.in/mgo.v2 master
)
