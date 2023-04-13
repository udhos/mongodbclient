# mongodbclient
mongodbclient

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

```
$ mongodbclient-ping
2023/04/13 14:55:23 SECRET_ROLE_ARN=''
2023/04/13 14:55:23 MONGO_URI=[] using MONGO_URI=mongodb://localhost:27017/ default=mongodb://localhost:27017/
2023/04/13 14:55:23 MONGO_USER=[] using MONGO_USER= default=
2023/04/13 14:55:23 MONGO_PASS=[] using MONGO_PASS= default=
2023/04/13 14:55:23 MONGO_CA_FILE=[] using MONGO_CA_FILE= default=
2023/04/13 14:55:23 mongodbclient.New: mongo connect URI: mongodb://localhost:27017/
2023/04/13 14:55:23 mongo ping: ok
2023/04/13 14:55:25 mongo ping: ok
2023/04/13 14:55:27 mongo ping: ok
2023/04/13 14:55:29 mongo ping: ok
2023/04/13 14:55:31 mongo ping: ok
2023/04/13 14:55:33 mongo ping: ok
```
