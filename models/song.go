package models

import "time"

// Describe a song
type Song struct {
	Id string `json:"id" gorm:"primaryKey;autoIncrement"`
	// Title of the song
	Title string `json:"title"`
	// Lyrics of the song
	Lyrics string `json:"lyrics"`
	// Audio of the song
	AudioUrl string `json:"audio_url"`
	// Release version
	ReleaseVersionId uint           `json:"release_version"`
	ReleaseVersion   ReleaseVersion `gorm:"foreignKey:ReleaseVersionId"`
	// Created at
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
