module github.com/divjotarora/bson-benchmarks

go 1.13

replace go.mongodb.org/mongo-driver => /home/divjot/code/mongo-go-driver

replace gopkg.in/mgo.v2 => /home/divjot/go/src/github.com/10gen/mgo

require (
	go.mongodb.org/mongo-driver v1.3.4
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)
