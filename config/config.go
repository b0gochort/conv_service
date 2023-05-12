package config

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		HTTP
		PG
		JWTSecretKey
	}

	PG struct {
		Addr         string `env:"ADDRESS" env-default:"localhost:5432"`
		User         string `env:"USER" env-default:"postgres"`
		Password     string `env:"PASSWORD" env-default:"king1337"`
		DataBaseName string `env:"DATABASE" env-default:"coursch"`
	}
	HTTP struct {
		PORT string `env:"HTTP_PORT" env-default:":8000"`
	}
	JWTSecretKey struct {
		JWTSecretKey string `env:"JWT_SECRET_KEY" env-default:"secretKey"`
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

var i int = 10
