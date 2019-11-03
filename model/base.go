package model

import (
    "time"

    "github.com/jinzhu/gorm"
    "github.com/satori/go.uuid"
)

// Base contains common columns for all tables.
type Base struct {
    ID        uuid.UUID  `json:"id";gorm:"primary_key;type:char(36);`
    CreatedAt time.Time  `json:"created_at";`
    UpdatedAt time.Time  `json:"updated_at";`
    DeletedAt *time.Time `json:"-";`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
    uuid := uuid.NewV4()

    return scope.SetColumn("ID", uuid)
}
