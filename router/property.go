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
		group.GET("/getProperty", getProperty)
		group.GET("/payProperty", payProperty)
	}
}

func getProperty(c *gin.Context) {
	query := c.Query("houseId")
	id, f := model.GetPropertyByHouseId(query)
	if id {
		c.JSON(200, util.CommonResult{
			Code:    666,
			Message: util.GetDataSuccess,
			Data:    f,
		})
		return
	}
	c.JSON(200, util.CommonResult{
		Code:    555,
		Message: util.GetDataFail,
		Data:    nil,
	})
}

func payProperty(c *gin.Context) {
	query := c.Query("houseId")
	s := c.Query("amount")
	float, err := strconv.ParseFloat(s, 10)
	util.CheckError(err)
	property := model.PayProperty(float32(float), query)
	if property {
		c.JSON(200, gin.H{
			"code":    666,
			"message": util.PropertySuccess,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    555,
		"message": util.PropertyFail,
	})
}
