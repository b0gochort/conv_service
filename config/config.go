package config

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		DataBase     PG
		JWTSecretKey string
	}

	PG struct {
		Addr         string `env:"ADDRESS" env-default:"localhost:5432"`
		User         string `env:"USER" env-default:"postgres"`
		Password     string `env:"PASSWORD" env-default:"king1337"`
		DataBaseName string `env:"DATABASE" env-default:"coursch"`
	}
)

func LoadConfig() (config *Config, err error) {
	cfg := &Config{}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
