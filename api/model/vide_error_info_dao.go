package model

import (
	"videoDownload/middleware"
)

// Get 获取
func (videInfo VideErrorInfo) GetUrl(url string) (results *VideErrorInfo, err error) {
	err = middleware.MysqlDef().GetDb().Model(&VideErrorInfo{}).Where("vod_url = ?", url).Find(&results).Error
	return
}

// Gets 获取批量结果
func (videInfo VideErrorInfo) Gets() (results []*VideErrorInfo, err error) {
	err = middleware.MysqlDef().GetDb().Model(&VideErrorInfo{}).Find(&results).Error
	return
}

// Create 批量创建
func (videInfo VideErrorInfo) Create(results []VideErrorInfo) (err error) {
	if len(results) != 0 {
		err = middleware.MysqlDef().GetDb().Model(&VideErrorInfo{}).Create(&results).Error
		return
	}
	return nil
}

// CreateOne 创建唯一
func (videInfo VideErrorInfo) CreateOne(results VideErrorInfo) (err error) {
	err = middleware.MysqlDef().GetDb().Model(&VideErrorInfo{}).Create(&results).Error
	return nil
}

func (videInfo VideErrorInfo) DeleteByUrl(url string) (err error) {
	err = middleware.MysqlDef().GetDb().Model(&VideErrorInfo{}).Where("vod_url = ?", url).Delete(&VideErrorInfo{}).Error
	return err
}
