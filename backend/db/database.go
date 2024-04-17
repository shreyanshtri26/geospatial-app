package database

import (
    "fmt"
    "log"
    "os"

    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq" // PostgreSQL driver
    "github.com/shreyanshtri26/geospatial-app/models"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")

    dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword) 

    db, err := gorm.Open("postgres", dbURI)
    if err != nil {
        log.Fatal("Error connecting to database:", err)
    }

    // Auto-migrate models (create tables if they don't exist)
    db.AutoMigrate(&models.User{}, &models.GeoData{}, &models.File{}) 

    DB = db
    return db
}
