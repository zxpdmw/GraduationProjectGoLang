package router

import (
	"github.com/gin-gonic/gin"
	"graduationproject/model"
	"graduationproject/util"
)

func HousekeepingRouters(e *gin.Engine) {
	group := e.Group("/housekeeping")
	{
		group.GET("/get", getHouseKeepingByUsername)
		group.POST("/add", addHouseKeeping)
		group.GET("/delete", deleteHouseKeeping)
	}
}

//@Summary addHouseKeeping
//@Tags 家政服务模块
//@Description 添加一个家政服务请求
//@Param houseKeeping body model.HouseKeeping true "家政服务结构体"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /housekeeping/add [post]
func addHouseKeeping(c *gin.Context) {
	var hk model.HouseKeeping
	if err := c.ShouldBindJSON(&hk); err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
	}
	ks, err := model.AddHouseKeeping(hk)
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: util.RequestSuccess,
		Data:    ks,
	})
}

//@Summary getHouseKeepingByUsername
//@Tags 家政服务模块
//@Description 获取每个人申请的家政服务
//@Param username query string true "用户名"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /housekeeping/get [get]
func getHouseKeepingByUsername(c *gin.Context) {
	query := c.Query("username")
	data, err := model.GetHouseKeepingByUsername(query)
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

//@Summary deleteHouseKeeping
//@Tags 家政服务模块
//@Description 如果用户不需要已申请的家政服务，可以选择取消，
//@Param username query string true "用户账户"
//@Param hk_type query string true "家政服务类型"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /housekeeping/delete [get]
func deleteHouseKeeping(c *gin.Context) {
	query := c.Query("username")
	s := c.Query("hk_type")
	err := model.DeleteHouseKeeping(s, query)
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
		Data:    nil,
	})
}
