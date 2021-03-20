package router

import (
	"github.com/gin-gonic/gin"
	"graduationproject/model"
	"graduationproject/util"
	"strconv"
)

func AdminRouters(e *gin.Engine) {
	group := e.Group("/admin")
	{
		group.GET("/", indexPage)
		group.GET("login", adminLogin)
		group.GET("/notice", getAllNotice)
		group.GET("/housekeeping", getAllHouseKeeping)
		group.GET("/complainrepair", getAllComplainRepair)
		group.GET("/editcr", editComplainRepairStatus)
		group.GET("/edithk", editHouseKeepingStatus)
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
//@Router /admin/complainrepair [get]
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

//@Summary editComplainRepairStatus
//@Tags 管理员模块
//@Description 修改投诉报修的状态
//@Param id query string true "投诉报修在数据库的主键id"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /admin/editcr [get]
func editComplainRepairStatus(c *gin.Context) {
	query := c.Query("id")
	atoi, _ := strconv.Atoi(query)
	err := model.EditComplainRepairStatus(atoi)
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: "投诉报修状态修改成功",
			Data:    nil,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: "投诉报修状态修改失败",
		Data:    nil,
	})
}

//@Summary editHouseKeeping
//@Tags 管理员模块
//@Description 修改家政服务的服务状态
//@Param id query string true "家政服务在数据库中的主键id"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /admin/edithk [get]
func editHouseKeepingStatus(c *gin.Context) {
	query := c.Query("id")
	atoi, _ := strconv.Atoi(query)
	err := model.EditHouseKeepingStatus(atoi)
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: "家政服务状态修改失败",
			Data:    nil,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: "家政服务状态修改成功",
		Data:    nil,
	})

}
