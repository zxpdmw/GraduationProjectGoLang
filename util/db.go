package util

import (
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error
var Rdb *redis.Client

func init() {
	Db, err = gorm.Open(mysql.Open("zxpdmw:Zxpdmw520@tcp(rm-2zeqer8186x8o6hi9vo.mysql.rds.aliyuncs.com:3306)/graduationproject?parseTime=true&loc=Asia%2fShanghai"), &gorm.Config{})
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "39.96.113.190:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err3 := Rdb.Ping().Result()
	CheckError(err3)
	//Db,err=sqlx.Connect("mysql","zxpdmw:Zxpdmw520@tcp(rm-2zeqer8186x8o6hi9vo.mysql.rds.aliyuncs.com:3306)/graduationproject?parseTime=true")
	CheckError(err)
}
