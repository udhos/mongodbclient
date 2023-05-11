[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/udhos/mongodbclient/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/udhos/mongodbclient)](https://goreportcard.com/report/github.com/udhos/mongodbclient)
[![Go Reference](https://pkg.go.dev/badge/github.com/udhos/mongodbclient.svg)](https://pkg.go.dev/github.com/udhos/mongodbclient)

# mongodbclient

mongodbclient creates a MongoDB client.

## Run mongodb server locally for testing

```
docker run --rm --name mongo-main -p 27017:27017 -d mongo
```

## Run sample application

See sample application: [examples/mongodbclient-ping/main.go](examples/mongodbclient-ping/main.go)

```
./build.sh

mongodbclient-ping
```

Example:

Start mongodb:

```bash
docker run --rm --name mongo-main -p 27017:27017 mongo
```

Fire up `mongodbclient-ping`:

```bash
$ mongodbclient-ping 
2023/05/11 02:18:36 mongodbclient-ping version=0.1.0 runtime=go1.20.4 boilerplate=1.0.1 GOOS=linux GOARCH=amd64 GOMAXPROCS=12
2023/05/11 02:18:36 envconfig.NewSimple: SECRET_ROLE_ARN=''
2023/05/11 02:18:36 MONGO_URI=[] using MONGO_URI=mongodb://localhost:27017/ default=mongodb://localhost:27017/
2023/05/11 02:18:36 MONGO_USER=[] using MONGO_USER= default=
2023/05/11 02:18:36 MONGO_PASS=[] using MONGO_PASS= default=
2023/05/11 02:18:36 MONGO_CA_FILE=[] using MONGO_CA_FILE= default=
2023/05/11 02:18:36 mongodbclient.New: mongo connect URI: mongodb://localhost:27017/
2023/05/11 02:18:36 mongo ping: ok: elapsed 2.027956ms
2023/05/11 02:18:38 mongo ping: ok: elapsed 541.742µs
2023/05/11 02:18:40 mongo ping: ok: elapsed 1.138481ms
2023/05/11 02:18:42 mongo ping: ok: elapsed 1.178109ms
2023/05/11 02:18:44 mongo ping: ok: elapsed 1.12218ms
```
