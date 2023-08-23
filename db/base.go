package db

import (
	"errors"
	"gorm.io/gorm"
)

// IsNotFound ErrRecordNotFound
func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

type MysqlMgr struct {
	DB *gorm.DB
}

// GetDB get gorm.DB info
func (obj *MysqlMgr) GetDB() *gorm.DB {
	return obj.DB
}

// UpdateDB update gorm.DB info
func (obj *MysqlMgr) UpdateDB(db *gorm.DB) {
	obj.DB = db
}

// New new gorm.新gorm,重置条件
func (obj *MysqlMgr) New() {
	obj.DB = obj.NewDB()
	obj.DB.Update()
}

// NewDB new gorm.新gorm
func (obj *MysqlMgr) NewDB() *gorm.DB {
	return obj.DB.Session(&gorm.Session{NewDB: true, Context: obj.ctx})
}
