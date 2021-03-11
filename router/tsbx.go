package router

import (
	"github.com/gin-gonic/gin"
	"graduationproject/model"
	"graduationproject/util"
)

func ComplainRepairRouters(e *gin.Engine) {
	group := e.Group("/tsbx")
	{
		group.POST("/add", addComplainRepair)
		group.GET("/get", getComplainRepairByUsername)
		group.GET("/all", GetAllComplainRepair)
	}

}

//@Summary addComplainRepair
//@Tags 投诉报修模块
//@Description 添加投诉报修
//@Param tb body model.TSBX true "TSBX结构体"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /tsbx/add [post]
func addComplainRepair(c *gin.Context) {
	var tb model.TSBX
	if err := c.ShouldBindJSON(&tb); err != nil {
		util.CheckError(err)
	}
	add := model.AddTB(tb)
	if add {
		c.JSON(200, gin.H{
			"code":    666,
			"message": util.ComplainRepairSuccess,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    555,
		"message": util.ComplainRepairFail,
	})
}

//@Summary getComplainRepairByUsername
//@Description 获取本人的投诉报修
//@Tags 投诉报修模块
//@Param username query string true "用户账号"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /tsbx/get [get]
func getComplainRepairByUsername(c *gin.Context) {
	query := c.Query("username")
	username := model.GetTBByUsername(query)
	if username != nil {
		c.JSON(200, util.Response{
			Code:    1,
			Message: util.GetDataSuccess,
			Data:    username,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    0,
		Message: util.GetDataFail,
		Data:    nil,
	})

}

//@Summary GetAllComplainRepair
//@Tags 投诉报修模块
//@Description 后天管理员获取全部的投诉报修
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /tsbx/all [get]
func GetAllComplainRepair(c *gin.Context) {
	cr, b := model.GetAllCR()
	if b {
		c.JSON(200, util.Response{
			Code:    666,
			Message: util.RequestSuccess,
			Data:    cr,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    555,
		Message: util.RequestFail,
		Data:    nil,
	})
}
