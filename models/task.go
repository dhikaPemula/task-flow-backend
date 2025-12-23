package models

import (
    "time"

    "github.com/google/uuid"
)

type Task struct {
    ID          uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    Title       string     `gorm:"not null" json:"title"`
    Description string     `json:"description"`
    Status      string     `gorm:"not null;default:'todo'" json:"status"`
    Priority    string     `gorm:"not null;default:'medium'" json:"priority"`
    DueDate     *time.Time `json:"due_date"`
    CategoryID  *uuid.UUID `gorm:"type:uuid" json:"category_id"`
    Category    *Category  `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
    UserID      uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
    User        User       `gorm:"foreignKey:UserID" json:"-"`
    CreatedAt   time.Time  `json:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at"`
}