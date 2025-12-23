package models

import (
    "time"

    "github.com/google/uuid"
)

type Category struct {
    ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    Name      string    `gorm:"not null" json:"name"`
    Color     string    `gorm:"not null" json:"color"`
    UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
    User      User      `gorm:"foreignKey:UserID" json:"-"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}