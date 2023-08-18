package types

import (
	"gorm.io/gorm"
)

type (
	ResourceResp struct {
		Code  int               `json:"code"`
		Total int               `json:"total"`
		List  []VideoDetailList `json:"list"`
	}
	VideoDetailList struct {
		VodID      int    `json:"vod_id" gorm:"uniqueIndex:idx_name,sort:desc"`
		VodName    string `json:"vod_name"`
		VodStatus  int    `json:"vod_status"`
		VodYear    string `json:"vod_year"`
		VodPlayURL string `json:"vod_play_url"`
		VodSerial  string `json:"vod_serial"`
	}
)

type VideoRecords struct {
	gorm.Model
	VideoDetailList
	Downloaded int
}
