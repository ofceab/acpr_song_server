package models

import "time"

// Describe a song
type Song struct {
	Id uint `json:"id" gorm:"primaryKey;autoIncrement"`
	// Title of the song
	Title string `json:"title"`
	// Lyrics of the song
	Lyrics string `json:"lyrics"`
	// Audio of the song
	AudioUrl string `json:"audio_url"`
	// Release version
	ReleaseVersion uint `json:"release_version" gorm:"foreignKey"`
	// Created at
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
