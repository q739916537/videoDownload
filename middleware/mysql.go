package middleware

import (
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	GromMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var mysqlConfig *mysql.Config

type Repo interface {
	GetDb() *gorm.DB
	Close() error
}

type MysqlMgr struct {
	DB *gorm.DB
}

// 用于编译检查
var _ Repo = (*MysqlMgr)(nil)

var defaultRepo Repo

func MysqlDef() Repo {
	return defaultRepo
}

func InitMysql() {
	err := configureYAML()
	if err != nil {
		DefaultLog().Error("configure yaml err:", zap.Error(err))
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", mysqlConfig.User, mysqlConfig.Passwd, "82.157.166.22", 3306, "voide", time.Second*3)

	db, err := gorm.Open(GromMysql.Open(dsn))
	if err != nil {
		DefaultLog().Error("gorm open err:", zap.Error(err))
		return
	}
	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	repo := &MysqlMgr{
		DB: db,
	}
	// 将第一次数据库连接设为默认值
	defaultRepo = repo
}

// UpdateDB update gorm.DB info
func (obj *MysqlMgr) UpdateDB(db *gorm.DB) {
	obj.DB = db
}

// GetDb get gorm.DB info
func (d *MysqlMgr) GetDb() *gorm.DB {
	return d.DB
}

func (d *MysqlMgr) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// IsNotFound ErrRecordNotFound
func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func configureYAML() error {
	//data, err := os.ReadFile("./config/config.yaml")
	//if err != nil {
	//	DefaultLog().Error("Read config.yaml err", zap.Error(err))
	//	return err
	//}
	//
	//err = yaml.Unmarshal(data, mysqlConfig)
	//if err != nil {
	//	DefaultLog().Error("Unmarshal config.yaml err", zap.Error(err))
	//	return err
	//}

	mysqlConfig = &mysql.Config{}
	mysqlConfig.Net = "82.157.166.22:3306"
	mysqlConfig.User = "root"
	mysqlConfig.Passwd = "123456"
	mysqlConfig.DBName = "voide"
	return nil
}
