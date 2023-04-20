// Package mongodbclient creates a mongodb client.
package mongodbclient

import (
	"context"
	"log"
	"net/url"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Options specifies parameters for mongodb client.
type Options struct {
	//Database  string
	URI         string
	Username    string
	Password    string
	TLSCAFile   string
	Timeout     time.Duration // Defaults to 10 seconds
	MinPoolSize uint64
	Logf        func(format string, v ...any) // Defaults to log.Printf
	Debug       bool                          // Log debug messages
}

// DefaultTimeout is used when Options Timeout isn't specified.
const DefaultTimeout = 10 * time.Second

// New creates a mongodb client.
func New(opt Options) (*mongo.Client, error) {
	const me = "mongodbclient.New"

	debug := opt.Debug

	logf := opt.Logf
	if logf == nil {
		logf = log.Printf
	}

	timeout := opt.Timeout
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	//
	// URI
	//

	var uri string

	{
		u, errURI := url.Parse(opt.URI)
		if errURI != nil {
			if debug {
				logf("%s: mongo connect parse URI: %v", me, errURI)
			}
			return nil, errURI
		}

		if opt.TLSCAFile != "" {
			q := u.Query()

			q.Set("ssl", "true")
			q.Set("tlsCAFile", opt.TLSCAFile)
			q.Set("ssl_ca_certs", opt.TLSCAFile) // documentdb?

			u.RawQuery = q.Encode()
		}

		uri = u.String()

		if debug {
			logf("%s: mongo connect URI: %s", me, uri)
		}
	}

	//
	// connect
	//

	var client *mongo.Client

	{
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		mongoOptions := options.Client().ApplyURI(uri).SetRetryWrites(false).SetMinPoolSize(opt.MinPoolSize)

		if opt.Username != "" || opt.Password != "" {
			cred := options.Credential{
				Username: opt.Username,
				Password: opt.Password,
			}
			mongoOptions.SetAuth(cred)
		}

		var errConnect error
		client, errConnect = mongo.Connect(ctx, mongoOptions)
		if errConnect != nil {
			if debug {
				logf("%s: mongo connect: %v", me, errConnect)
			}
			return nil, errConnect
		}
	}

	return client, nil
}
