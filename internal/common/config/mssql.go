package config

import "os"

func MSSQL() mssqlConfig {
	empty := mssqlConfig{}
	if mssql == empty {
		mssql = mssqlConfig{
			Username:     getEnvWithDefault("DB_USERNAME", "sa"),
			Password:     getEnvWithDefault("DB_PASSWORD", "MyStrongPassword123"),
			Host:         getEnvWithDefault("DB_HOST", "localhost"),
			Port:         getEnvWithDefault("DB_PORT", "1433"),
			DatabaseName: getEnvWithDefault("DB_NAME", "todoapp"),
		}
	}
	return mssql
}

type mssqlConfig struct {
	Username     string
	Password     string
	Host         string
	Port         string
	DatabaseName string
}

var mssql mssqlConfig

func getEnvWithDefault(envName, defaultVal string) string {
	env := os.Getenv(envName)
	if env == "" {
		env = defaultVal
	}
	return env
}
