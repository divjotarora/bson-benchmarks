module github.com/divjotarora/bson-benchmarks

go 1.13

replace gopkg.in/mgo.v2 => github.com/divjotarora/mgo v0.0.0-20200626010915-7f441db88ff2

require (
	go.mongodb.org/mongo-driver v1.4.0-beta2.0.20200727183953-ec900457b075
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)
