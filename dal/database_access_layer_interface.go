package dal

import "acpr_songs_server/models"

// import "acpr_songs_server/models"

// type IDatabaseAccessLayer interface {
// 	IDataSaver
// 	IDataRetriever
// }

// type Data interface {
// 	models.ReleaseVersion | []models.ReleaseVersion | models.Song | []models.Song
// }

// // Allow to Save Data in store
// type IDataSaver interface {
// 	Save(d interface{}) error
// }

// type IDataRetriever interface {
// 	Retrieve() (interface{}, error)
// }

type ISongDatabaseAccessLayer interface {
	// Save song in store
	SaveSong(s *models.Song) models.Song
	// Fetch songs
	FetchSongs() []models.Song
	// Fetch all sounds per version id for fetching release song of a certain `version Id`
	FetchSongsPerVersionId() []models.Song
	// Remove song from a certain release
	DeleteSong(s string) (models.Song, error)
}
