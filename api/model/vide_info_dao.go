package model

import (
	"gorm.io/gorm"
	"videoDownload/middleware"
)

// Get 获取
func (obj VideInfo) GetId(id int) (results *VideInfo, err error) {
	err = middleware.MysqlDef().GetDb().Model(&VideInfo{}).Where("id = ?", id).Find(&results).Error
	return
}

// Gets 获取批量结果
func (obj VideInfo) Gets() (results []*VideInfo, err error) {
	err = middleware.MysqlDef().GetDb().Model(&VideInfo{}).Find(&results).Error
	return
}

// Gets 获取批量结果
func (obj VideInfo) Create(results []VideInfo) (err error) {
	if len(results) != 0 {
		err = middleware.MysqlDef().GetDb().Model(&VideInfo{}).Create(&results).Error
		return
	}
	return nil
}

// DeleteById 批量删除
func (obj VideInfo) DeleteById(results []VideInfo) (err error) {
	var delId []int
	for result := range results {
		delId = append(delId, results[result].VodId)
	}
	err = middleware.MysqlDef().GetDb().Model(&VideInfo{}).Where("vod_id IN (?)", delId).Delete(&VideInfo{}).Error
	return err
}

// GetBatchFromVodID 批量查找
func (obj VideInfo) GetBatchFromVodID(vodIDs []int) (results []*VideInfo, err error) {
	err = middleware.MysqlDef().GetDb().Model(&VideInfo{}).Where("`vod_id` IN (?)", vodIDs).Find(&results).Error
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj VideInfo) Count(count *int64) (tx *gorm.DB) {
	return middleware.MysqlDef().GetDb().Model(&VideInfo{}).Count(count)
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromVodID 通过vod_id获取内容
func (obj VideInfo) GetFromVodID(vodID int) (results []*VideInfo, err error) {
	err = middleware.MysqlDef().GetDb().Model(&VideInfo{}).Where("`vod_id` = ?", vodID).Find(&results).Error

	return
}
func (obj VideInfo) GetVodName(vodName string) (results []*VideInfo, err error) {
	err = middleware.MysqlDef().GetDb().Where("`vod_name` LIKE ?", "%"+vodName+"%").Find(&results).Error
	return
}

// GetFromVodName 通过vod_name获取内容
func (obj VideInfo) GetFromVodName(vodName string) (results []*VideInfo, err error) {
	err = middleware.MysqlDef().GetDb().Model(&VideInfo{}).Where("`vod_name` = ?", vodName).Find(&results).Error

	return
}

// GetBatchFromVodName 批量查找
func (obj VideInfo) GetBatchFromVodName(vodNames []string) (results []*VideInfo, err error) {
	err = middleware.MysqlDef().GetDb().Model(&VideInfo{}).Where("`vod_name` IN (?)", vodNames).Find(&results).Error

	return
}

// GetFromVodStatus 通过vod_status获取内容
func (obj VideInfo) GetFromVodStatus(vodStatus int) (results []*VideInfo, err error) {
	err = middleware.MysqlDef().GetDb().Model(&VideInfo{}).Where("`vod_status` = ?", vodStatus).Find(&results).Error

	return
}

// GetBatchFromVodStatus 批量查找
func (obj VideInfo) GetBatchFromVodStatus(vodStatuss []int) (results []*VideInfo, err error) {
	err = middleware.MysqlDef().GetDb().Model(&VideInfo{}).Where("`vod_status` IN (?)", vodStatuss).Find(&results).Error

	return
}

// GetFromVodYear 通过vod_year获取内容
func (obj VideInfo) GetFromVodYear(vodYear string) (results []*VideInfo, err error) {
	err = middleware.MysqlDef().GetDb().Model(&VideInfo{}).Where("`vod_year` = ?", vodYear).Find(&results).Error

	return
}

// GetBatchFromVodYear 批量查找
func (obj VideInfo) GetBatchFromVodYear(vodYears []string) (results []*VideInfo, err error) {
	err = middleware.MysqlDef().GetDb().Model(&VideInfo{}).Where("`vod_year` IN (?)", vodYears).Find(&results).Error

	return
}

// GetFromVodPlayURL 通过vod_play_url获取内容
func (obj VideInfo) GetFromVodPlayURL(vodPlayURL string) (results []*VideInfo, err error) {
	err = middleware.MysqlDef().GetDb().Model(&VideInfo{}).Where("`vod_play_url` = ?", vodPlayURL).Find(&results).Error

	return
}

// GetBatchFromVodPlayURL 批量查找
func (obj VideInfo) GetBatchFromVodPlayURL(vodPlayURLs []string) (results []*VideInfo, err error) {
	err = middleware.MysqlDef().GetDb().Model(&VideInfo{}).Where("`vod_play_url` IN (?)", vodPlayURLs).Find(&results).Error

	return
}

// GetFromVodSerial 通过vod_serial获取内容
func (obj VideInfo) GetFromVodSerial(vodSerial string) (results []*VideInfo, err error) {
	err = middleware.MysqlDef().GetDb().Model(&VideInfo{}).Where("`vod_serial` = ?", vodSerial).Find(&results).Error

	return
}

// GetBatchFromVodSerial 批量查找
func (obj VideInfo) GetBatchFromVodSerial(vodSerials []string) (results []*VideInfo, err error) {
	err = middleware.MysqlDef().GetDb().Model(&VideInfo{}).Where("`vod_serial` IN (?)", vodSerials).Find(&results).Error

	return
}
