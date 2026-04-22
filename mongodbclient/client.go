// Package mongodbclient creates a mongodb client.
package mongodbclient

import (
	"log"
	"net/url"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Options specifies parameters for mongodb client.
type Options struct {
	URI         string
	Username    string
	Password    string
	TLSCAFile   string
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
		mongoOptions := options.Client().ApplyURI(uri).SetRetryWrites(false).SetMinPoolSize(opt.MinPoolSize)

		if opt.Username != "" || opt.Password != "" {
			cred := options.Credential{
				Username: opt.Username,
				Password: opt.Password,
			}
			mongoOptions.SetAuth(cred)
		}

		var errConnect error
		client, errConnect = mongo.Connect(mongoOptions)
		if errConnect != nil {
			if debug {
				logf("%s: mongo connect: %v", me, errConnect)
			}
			return nil, errConnect
		}
	}

	return client, nil
}
