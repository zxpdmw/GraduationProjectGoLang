package router

import (
	"github.com/gin-gonic/gin"
	"graduationproject/model"
)

func AdminRouters(e *gin.Engine) {
	group := e.Group("/admin")
	{
		group.GET("/", indexPage)
		group.GET("login", adminLogin)
	}
}

func indexPage(c *gin.Context) {
	c.HTML(200, "login.gohtml", nil)
}

func adminLogin(c *gin.Context) {
	query := c.Query("username")
	s := c.Query("password")
	login := model.AdminLogin(query, s)
	if login {
		c.HTML(200, "main.gohtml", gin.H{"message": "登录成功"})
		return
	}
	c.HTML(200, "login.gohtml", gin.H{"message": "密码错误"})
}
