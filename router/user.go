package router

import (
	"github.com/gin-gonic/gin"
	"graduationproject/model"
	"net/http"
)

func login(context *gin.Context)  {
	var username= context.Query("username")
	var password= context.Query("password")
	register := model.Login(username,password)
	if register {
		context.String(http.StatusOK,"true")
	}else {
		context.String(http.StatusOK,"false")
	}
}

func register(context *gin.Context)  {
	var username=context.Query("username")
	var nickname=context.Query("nickname")
	var password=context.Query("password")
	var houserid=context.Query("houseid")
	login := model.Register(username,password,nickname,houserid)
	if login {
		context.String(http.StatusOK,"true")
	}else {
		context.String(http.StatusOK,"false")
	}
}

func UserRouters(engine *gin.Engine)  {
	engine.GET("/login",login)
	engine.GET("/register",register)
}