# Local Setup

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

# CI Setup

When running in CI, check out the `nonlocal-gomod` branch in this project. The changes on that branch modify the
`go.mod` file to only depend on remote repositories rather than local copies. For benchmarks using mgo, the `go.mod`
file replaces `gopkg.in/mgo.v2` with `github.com/divjotarora` pinned to the `modules` branch.

# Run Explanations

## Starting

This run was used to get the beginning baseline numbers:

- Driver commit: 09ccd6fca3f93c26d6b2c74c40ed4ba4705eb163
- mgo commit: 7f441db88ff27b0d5be438510c30e1881b3fa2f6 (tip of divjotarora/mgo@modules)

Numbers for both the driver and mgo, as well as comparison numbers generated using `benchstat` are available in the
`benchmarks/starting` directory. The `benchstat` tool was invoked using `benchstat mgo.bench driver.bench > comparison.bench`.

## Basic Unmarshal Improvements

This run was conducted to measure the effect of making the bare minimum improvements to the `bson.Unmarshal` code path.
These improvements included:

- Add a `typeDecoder` interface
- Add a dedicated `DDecodeValue` decoder for `primitive.D` instances
- Change `EmptyInterfaceCodec` to implement `typeDecoder` and delegate to `typeDecoder` when possible, falling back to `ValueDecoder` when not.

Details:

- Driver commit: [e33fb05539f40505048d3a60ee2f8620811b16ae](https://github.com/divjotarora/mongo-go-driver/commit/e33fb05539f40505048d3a60ee2f8620811b16ae)
- Same mgo commit

Numbers can be found in the `benchmarks/basic-unmarshal-improvements` directory.