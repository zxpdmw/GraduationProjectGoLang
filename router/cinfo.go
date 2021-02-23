package router

import "github.com/gin-gonic/gin"

func CInfoRouters(e *gin.Engine)  {
	e.GET("/firstpageinfotitle",FirstPageInfoTitle)
	e.GET("/getinfobytitle",GetInfoByTitle)
}

func FirstPageInfoTitle(c *gin.Context)  {
	
}

func GetInfoByTitle(c *gin.Context)  {
	
}
