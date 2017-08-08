package config

import (
	"fmt"

	"github.com/namsral/flag"
)

func makeConfig() IAppConfig {
	var port, mongoURL, mongoDB string

	flag.StringVar(&port, "PORT", "8000", "port number")
	flag.StringVar(&mongoURL, "MONGO_URL", "", "MongoDB connection string")
	flag.StringVar(&mongoDB, "MONGO_DB", "simple-api-db", "MongoDB database name")

	flag.Parse()

	if mongoURL == "" {
		fmt.Println("[!] Missing MONGO_URL env variable, database won't be accessible.")
	}

	return IAppConfig{
		Port:     port,
		MongoURL: mongoURL,
		MongoDB:  mongoDB,
	}
}

// IAppConfig contains app configuration
type IAppConfig struct {
	Port     string
	MongoURL string
	MongoDB  string
}

// App contains all dynamic configuration
var App = makeConfig()
