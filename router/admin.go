package router

import (
	"github.com/gin-gonic/gin"
	"graduationproject/model"
	"graduationproject/util"
)

func AdminRouters(e *gin.Engine) {
	group := e.Group("/admin")
	{
		group.GET("/", indexPage)
		group.GET("login", adminLogin)
		group.GET("/notice", getAllNotice)
		group.GET("/housekeeping", getAllHouseKeeping)
		group.GET("/tsbx", getAllComplainRepair)
	}
}

func indexPage(c *gin.Context) {
	c.HTML(200, "login.gohtml", nil)
}

//@Summary adminLogin
//@Tags 管理员模块
//@Description 管理员登录
//@Param username query string true "管理员账号"
//@Param password query string true "管理员密码"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /admin/login [get]
func adminLogin(c *gin.Context) {
	query := c.Query("username")
	s := c.Query("password")
	err := model.AdminLogin(query, s)
	if err != nil {
		c.HTML(200, "login.gohtml", gin.H{"message": "密码错误"})
		return
	}
	c.HTML(200, "main.gohtml", gin.H{"message": "登录成功"})

}

//@Summary getAllNotice
//@Tags 管理员模块
//@Description 获取全部社区公告信息
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /admin/notice [get]
func getAllNotice(c *gin.Context) {
	data, err := model.GetAllNotice()
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.NoticeFail,
			Data:    nil,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: util.NoticeSuccess,
		Data:    data,
	})

}

//@Summary gtAllComplainRepair
//@Tags 管理员模块
//@Description 后台管理员获取全部的投诉报修
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /admin/tsbx [get]
func getAllComplainRepair(c *gin.Context) {
	data, err := model.GetAllCR()
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: util.RequestSuccess,
		Data:    data,
	})

}

//@Summary getAllHouseKeeping
//@Tags 管理员模块
//@Description 获取全部的家政服务请求
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /admin/housekeeping [get]
func getAllHouseKeeping(c *gin.Context) {
	data, err := model.GetAllHouseKeeping()
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    data,
		})
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: util.RequestSuccess,
		Data:    data,
	})
}
