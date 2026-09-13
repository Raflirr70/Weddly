package config

import "os"

type Config struct {
	DBHost      string
	DBPort      string
	DBUser      string
	DBPass      string
	DBName      string
	RedisHost   string
	RedisPort   string
	KafkaBroker string
	JWTSecret   string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:      getEnv("DB_HOST"),
		DBPort:      getEnv("DB_PORT"),
		DBUser:      getEnv("DB_USER"),
		DBPass:      getEnv("DB_PASS"),
		DBName:      getEnv("DB_NAME"),
		RedisHost:   getEnv("REDIS_HOST"),
		RedisPort:   getEnv("REDIS_PORT"),
		KafkaBroker: getEnv("KAFKA_BROKER"),
		JWTSecret:   getEnv("JWT_SECRET"),
	}
}

func getEnv(key string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return ""
}
