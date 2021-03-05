package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"graduationproject/model"
	"graduationproject/util"
)

func ComplainRepairRouters(e *gin.Engine) {
	group := e.Group("/tsbx")
	{
		group.POST("/addcr", addComplainRepair)
		group.GET("/getcr", getComplainRepair)
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

func getComplainRepair(c *gin.Context) {
	query := c.Query("username")
	fmt.Println(query)
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
