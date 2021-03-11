package router

import (
	"github.com/gin-gonic/gin"
	"graduationproject/model"
	"graduationproject/util"
	"net/http"
)

//@Summary 用户登录
//@Description 用户使用账户和密码登录
//@Tags 用户模块
//@Param username query string true "用户账户"
//@Param password query string true "用户密码"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /user/login [get]
func userLogin(context *gin.Context) {
	var username = context.Query("username")
	var password = context.Query("password")
	register := model.Login(username, password)
	if register {
		context.JSON(http.StatusOK, util.Response{
			Code:    666,
			Message: util.LoginSuccess,
			Data:    nil,
		})
		return
	}
	context.JSON(http.StatusOK, util.Response{
		Code:    555,
		Message: util.LoginFail,
		Data:    nil,
	})
}

//@Summary 用户注册
//@Description 用户输入在账号，密码，并绑定自己的房屋ID，昵称
//@Tags 用户模块
//@Param username query string true "用户账户"
//@Param password query string true "用户密码"
//@Param nickname query string true "用户昵称"
//@Param houseId query string true "房屋ID"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /user/register [get]
func userRegister(context *gin.Context) {
	var username = context.Query("username")
	var nickname = context.Query("nickname")
	var password = context.Query("password")
	var houseId = context.Query("houseId")
	login := model.Register(username, password, nickname, houseId)
	if login {
		context.JSON(http.StatusOK, util.Response{
			Code:    666,
			Message: util.RegisterSuccess,
			Data:    nil,
		})
		return
	}
	context.JSON(http.StatusOK, util.Response{
		Code:    555,
		Message: util.RegisterFail,
		Data:    nil,
	})

}

//@Summary editUserInfo
//@Description 修改用户信息
//@Tags 用户模块
//@Param user body model.User true "User结构体包含要修改的字段"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /user/info [post]
func editUserInfo(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		util.CheckError(err)
	} else {
		info := model.EditInfo(u)
		if info {
			c.JSON(200, util.Response{
				Code:    666,
				Message: util.InfoSuccess,
				Data:    nil,
			})
			return
		}
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.InfoFail,
			Data:    nil,
		})
	}
}

//@Summary editUserPassword
//@Description 修改用户密码
//@Tags 用户模块
//@Param user body model.User true "User结构体只包含username和password"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /user/password [post]
func editUserPassword(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		util.CheckError(err)
	} else {
		password := model.EditPassword(u)
		if password {
			c.JSON(200, util.Response{
				Code:    666,
				Message: util.PasswordSuccess,
				Data:    nil,
			})
			return
		}
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.PasswordFail,
			Data:    nil,
		})
	}
}

//@Summary getUserInfo
//@Description 获取用户信息根据username
//@Tags 用户模块
//@Param username query string true "用户账户"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /user/get [get]
func getInfo(c *gin.Context) {
	query := c.Query("username")
	info, userInfo := model.GetInfo(query)
	if info {
		c.JSON(http.StatusOK, util.Response{
			Code:    111,
			Message: util.GetDataSuccess,
			Data:    userInfo,
		})
		return
	}
	c.JSON(http.StatusOK, util.Response{
		Code:    000,
		Message: util.GetDataFail,
		Data:    userInfo,
	})
}

func UserRouters(engine *gin.Engine) {
	group := engine.Group("/user")
	{
		group.GET("/get", getInfo)
		group.GET("/login", userLogin)
		group.GET("/register", userRegister)
		group.POST("/info", editUserInfo)
		group.POST("/password", editUserPassword)
	}
}
