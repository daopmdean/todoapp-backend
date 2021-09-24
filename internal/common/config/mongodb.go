package config

import "os"

func GetMongoUri() string {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		panic("Please provide MONGO_URI environment")
	}
	return uri
}

func GetMongoDBName() string {
	uri := os.Getenv("MONGO_DB_NAME")
	if uri == "" {
		return "todoapp"
	}
	return uri
}
