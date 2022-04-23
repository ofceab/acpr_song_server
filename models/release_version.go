package models

import (
	"time"
)

// A ReleaseVersion is a number for keeping in track songs
// To each version of song will be associate a unique version

type ReleaseVersion struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreatedTime"`
}
