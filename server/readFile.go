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
	"github.com/go-sql-driver/mysql"
	"videoDownload/middleware"
)

var server Server

type Server struct {
	mysqlConfig *mysql.Config
}

func InitMiddleWare() {
	middleware.InitLog()
	middleware.InitMysql()

}
