package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "graduationproject/docs"
	"graduationproject/env"
	"graduationproject/router"
	"graduationproject/util"
	"net/http"
)

//@title 社区便民服务接口
//@version 1.0
//@license.name zwy
//@author 张惟宇
func main() {
	server := gin.Default()
	path := env.GetTemplatePath()
	server.LoadHTMLGlob(path)
	url := ginSwagger.URL(env.GetIp())
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	server.Static("/static", "./static")
	server.GET("/favicon.ico", func(c *gin.Context) {
		response, err := http.Get("https://cdn.jsdelivr.net/gh/zxpdmw/pictureBed/img/touxiang.jpg")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		defer reader.Close()
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, nil)
	})
	//server.StaticFile("/favicon.ico","static/img/favicon.ico")
	router.AdminRouters(server)
	router.UserRouters(server)
	router.NoticeRouters(server)
	router.HousekeepingRouters(server)
	router.HouseRentingRouters(server)
	router.PropertyRouters(server)
	router.ComplainRepairRouters(server)
	err := server.Run(":80")
	util.CheckError(err)
}
