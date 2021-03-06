package router

import (
	"github.com/gin-gonic/gin"
	"graduationproject/model"
	"graduationproject/util"
	"net/http"
)

func userLogin(context *gin.Context) {
	var username = context.Query("username")
	var password = context.Query("password")
	register := model.Login(username, password)
	if register {
		context.JSON(http.StatusOK, gin.H{
			"code":    666,
			"message": util.LoginSuccess,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    555,
		"message": util.LoginFail,
	})
}

func userRegister(context *gin.Context) {
	var username = context.Query("username")
	var nickname = context.Query("nickname")
	var password = context.Query("password")
	var houseId = context.Query("houseId")
	login := model.Register(username, password, nickname, houseId)
	if login {
		context.JSON(http.StatusOK, gin.H{
			"code":    666,
			"message": util.RegisterSuccess,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    555,
		"message": util.RegisterFail,
	})

}

func editUserInfo(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		util.CheckError(err)
	} else {
		info := model.EditInfo(u)
		if info {
			c.JSON(200, gin.H{
				"code":    666,
				"message": util.InfoSuccess,
			})
			return
		}
		c.JSON(200, gin.H{
			"code":    555,
			"message": util.InfoFail,
		})
	}
}

func editUserPassword(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		util.CheckError(err)
	} else {
		password := model.EditPassword(u)
		if password {
			c.JSON(200, gin.H{
				"code":    666,
				"message": util.PasswordSuccess,
			})
			return
		}
		c.JSON(200, gin.H{
			"code":    555,
			"message": util.PasswordFail,
		})
	}
}

func getInfo(c *gin.Context) {
	query := c.Query("username")
	info, userInfo := model.GetInfo(query)
	if info {
		c.JSON(http.StatusOK, util.CommonResult{
			Code:    111,
			Message: util.GetDataSuccess,
			Data:    userInfo,
		})
		return
	}
	c.JSON(http.StatusOK, util.CommonResult{
		Code:    000,
		Message: util.GetDataFail,
		Data:    userInfo,
	})
}

func UserRouters(engine *gin.Engine) {
	group := engine.Group("/user")
	{
		group.GET("/getInfo", getInfo)
		group.GET("/login", userLogin)
		group.GET("/register", userRegister)
		group.POST("/info", editUserInfo)
		group.POST("/password", editUserPassword)
	}
}
