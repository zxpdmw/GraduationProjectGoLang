package router

import (
	"github.com/gin-gonic/gin"
	"graduationproject/model"
	"graduationproject/util"
)

func ComplainRepairRouters(e *gin.Engine) {
	group := e.Group("/complainrepair")
	{
		group.POST("/add", addComplainRepair)
		group.GET("/get", getComplainRepairByUsername)
		group.GET("/complain")
		group.GET("repair")
	}

}

//@Summary addComplainRepair
//@Tags 投诉报修模块
//@Description 添加投诉报修
//@Param tb body model.ComplainRepair true "投诉报修结构体"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /complainrepair/add [post]
func addComplainRepair(c *gin.Context) {
	var tb model.ComplainRepair
	if err := c.ShouldBindJSON(&tb); err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.ComplainRepairFail,
			Data:    nil,
		})
		return
	}
	addTB, err := model.AddTB(tb)
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.ComplainRepairFail,
			Data:    nil,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: util.ComplainRepairSuccess,
		Data:    addTB,
	})
}

//@Summary getComplainRepairByUsername
//@Description 获取本人的投诉报修
//@Tags 投诉报修模块
//@Param username query string true "用户账号"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /complainrepair/get [get]
func getComplainRepairByUsername(c *gin.Context) {
	query := c.Query("username")
	data, err := model.GetTBByUsername(query)
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.GetDataFail,
			Data:    nil,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: util.GetDataSuccess,
		Data:    data,
	})
}
