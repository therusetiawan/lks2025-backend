package migrations

import (
    "log"
    "github.com/therusetiawan/lks2025cc-backend/config"
    "github.com/therusetiawan/lks2025cc-backend/models"
)

func Migrate() {
    err := config.DB.AutoMigrate(&models.Student{})
    if err != nil {
        log.Fatal("Migration failed:", err)
    }
    log.Println("Database migration completed successfully.")
}
