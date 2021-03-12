package router

import (
	"github.com/gin-gonic/gin"
	"graduationproject/model"
	"graduationproject/util"
	"strconv"
)

func PropertyRouters(e *gin.Engine) {
	group := e.Group("/property")
	{
		group.GET("/get", getProperty)
		group.GET("/pay", payProperty)
	}
}

//@Summary getProperty
//@Tags 物业费模块
//@Description 获取物业费
//@Param houseId query string true "房屋ID"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /property/get [get]
func getProperty(c *gin.Context) {
	query := c.Query("houseId")
	data, err := model.GetPropertyByHouseId(query)
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

//@Summary payProperty
//@Tags 物业费模块
//@Description 缴纳物业费
//@Param houseId query string true "房屋ID"
//@Param property query string true "缴纳金额"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /property/pay [get]
func payProperty(c *gin.Context) {
	query := c.Query("houseId")
	s := c.Query("property")
	data, err := strconv.ParseFloat(s, 10)
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.PropertyFail,
			Data:    nil,
		})
	}
	err = model.PayProperty(float32(data), query)
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.PropertyFail,
			Data:    nil,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: util.PropertySuccess,
		Data:    data,
	})
}

func CronProperty() {

}
