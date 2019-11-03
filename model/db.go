package model

import (
    "os"
    "log"
    "time"

    "github.com/joho/godotenv"
    "github.com/satori/go.uuid"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    _ "github.com/jinzhu/gorm/dialects/mssql"

)

var database *gorm.DB

func init() {
    if database != nil {
        if err := database.DB().Ping(); err == nil {
            log.Fatalln(err.Error())
        }
    }
    
    if err := godotenv.Load(); err != nil {
        log.Fatalln("Error loading .env file")
    }

    dbEngine := os.Getenv("DB_ENGINE")
    dbConnect := os.Getenv("DB_CONNECT")
    dbLogMode := os.Getenv("DB_LOG_MODE")

    var db *gorm.DB
    var err error

    db, err = gorm.Open(dbEngine, dbConnect)
    if err != nil {
        log.Fatalln(err.Error())
    }

    // defer db.Close()

    db.LogMode(dbLogMode == "true")
    db.AutoMigrate(&Chemical{})
    
    database = db
}

func GetDatabase() *gorm.DB {
    return database
}

// Base contains common columns for all tables.
type Base struct {
    ID        uuid.UUID `json:"id";gorm:"primary_key;type:char(36);`
    CreatedAt time.Time `json:"created_at";`
    UpdatedAt time.Time `json:"updated_at";`
    DeletedAt *time.Time `json:"-";`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
    uuid := uuid.NewV4()

    return scope.SetColumn("ID", uuid)
}