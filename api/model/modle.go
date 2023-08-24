/*
--------------------------------------------------
 @FileName: modle.go
 @Author:  yuanzibo@firecloud.ai
 @Company:  Firecloud
 @CreatedAt: 2023-08-22 14:33:25
---------------------说明--------------------------

---------------------------------------------------
*/

package model

type ResourceResp struct {
	Code  int        `json:"code" `
	Total int        `json:"total"`
	List  []VideInfo `json:"list"`
}

type VideInfo struct {
	//VodID      int    `gorm:"column:vod_id" json:"vod_id"`
	//VodName    string `gorm:"column:vod_name" json:"vod_name"`
	//VodStatus  int    `gorm:"column:vod_status" json:"vod_status"`
	//VodYear    string `gorm:"column:vod_year" json:"vod_year"`
	//VodPlayURL string `gorm:"column:vod_play_url" json:"vod_play_url"`
	//VodSerial  string `gorm:"column:vod_serial" json:"vod_serial"`

	VodId       int    `json:"vod_id" gorm:"column:vod_id"`
	VodName     string `json:"vod_name" gorm:"column:vod_name"`
	TypeId      int    `json:"type_id" gorm:"column:type_id"`
	TypeName    string `json:"type_name" gorm:"column:type_name"`
	VodEn       string `json:"vod_en" gorm:"column:vod_en"`
	VodTime     string `json:"vod_time" gorm:"column:vod_time"`
	VodRemarks  string `json:"vod_remarks" gorm:"column:vod_remarks"`
	VodPlayFrom string `json:"vod_play_from" gorm:"column:vod_play_from"`
}

func (videInfo VideInfo) TableName() string {
	return "vide_info"
}

func NewVideInfo() *VideInfo {
	return new(VideInfo)
}
