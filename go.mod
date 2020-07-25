module github.com/divjotarora/bson-benchmarks

go 1.13

replace gopkg.in/mgo.v2 => github.com/divjotarora/mgo v0.0.0-20200626010915-7f441db88ff2

replace go.mongodb.org/mongo-driver => github.com/divjotarora/mongo-go-driver v0.0.7-0.20200725012058-3c67c25251b3

require (
	go.mongodb.org/mongo-driver v1.4.0-beta2.0.20200724175811-46d5d7678dc1
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)
