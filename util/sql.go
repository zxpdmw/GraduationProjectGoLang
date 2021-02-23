package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func init()  {
	Db, err = gorm.Open(mysql.Open("zxpdmw:Zxpdmw520@tcp(rm-2zeqer8186x8o6hi9vo.mysql.rds.aliyuncs.com:3306)/graduationproject?parseTime=true"), &gorm.Config{})
	//Db,err=sqlx.Connect("mysql","zxpdmw:Zxpdmw520@tcp(rm-2zeqer8186x8o6hi9vo.mysql.rds.aliyuncs.com:3306)/graduationproject?parseTime=true")
	CheckError(err)
}
