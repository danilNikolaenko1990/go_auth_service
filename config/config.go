package config

type Config struct {
	DbUser     string `envconfig:"AUTH_SERVICE_DB_USER"`
	DbPassword string `envconfig:"AUTH_SERVICE_DB_PASSWORD"`
	DbHost     string `envconfig:"AUTH_SERVICE_DB_HOST"`
	DbPort     string `envconfig:"AUTH_SERVICE_DB_PORT"`
	DbName     string `envconfig:"AUTH_SERVICE_DB_NAME"`
	ServerPort string `envconfig:"AUTH_SERVICE_DB_SERVER_PORT"`
}
