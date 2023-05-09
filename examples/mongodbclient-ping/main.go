// Package main implements an example application.
package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/udhos/boilerplate/boilerplate"
	"github.com/udhos/boilerplate/envconfig"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/udhos/mongodbclient/mongodbclient"
)

const version = "0.1.0"

func main() {

	me := filepath.Base(os.Args[0])

	log.Print(boilerplate.LongVersion(me + " version=" + version))

	env := envconfig.NewSimple(me)

	clientOptions := mongodbclient.Options{
		Debug:     true,
		URI:       env.String("MONGO_URI", "mongodb://localhost:27017/"),
		Username:  env.String("MONGO_USER", ""),
		Password:  env.String("MONGO_PASS", ""),
		TLSCAFile: env.String("MONGO_CA_FILE", ""),
	}

	client, errClient := mongodbclient.New(clientOptions)
	if errClient != nil {
		log.Fatalf("%s: create client error: %v", me, errClient)
	}

	for {
		ping(client, 2*time.Second)
		time.Sleep(2 * time.Second)
	}
}

func ping(client *mongo.Client, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	begin := time.Now()

	errPing := client.Ping(ctx, nil)

	elap := time.Since(begin)

	if errPing != nil {
		log.Printf("mongo ping: error: %v", errPing)
	} else {
		log.Printf("mongo ping: ok: elapsed %v", elap)
	}
}
