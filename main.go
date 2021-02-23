package main

import (
	"github.com/gin-gonic/gin"
	"graduationproject/router"
)

func main() {
	server := gin.Default()
	router.CInfoRouters(server)
	router.CypageRouters(server)
	router.HkeepingRouters(server)
	router.HrentsaleRouters(server)
	router.WuyefeiRouters(server)
	router.TousubaoxiuRouters(server)
	router.UserRouters(server)
	server.Run(":80")
}
