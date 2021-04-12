package util

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

const GetDataFail = "数据获取失败"
const GetDataSuccess = "数据获取成功"
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
const NoticeSuccess = "获取全部公告成功"
const NoticeFail = "获取全部公告失败"

const RegisterSuccess = "注册成功"
const LoginSuccess = "登陆成功"
const UserIsExist = "用户已存在!"
const UserNotExist = "用户不存在,请检查用户名!"
const PasswordError = "密码错误,请检查后重新输入!"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type DbConfig struct {
	Username string
	Password string
	Ip       string
	Port     int
	Dbname   string
}

var C *cron.Cron
var Db *gorm.DB
var err error
var Rdb *redis.Client
var Logger *zap.Logger

func MySqlInit(config DbConfig) {

	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.Username, config.Password, config.Ip, config.Port, config.Dbname)
	Db, err = gorm.Open(mysql.Open(conn), &gorm.Config{})
	Db = Db.Debug()
	if err != nil {
		Logger.Error("get mysql conn is happen error", zap.Error(err))
	} else {
		Logger.Info("get mysql conn success")
	}
}

func CronInit() {
	C = cron.New()
}

func RedisInit() {
	ctx := context.Background()
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "39.96.113.190:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err = Rdb.Ping(ctx).Result()
	if err != nil {
		Logger.Error("get redis conn is happen error", zap.Error(err))
	} else {
		Logger.Info("get redis conn success")
	}
}

func ErrorHandler(err error) {
	if err != nil {
		Logger.Error("happer error", zap.Error(err))
	}
}

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	Logger = zap.New(core, zap.AddCaller())
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	//file, _ := os.Create("./log/dev.log")
	return zapcore.AddSync(os.Stdout)
}
