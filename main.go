package main

import (
    "log"
    "github.com/joho/godotenv"
    "github.com/therusetiawan/lks2025cc-backend/config"
    "github.com/therusetiawan/lks2025cc-backend/migrations"
    "github.com/therusetiawan/lks2025cc-backend/routes"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    config.ConnectDatabase()
    config.ConnectRedis()
    config.ConnectS3()
    migrations.Migrate()

    r := routes.SetupRouter()
    r.Run(":8080")
}
