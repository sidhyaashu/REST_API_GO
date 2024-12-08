package main

import (
	"fmt"
	"os" 
	"log" 
)

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	JWTSecret  string
}

// Envs holds the configuration loaded from environment variables.
var Envs = initConfig()

func initConfig() Config {
	return Config{
		Port:       getEnv("PORT", "8080"),        
		DBUser:     getEnv("DB_USER", "root"),     
		DBPassword: getEnv("DB_PASSWORD", "rootpassword"), 
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "projectmanager"), 
		JWTSecret:  getEnv("JWT_SECRET", "myrandomsecret"), 
	}
}

func getEnv(key, fallback string) string {
	// Look up the environment variable, if it exists return its value, otherwise return the fallback value
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	// Log a message if the environment variable is missing
	log.Printf("Warning: Environment variable %s not set, using fallback value: %s", key, fallback)
	return fallback
}