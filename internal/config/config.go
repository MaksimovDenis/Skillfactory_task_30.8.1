package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	APIPort         int    `split_words:"true" default:"8080"`
	MongoConnString string `split_words:"true" default:"mongodb://localhost:27017"`
	PgConnString    string `split_words:"true" default:"postgres://admin:admin@localhost:5432/skillfactory?sslmode=disable"`
	LogLevel        string `split_words:"true" default:"debug"`
}

// InitConfig init config
func InitConfig() (*Config, error) {
	var cnf Config
	err := envconfig.Process("SKILLFACTORY", &cnf)
	if err != nil {
		return nil, err
	}

	return &cnf, nil
}
