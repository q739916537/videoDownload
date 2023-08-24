package model

// VideInfoColumns get sql column name.获取数据库列名
var VideInfoColumns = struct {
	VodID      string
	VodName    string
	VodStatus  string
	VodYear    string
	VodPlayURL string
	VodSerial  string
}{
	VodID:      "vod_id",
	VodName:    "vod_name",
	VodStatus:  "vod_status",
	VodYear:    "vod_year",
	VodPlayURL: "vod_play_url",
	VodSerial:  "vod_serial",
}
