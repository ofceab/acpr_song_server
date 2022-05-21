package dataformat

// Data to send to create Song
type CreateSong struct {
	// Title of the song
	Title string `json:"title" binding:"required"`
	// Lyrics of the song
	Lyrics string `json:"lyrics" binding:"required"`
	// Audio of the song
	AudioUrl string `json:"audio_url"`
}

type UpdateSong struct {
	CreateSong
	// Unique song Id
	SongUniqueId string `json:"song_unique_id" binding:"required"`
}
