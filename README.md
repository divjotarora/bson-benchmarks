# Setup

Clone the driver from `git@github.com:mongodb/mongo-go-driver.git`.

Clone divjotarora/mgo. This is a fork of 10gen/mgo with a new
`modules` branch to enable Go modules support. The branch is based on mgo commit
`7ddd511871dec26ace0517365e4e496b545159b5`:

```
git clone git@github.com:divjotarora/mgo.git
cd mgo
git checkout modules
```

After cloning both repositories, modify the relevant `replace` directives in `go.mod` to point to your local
installations:

```
replace go.mongodb.org/mongo-driver => /path/to/driver/installation
replace gopkg.in/mgo.v2 => /patch/to/mgo/installation
```
