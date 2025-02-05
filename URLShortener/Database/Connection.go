package Database

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)


const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "09011155"
	DBName   = "url_shortener"
)

var DB *gorm.DB

func ConnectDatabase() {
    // ایجاد رشته اتصال (DSN) برای PostgreSQL
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
        Host, User, Password, DBName, Port)

    // اتصال به پایگاه داده PostgreSQL
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatal("Failed to connect to the database!", err)
    }

    DB = db
    log.Println("Database connected...")
}