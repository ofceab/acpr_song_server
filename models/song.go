package models

// Describe a song
type Song struct {
	Id string `json:"id"`
	// Title of the song
	Title string `json:"title"`
	// Lyrics of the song
	Lyrics string `json:"lyrics"`
	// Audio of the song
	AudioUrl string `json:"audio_url"`
	// Release version
	ReleaseVersion string `json:"release_version"`
}
