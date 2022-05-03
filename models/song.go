package models

import "time"

// Describe a song
type Song struct {
	ID uint `json:"id"`
	// Title of the song
	Title string `json:"title"`
	// Lyrics of the song
	Lyrics string `json:"lyrics"`
	// Audio of the song
	AudioUrl string `json:"audio_url"`
	// Release version
	ReleaseVersionId uint           `json:"release_version_id"`
	ReleaseVersion   ReleaseVersion `json:"release_version" gorm:"foreignKey:ReleaseVersionId"`
	// Created at
	CreatedAt time.Time `json:"created_at" gorm:"autoCreatedTime"`
}
