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
	VodID      int    `gorm:"column:vod_id" json:"vod_id"`
	VodName    string `gorm:"column:vod_name" json:"vod_name"`
	VodStatus  int    `gorm:"column:vod_status" json:"vod_status"`
	VodYear    string `gorm:"column:vod_year" json:"vod_year"`
	VodPlayURL string `gorm:"column:vod_play_url" json:"vod_play_url"`
	VodSerial  string `gorm:"column:vod_serial" json:"vod_serial"`
}
