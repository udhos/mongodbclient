// Package main implements an example application.
package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/udhos/boilerplate/envconfig"
	"github.com/udhos/boilerplate/secret"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/udhos/mongodbclient/mongodbclient"
)

func main() {

	me := filepath.Base(os.Args[0])

	env := getEnv(me)

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

func getEnv(me string) *envconfig.Env {
	roleArn := os.Getenv("SECRET_ROLE_ARN")

	log.Printf("SECRET_ROLE_ARN='%s'", roleArn)

	secretOptions := secret.Options{
		RoleSessionName: me,
		RoleArn:         roleArn,
	}
	secret := secret.New(secretOptions)
	envOptions := envconfig.Options{
		Secret: secret,
	}
	env := envconfig.New(envOptions)
	return env
}

func ping(client *mongo.Client, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if errPing := client.Ping(ctx, nil); errPing != nil {
		log.Printf("mongo ping: error: %v", errPing)
	} else {
		log.Printf("mongo ping: ok")
	}
}
