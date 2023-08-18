package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"videoDownload/types"
)

type MDB struct {
	db *gorm.DB
}

func NewDb() (*MDB, error) {
	mdb := &MDB{}
	db, err := gorm.Open(sqlite.Open("video.db"), &gorm.Config{})
	mdb.db = db
	if err != nil {
		panic("failed to connect database")
	}

	mdb.db.AutoMigrate(&types.VideoRecords{})
	return mdb, err
}

func (d *MDB) InsertVideoRecords(vr *types.VideoRecords) (err error) {
	return d.db.Create(vr).Error
}

func (d *MDB) GetVideoRecords(ids int) (vr *types.VideoRecords, err error) {
	vr = &types.VideoRecords{}
	v := d.db.First(vr, " vod_id = ?", ids)
	return vr, v.Error
}

func (d *MDB) UpdateVideoRecords(vr *types.VideoRecords) (err error) {
	return d.db.Save(vr).Error
	// return d.db.Model(vr).Update("downloaded", complete).Error
}
