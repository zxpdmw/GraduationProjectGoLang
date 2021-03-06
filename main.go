package main

import (
	"github.com/gin-gonic/gin"
	"graduationproject/router"
	"graduationproject/util"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")
	server.StaticFile("/static", "./static")
	server.StaticFile("/favicon", "./favicon.ico")
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
