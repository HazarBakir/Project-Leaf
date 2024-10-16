package main

import (
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var db *gorm.DB
var err error

// User struct
type User struct {
    UserID    uint   `gorm:"primaryKey"`
    Username  string `gorm:"size:255"`
    Email     string `gorm:"size:255"`
    Password  string `gorm:"size:255"`
    CreatedAt string `gorm:"autoCreateTime"`
    UpdatedAt string `gorm:"autoUpdateTime"`
}

// SavedBook struct
type SavedBook struct {
    SaveID    uint   `gorm:"primaryKey"`
    UserID    uint   `gorm:"index"`
    BookTitle string `gorm:"size:255"`
    SavedAt   string `gorm:"autoCreateTime"`
}

func ConnectDatabase() {
    dsn := "host=localhost user=myuser password=K4n4sh1h4z4r! dbname=library_db port=5432 sslmode=disable"
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    // Auto Migrate
    db.AutoMigrate(&User{}, &SavedBook{})
}