package configuration

import "github.com/caarlos0/env"

type ServerConfig struct {
	Address AddressConfig
}

type AddressConfig struct {
	Port string `env:"PORT" envDefault:"8081"`
}

func Read() (*ServerConfig, error) {
	config := ServerConfig{}
	if err := env.Parse(&config); err != nil {
		return nil, err
	}
	if err := env.Parse(&config.Address); err != nil {
		return nil, err
	}
	return &config, nil
}
