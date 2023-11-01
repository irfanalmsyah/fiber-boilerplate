package config

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
)

func LoadEnv() {
    err := godotenv.Load(".env")
    if err != nil {
        fmt.Printf("Error loading .env file: %v", err)
    }
}

func Config(key string) (string, error) {
    value := os.Getenv(key)
    if value == "" {
        return "", fmt.Errorf("%s env variable not set", key)
    }
    return value, nil
}