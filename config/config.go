package config

import "os"

const CHANGE_ME_STATE = "CHANGE_ME"
const DEVELOP_MODE = "develop"
const PROD_MOD = "release"

type Config struct {
	Port               string
	JwtSecrets         string
	DBConnectionString string
	AppMode            string
}

func New() Config {
	return Config{
		Port:               getValue("PORT", "8080"),
		JwtSecrets:         getValue("JWT_SECRETS", CHANGE_ME_STATE),
		DBConnectionString: getValue("CONN_STRING", CHANGE_ME_STATE),
		AppMode:            getValue("APP_MODE", "develop"),
	}
}

func getValue(key, def string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return def
}
