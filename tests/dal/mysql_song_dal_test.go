package dal_test

import (
	"acpr_songs_server/dal/dal_interfaces"
	"acpr_songs_server/dal/mysql"
	dataformat "acpr_songs_server/data_format"
	"fmt"
	"testing"
)

func TestSaveSong(t *testing.T) {
	_db := mysql.InitDb()
	_releaseVersionDal := dal_interfaces.MysqlReleaseVersionDataAccessLayer{DbConnection: _db}
	_d := dal_interfaces.MysqlSongDataAccessLayer{DbConnection: _db, ReleaseVersionDal: &_releaseVersionDal}

	_tempSongData := dataformat.CreateSong{Title: "Alleluia", Lyrics: "Hosana", AudioUrl: ""}
	_r, err := _d.SaveSong(&_tempSongData, 21)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(_r)
}
