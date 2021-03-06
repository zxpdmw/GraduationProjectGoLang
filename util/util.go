package util

import (
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const GetDataFail = "数据获取失败"
const GetDataSuccess = "数据获取成功"
const LoginFail = "登陆失败"
const LoginSuccess = "登陆成功"
const RegisterSuccess = "注册成功"
const RegisterFail = "注册失败"
const InfoFail = "修改用户信息失败"
const InfoSuccess = "修改用户信息成功"
const PasswordSuccess = "修改用户密码成功"
const PasswordFail = "修改用户密码失败"
const ComplainRepairSuccess = "投诉报修成功"
const ComplainRepairFail = "投诉报修失败"
const RecommendNoticeSuccess = "获取推荐公告成功"
const RecommendNoticeFail = "获取推荐公告失败"
const DetailNoticeSuccess = "获取公告详细信息成功"
const DetailNoticeFail = "获取公告详细信息失败"
const PublishNoticeSuccess = "公告发布成功"
const PublishNoticeFail = "公告发布失败"
const RequestFail = "请求失败"
const RequestSuccess = "请求成功"
const PropertyFail = "物业费缴纳失败"
const PropertySuccess = "物业费缴纳成功"
const RentSuccess = "获取全部租房信息成功"
const RentFail = "获取全部租房信息失败"
const SaleFail = "获取全部售房信息失败"
const SaleSuccess = "获取全部售房信息成功"
const RentSaleSuccess = "获取房屋租售信息成功"
const RentSaleFail = "获取房屋租售信息失败"
const DeleteHouseSuccess = "删除房屋信息成功"
const DeleteHouseFail = "删除房屋信息失败"

type CommonResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

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
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
