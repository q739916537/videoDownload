package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"videoDownload/db"
)

func NewVideInfoMgr() string {
	return "vide_info"
}

// VideInfoMgr open func
func VideInfoMgr(db *gorm.DB) *_VideInfoMgr {
	if db == nil {
		panic(fmt.Errorf("VideInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_VideInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("vide_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}

}

// GetTableName get sql table name.获取数据库名字
func (obj *_VideInfoMgr) GetTableName() string {
	db.MysqlMgr{}
	return "vide_info"
}

// Reset 重置gorm会话
func (obj *_VideInfoMgr) Reset() *_VideInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_VideInfoMgr) Get() (result VideInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_VideInfoMgr) Gets() (results []*VideInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideInfo{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_VideInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(VideInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithVodID vod_id获取
func (obj *_VideInfoMgr) WithVodID(vodID int) Option {
	return optionFunc(func(o *options) { o.query["vod_id"] = vodID })
}

// WithVodName vod_name获取
func (obj *_VideInfoMgr) WithVodName(vodName string) Option {
	return optionFunc(func(o *options) { o.query["vod_name"] = vodName })
}

// WithVodStatus vod_status获取
func (obj *_VideInfoMgr) WithVodStatus(vodStatus int) Option {
	return optionFunc(func(o *options) { o.query["vod_status"] = vodStatus })
}

// WithVodYear vod_year获取
func (obj *_VideInfoMgr) WithVodYear(vodYear string) Option {
	return optionFunc(func(o *options) { o.query["vod_year"] = vodYear })
}

// WithVodPlayURL vod_play_url获取
func (obj *_VideInfoMgr) WithVodPlayURL(vodPlayURL string) Option {
	return optionFunc(func(o *options) { o.query["vod_play_url"] = vodPlayURL })
}

// WithVodSerial vod_serial获取
func (obj *_VideInfoMgr) WithVodSerial(vodSerial string) Option {
	return optionFunc(func(o *options) { o.query["vod_serial"] = vodSerial })
}

// GetByOption 功能选项模式获取
func (obj *_VideInfoMgr) GetByOption(opts ...Option) (result VideInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(VideInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_VideInfoMgr) GetByOptions(opts ...Option) (results []*VideInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(VideInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_VideInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]VideInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(VideInfo{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromVodID 通过vod_id获取内容
func (obj *_VideInfoMgr) GetFromVodID(vodID int) (results []*VideInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideInfo{}).Where("`vod_id` = ?", vodID).Find(&results).Error

	return
}

// GetBatchFromVodID 批量查找
func (obj *_VideInfoMgr) GetBatchFromVodID(vodIDs []int) (results []*VideInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideInfo{}).Where("`vod_id` IN (?)", vodIDs).Find(&results).Error

	return
}

// GetFromVodName 通过vod_name获取内容
func (obj *_VideInfoMgr) GetFromVodName(vodName string) (results []*VideInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideInfo{}).Where("`vod_name` = ?", vodName).Find(&results).Error

	return
}

// GetBatchFromVodName 批量查找
func (obj *_VideInfoMgr) GetBatchFromVodName(vodNames []string) (results []*VideInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideInfo{}).Where("`vod_name` IN (?)", vodNames).Find(&results).Error

	return
}

// GetFromVodStatus 通过vod_status获取内容
func (obj *_VideInfoMgr) GetFromVodStatus(vodStatus int) (results []*VideInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideInfo{}).Where("`vod_status` = ?", vodStatus).Find(&results).Error

	return
}

// GetBatchFromVodStatus 批量查找
func (obj *_VideInfoMgr) GetBatchFromVodStatus(vodStatuss []int) (results []*VideInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideInfo{}).Where("`vod_status` IN (?)", vodStatuss).Find(&results).Error

	return
}

// GetFromVodYear 通过vod_year获取内容
func (obj *_VideInfoMgr) GetFromVodYear(vodYear string) (results []*VideInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideInfo{}).Where("`vod_year` = ?", vodYear).Find(&results).Error

	return
}

// GetBatchFromVodYear 批量查找
func (obj *_VideInfoMgr) GetBatchFromVodYear(vodYears []string) (results []*VideInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideInfo{}).Where("`vod_year` IN (?)", vodYears).Find(&results).Error

	return
}

// GetFromVodPlayURL 通过vod_play_url获取内容
func (obj *_VideInfoMgr) GetFromVodPlayURL(vodPlayURL string) (results []*VideInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideInfo{}).Where("`vod_play_url` = ?", vodPlayURL).Find(&results).Error

	return
}

// GetBatchFromVodPlayURL 批量查找
func (obj *_VideInfoMgr) GetBatchFromVodPlayURL(vodPlayURLs []string) (results []*VideInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideInfo{}).Where("`vod_play_url` IN (?)", vodPlayURLs).Find(&results).Error

	return
}

// GetFromVodSerial 通过vod_serial获取内容
func (obj *_VideInfoMgr) GetFromVodSerial(vodSerial string) (results []*VideInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideInfo{}).Where("`vod_serial` = ?", vodSerial).Find(&results).Error

	return
}

// GetBatchFromVodSerial 批量查找
func (obj *_VideInfoMgr) GetBatchFromVodSerial(vodSerials []string) (results []*VideInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideInfo{}).Where("`vod_serial` IN (?)", vodSerials).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
