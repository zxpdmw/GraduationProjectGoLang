package router

import (
	"github.com/gin-gonic/gin"
	"graduationproject/model"
	"graduationproject/util"
	"net/http"
)

func login(context *gin.Context) {
	var username = context.Query("username")
	var password = context.Query("password")
	register := model.Login(username, password)
	if register {
		context.String(http.StatusOK, "true")
	} else {
		context.String(http.StatusOK, "false")
	}
}

func register(context *gin.Context) {
	var username = context.Query("username")
	var nickname = context.Query("nickname")
	var password = context.Query("password")
	var houserid = context.Query("houseid")
	login := model.Register(username, password, nickname, houserid)
	if login {
		context.String(http.StatusOK, "true")
	} else {
		context.String(http.StatusOK, "false")
	}
}

func editInfo(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		util.CheckError(err)
	} else {
		update := util.Db.Model(&model.User{}).Where("username=?", u.Username).Update("nickname", u.Nickname).Update("address", u.Address).Update("phone", u.Phone).Update("house_id", u.HouseId)
		if update.RowsAffected == 1 {
			c.String(200, "true")
		} else {
			c.String(200, "false")
		}
	}
}

func editPassword(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		util.CheckError(err)
	} else {
		update := util.Db.Model(&model.User{}).Where("username=?", u.Username).Update("password", u.Password)
		if update.RowsAffected == 1 {
			c.String(200, "true")
		} else {
			c.String(200, "false")
		}
	}

}

func getInfo(c *gin.Context) {

}

func UserRouters(engine *gin.Engine) {
	group := engine.Group("/user")
	{
		group.GET("/login", login)
		group.GET("/register", register)
		group.POST("/editinfo", editInfo)
		group.POST("/editpassword", editPassword)
	}

}
