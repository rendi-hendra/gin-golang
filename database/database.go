package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gin-golang/repository"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	if errEnv := godotenv.Load(); errEnv != nil {
		log.Fatal("Error loading .env file", errEnv)
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if user == "" || host == "" || port == "" || name == "" {
		log.Fatal("Database environment variables (DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME) must be set")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed getting sql DB object: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)

	if err := db.AutoMigrate(&repository.User{}); err != nil {
		log.Fatalf("failed auto migrate: %v", err)
	}

	DB = db
	fmt.Println("Database Connected!")
}
