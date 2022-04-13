package dal

import (
	"acpr_songs_server/dal"
	"acpr_songs_server/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDataAccessLayer struct {
	dbConnection *gorm.DB
}

// Save song in store
func (s *MysqlDataAccessLayer) SaveSong(s *models.Song) models.Song {
	s.dbConnection.Create(s)
	return s
}

// Fetch songs
func (s *MysqlDataAccessLayer) FetchSongs() []models.Song {
	s.dbConnection.Find(&models.Song{})
}

// Fetch all sounds per version id for fetching release song of a certain `version Id`
FetchSongsPerVersionId() []models.Song
// Remove song from a certain release
DeleteSong(s string) (models.Song, error)

// func New() (dal.IDatabaseAccessLayer, error) {
// 	_dsn := "acpr_user:acpr_user@tcp(127.0.0.1:3306)/acpr_song_db?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(_dsn), &gorm.Config{})

// 	if err != nil {
// 		return nil, err
// 	}
// 	return &MysqlDataAccessLayer{
// 		dbConnection: db,
// 	}, nil
// }

// func (m *MysqlDataAccessLayer) Save(data interface{}) error {
// 	m.dbConnection.Create(&data)
// 	return nil
// }

// func (m *MysqlDataAccessLayer) Retrieve() (interface{}, error) {
// 	m.dbConnection.Find()
// }
