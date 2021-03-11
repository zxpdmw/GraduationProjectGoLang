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
		group.GET("/cr", getAllCR)
		group.GET("/notice", getAllNotice)
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
	login := model.AdminLogin(query, s)
	if login {
		c.HTML(200, "main.gohtml", gin.H{"message": "登录成功"})
		return
	}
	c.HTML(200, "login.gohtml", gin.H{"message": "密码错误"})
}

//@Summary getAllComplainRepair
//@Tags 管理员模块
//@Description 获取全部的投诉报修信息
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /admin/cr [get]
func getAllCR(c *gin.Context) {
	cr, b := model.GetAllCR()
	if b {
		c.JSON(200, util.Response{
			Code:    666,
			Message: util.ComplainRepairSuccess,
			Data:    cr,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    555,
		Message: util.ComplainRepairFail,
		Data:    nil,
	})
}

//@Summary getAllNotice
//@Tags 管理员模块
//@Description 获取全部社区公告信息
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /admin/notice [get]
func getAllNotice(c *gin.Context) {
	notice, b := model.GetAllNotice()
	if b {
		c.JSON(200, util.Response{
			Code:    666,
			Message: util.NoticeSuccess,
			Data:    notice,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    555,
		Message: util.NoticeFail,
		Data:    nil,
	})

}
