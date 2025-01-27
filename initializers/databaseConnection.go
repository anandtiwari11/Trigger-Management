package initializers

import (
	"log"
	"os"

	"github.com/anandtiwari11/event-trigger/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDB() {
    err = godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
    dsn := "host=" + os.Getenv("DB_HOST") +
        " user=" + os.Getenv("DB_USER") +
        " password=" + os.Getenv("DB_PASSWORD") +
        " dbname=" + os.Getenv("DB_NAME") +
        " port=" + os.Getenv("DB_PORT") +
        " sslmode=" + os.Getenv("DB_SSLMODE")
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    log.Println("Connected to PostgreSQL database successfully!")
    AutoMigrateTables()
}

func AutoMigrateTables() {
    err := DB.AutoMigrate(
        &models.Event{},
        &models.Trigger{},
    )
    if err != nil {
        log.Fatalf("Failed to auto-migrate tables: %v", err)
    }

    log.Println("Auto-migration completed successfully!")
}
