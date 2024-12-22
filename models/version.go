package models

import (
	"time"
)

type Version struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ServiceID    uint      `gorm:"not null" json:"service_id"`
	Version      string    `gorm:"type:text;not null" json:"version"`
	ReleaseNotes string    `gorm:"type:text" json:"release_notes"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
