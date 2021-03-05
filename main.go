package main

import (
	"github.com/gin-gonic/gin"
	"graduationproject/router"
	"graduationproject/util"
)

func main() {
	server := gin.Default()
	server.StaticFile("/favicon", "https://zwyblog.oss-cn-beijing.aliyuncs.com/touxiang.jpg")
	router.UserRouters(server)
	router.NoticeRouters(server)
	router.HousekeepingRouters(server)
	router.HouseRentingRouters(server)
	router.PropertyRouters(server)
	router.ComplainRepairRouters(server)
	err := server.Run(":80")
	util.CheckError(err)
}
