package router

import (
	"github.com/gin-gonic/gin"
	"graduationproject/model"
	"graduationproject/util"
	"net/http"
	"strings"
)

//@Summary userLogin
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
	exist, _ := model.CheckUserExist(username)
	if exist == 0 {
		context.JSON(200, util.Response{
			Code:    5551,
			Message: util.UserNotExist,
			Data:    nil,
		})
		return
	}
	c, err := model.Login(username, password)
	if err != nil {
		context.JSON(http.StatusOK, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
		return
	}

	if !strings.EqualFold(password, c.Password) {
		context.JSON(http.StatusOK, util.Response{
			Code:    5552,
			Message: util.PasswordError,
			Data:    nil,
		})
		return
	}
	context.JSON(http.StatusOK, util.Response{
		Code:    666,
		Message: util.LoginSuccess,
		Data:    c,
	})
}

//@Summary userRegister
//@Description 用户输入在账号，密码，并绑定自己的房屋ID，昵称
//@Tags 用户模块
//@Param userRegister body model.UserRegister true "用户注册结构体"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /user/register [get]
func userRegister(c *gin.Context) {
	var ur model.UserRegister
	if err := c.ShouldBindJSON(&ur); err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
		return
	}
	exist, err := model.CheckUserExist(ur.Username)
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
		return
	}
	if exist == 1 {
		c.JSON(200, util.Response{
			Code:    5550,
			Message: util.UserIsExist,
			Data:    nil,
		})
		return
	}
	err = model.Register(ur)
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
		return
	}
	model.UserBindHouseID(model.Property{
		HouseID: ur.HouseId,
		Balance: 500.0,
	})
	c.JSON(http.StatusOK, util.Response{
		Code:    666,
		Message: util.RegisterSuccess,
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
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
		return
	}
	err := model.EditInfo(u)
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    555,
		Message: util.InfoSuccess,
		Data:    nil,
	})

}

//@Summary editUserPassword
//@Description 修改用户密码
//@Tags 用户模块
//@Param username query string true "账号"
//@Param password query string true "密码"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /user/editpassword [post]
func editUserPassword(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	err := model.EditPassword(username, password)
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: util.PasswordSuccess,
		Data:    nil,
	})

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
	info, err := model.GetInfo(query)
	if err != nil {
		c.JSON(http.StatusOK, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    info,
		})
		return
	}
	c.JSON(http.StatusOK, util.Response{
		Code:    666,
		Message: util.GetDataSuccess,
		Data:    info,
	})
}

//@Summary getHouseId
//@Tags 用户模块
//@Description 根据username获取房屋号
//@Param username query string true "账号"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /user/houseid [get]
func getHouseId(c *gin.Context) {
	query := c.Query("username")
	id, err := model.GetHouseId(query)
	if err != nil {
		c.JSON(http.StatusOK, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: util.RequestSuccess,
		Data:    id,
	})
}

//@Summary editHouseId
//@Tags 用户模块
//@Description 修改用户的房屋号
//@Param username query string true "账号"
//@Param house query string true "房屋号"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /user/edithouseid [get]
func editHouseId(c *gin.Context) {
	username := c.Query("username")
	houseID := c.Query("house")
	err := model.EditHouseID(username, houseID)
	if err != nil {
		c.JSON(http.StatusOK, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: util.RequestSuccess,
		Data:    nil,
	})

}

//@Summary editNickname
//@Tags 用户模块
//@Description 修改用户的昵称
//@Param username query string true "账号"
//@Param nickname query string true "昵称"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /user/editnickname [get]
func editNickname(c *gin.Context) {
	username := c.Query("username")
	nickname := c.Query("nickname")
	err := model.EditNickname(username, nickname)
	if err != nil {
		c.JSON(http.StatusOK, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: util.RequestSuccess,
		Data:    nil,
	})
}

//@Summary editAddress
//@Tags 用户模块
//@Description 修改用户的地址
//@Param username query string true "账号"
//@Param address query string true "地址"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /user/editaddress [get]
func editAddress(c *gin.Context) {
	username := c.Query("username")
	houseID := c.Query("address")
	err := model.EditAddress(username, houseID)
	if err != nil {
		c.JSON(http.StatusOK, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: util.RequestSuccess,
		Data:    nil,
	})

}

//@Summary editPhone
//@Tags 用户模块
//@Description 修改用户的电话
//@Param username query string true "账号"
//@Param phone query string true "电话"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /user/editphone [get]
func editPhone(c *gin.Context) {
	username := c.Query("username")
	houseID := c.Query("phone")
	err := model.EditPhone(username, houseID)
	if err != nil {
		c.JSON(http.StatusOK, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: util.RequestSuccess,
		Data:    nil,
	})
}

func UserRouters(engine *gin.Engine) {
	group := engine.Group("/user")
	{
		group.GET("/get", getInfo)
		group.GET("/login", userLogin)
		group.POST("/register", userRegister)
		group.POST("/info", editUserInfo)
		group.GET("/houseid", getHouseId)
		group.GET("/editpassword", editUserPassword)
		group.GET("/edithouseid", editHouseId)
		group.GET("/editaddress", editAddress)
		group.GET("/editphone", editPhone)
		group.GET("/editnickname", editNickname)
	}
}
