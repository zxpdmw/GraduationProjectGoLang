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

func init() {
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(config))
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
	util.C.AddFunc("0 0 0 1/1 * ? *", func() {
		model.CronProperty()
	})
	util.C.Start()
}

//@title 社区便民服务接口
//@version 1.0
//@license.name 张惟宇
func main() {
	server := gin.Default()
	fs, _ := template.ParseFS(tmpl, "templates/*.gohtml")
	server.SetHTMLTemplate(fs)
	server.StaticFS("/public", http.FS(static))
	url := ginSwagger.URL(env.GetIp())
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
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
	server.Run()

}
