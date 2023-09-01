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

import "time"

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

type VideErrorInfo struct {
	VodUrl     string    `json:"vod_url" gorm:"column:vod_url"`
	Method     string    `json:"method" gorm:"column:method"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	Error      string    `json:"error" gorm:"column:error"`
	Lx         string    `json:"lx" gorm:"column:lx"`
}

func (videInfo VideErrorInfo) TableName() string {
	return "vide_error_info"
}

func NewVideErrorInfo() *VideErrorInfo {
	return new(VideErrorInfo)
}

// VideUrlInfo  url具体信息存储
type VideUrlInfo struct {
	VodID         int    `gorm:"column:vod_id" json:"vod_id"`
	VodName       string `gorm:"column:vod_name" json:"vod_name"`
	VodStatus     int    `gorm:"column:vod_status" json:"vod_status"`
	VodYear       string `gorm:"column:vod_year" json:"vod_year"`
	VodPlayURL    string `gorm:"column:vod_play_url" json:"vod_play_url"`
	VodLocalPath  string `gorm:"column:vod_local_path" json:"vod_local_path"`
	VodSerial     string `gorm:"column:vod_serial" json:"vod_serial"`
	VodDownStatus string `gorm:"column:vod_down_status" json:"vod_down_status"`
}

func (videInfo VideUrlInfo) TableName() string {
	return "vide_url_info"
}

func NewVideUrlInfo() *VideErrorInfo {
	return new(VideErrorInfo)
}
