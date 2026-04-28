package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pass string `env:"DB_PASS"`
	Name string `yaml:"name"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type JWTConfig struct {
	Secret string `env:"JWT_SECRET"`
}
type Config struct {
	DB     DBConfig     `yaml:"db"`
	Server ServerConfig `yaml:"server"`
	JWT    JWTConfig    `yaml:"jwt"`
}

// Reading the .env
// Loading config
func MustLoad() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := &Config{}
	err = cleanenv.ReadConfig("config.yml", config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
