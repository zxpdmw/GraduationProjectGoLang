package main

import (
	"bytes"
	"embed"
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "graduationproject/docs"
	"graduationproject/env"
	"graduationproject/model"
	"graduationproject/router"
	"graduationproject/util"
	"html/template"
	"net/http"
)

//go:embed config/config.yaml
var config []byte

//go:embed templates/*
var tmpl embed.FS

//go:embed assets/*
var static embed.FS

var err error

func init() {
	util.InitLogger()
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(config))
	util.ErrorHandler(err)
	var config = util.DbConfig{
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		Ip:       viper.GetString("db.ip"),
		Port:     viper.GetInt("db.port"),
		Dbname:   viper.GetString("db.dbname"),
	}
	util.MySqlInit(config)
	util.RedisInit()
	util.CronInit()
	err = util.C.AddFunc("0 0 0/1 * * ? ", func() {
		model.CronProperty()
	})
	util.ErrorHandler(err)
	util.C.Start()
}

//@title 社区便民服务接口
//@version 1.0
//@license.name 张惟宇
func main() {
	server := gin.Default()
	fs, err := template.ParseFS(tmpl, "templates/*.gohtml")
	util.ErrorHandler(err)
	server.SetHTMLTemplate(fs)
	server.StaticFS("/public", http.FS(static))
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(env.GetIp())))
	server.GET("/favicon.ico", func(context *gin.Context) {
		file, _ := static.ReadFile("assets/favicon.ico")
		context.Data(200,
			"image/x-icon", file)
	})
	router.AdminRouters(server)
	router.UserRouters(server)
	router.NoticeRouters(server)
	router.HousekeepingRouters(server)
	router.HouseRentingRouters(server)
	router.PropertyRouters(server)
	router.ComplainRepairRouters(server)
	server.NoRoute(func(context *gin.Context) {
		context.String(200, "404 page not found")
	})
	err = server.Run()
	util.ErrorHandler(err)
}
