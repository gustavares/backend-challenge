package config

import "os"

type Config struct {
	Env *Env
}

type Env struct {
	DbName     string
	DbUser     string
	DbPassword string
	DbHost     string
}

func envSetup() *Env {
	return &Env{
		DbName:     os.Getenv("DB_NAME"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbHost:     os.Getenv("DB_HOST"),
	}
}

func New() *Config {
	env := envSetup()

	return &Config{
		env,
	}
}
