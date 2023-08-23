/*
--------------------------------------------------
 @FileName: readFile.go
 @Author:  yuanzibo@firecloud.ai
 @Company:  Firecloud
 @CreatedAt: 2023-08-18 16:54:30
---------------------说明--------------------------

---------------------------------------------------
*/

package server

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	GromMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var mysqlConfig *mysql.Config

var server Server

type Server struct {
	zap   *zap.Logger
	mysql *gorm.DB
}

func DefaultLogger() *zap.Logger {
	if server.zap == nil {
		logger := zap.NewNop()
		return logger
	}
	return server.zap
}

func InitMiddleWare() {
	initLogger()
	initMysql()
}

func initLogger() *zap.Logger {
	logger, err2 := zap.NewProduction()
	if err2 != nil {
		fmt.Println("init zap err:", err2)
		return nil
	}
	return logger
}

func configureYAML() error {
	data, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		DefaultLogger().Error("Read config.yaml err", zap.Error(err))
		return err
	}

	err = yaml.Unmarshal(data, mysqlConfig)
	if err != nil {
		DefaultLogger().Error("Unmarshal config.yaml err", zap.Error(err))
		return err
	}
	return nil
}

func initMysql() {
	err := configureYAML()
	if err != nil {
		DefaultLogger().Error("configure yaml err:", zap.Error(err))
		return
	}

	db, err := gorm.Open(GromMysql.Open(mysqlConfig.FormatDSN()))
	if err != nil {
		DefaultLogger().Error("gorm open err:", zap.Error(err))
		return
	}
	db.Set("gorm:table_options", "CHARSET=utf8mb4")
	server.mysql = db
}

// Close 关闭mysql
func Close() error {
	/// 关闭mysql
	if server.mysql != nil {
		db, err := server.mysql.DB()
		if err != nil {
			DefaultLogger().Error("mysql close err:", zap.Error(err))
			return err
		}
		return db.Close()

	}
	return nil

}
