package config

import "os"

func GetAtlasUri() string {
	uri := os.Getenv("ATLAS_URI")
	if uri == "" {
		panic("Please provide ATLAS_URI environment")
	}
	return uri
}
