package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "graduationproject/docs"
	"graduationproject/router"
	"graduationproject/util"
)

//@title 社区便民服务接口
//@version 1.0
//@license.name zwy
//@author 张惟宇
func main() {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")
	url := ginSwagger.URL("http://localhost/swagger/doc.json")
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
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
