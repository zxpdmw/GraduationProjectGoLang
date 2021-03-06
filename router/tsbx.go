package router

import (
	"github.com/gin-gonic/gin"
	"graduationproject/model"
	"graduationproject/util"
)

func ComplainRepairRouters(e *gin.Engine) {
	group := e.Group("/tsbx")
	{
		group.POST("/addcr", addComplainRepair)
		group.GET("/getcr", getComplainRepairByUsername)
		group.GET("/getall", GetAllCR)
	}

}

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

func getComplainRepairByUsername(c *gin.Context) {
	query := c.Query("username")
	username := model.GetTBByUsername(query)
	if username != nil {
		c.JSON(200, util.CommonResult{
			Code:    1,
			Message: util.GetDataSuccess,
			Data:    username,
		})
		return
	}
	c.JSON(200, util.CommonResult{
		Code:    0,
		Message: util.GetDataFail,
		Data:    nil,
	})

}

func GetAllCR(c *gin.Context) {
	cr, b := model.GetAllCR()
	if b {
		c.JSON(200, util.CommonResult{
			Code:    666,
			Message: util.RequestSuccess,
			Data:    cr,
		})
		return
	}
	c.JSON(200, util.CommonResult{
		Code:    555,
		Message: util.RequestFail,
		Data:    nil,
	})
}
