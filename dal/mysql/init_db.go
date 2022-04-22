package mysql

import (
	"acpr_songs_server/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	db, err := openConnectionToDB()
	if err != nil {
		os.Exit(1)
	}

	/// Start migration
	db.AutoMigrate(&models.ReleaseVersion{}, &models.Song{})
	return db
}

func openConnectionToDB() (*gorm.DB, error) {
	_dsn := "obed:ofceab@tcp(127.0.0.1:3306)/acpr_songs?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(_dsn), &gorm.Config{})
	return db, err
}
