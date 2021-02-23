package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"graduationproject/model"
	"graduationproject/util"
)

func TousubaoxiuRouters(e *gin.Engine)  {
	e.POST("/tousubaoxiu",TousuOrBaoxiu)
	e.GET("/findtbbyusername",FindTsBxByUsername)
}

func TousuOrBaoxiu(c *gin.Context)  {
	var tb model.TSBX
	if err:=c.ShouldBindJSON(&tb);err!=nil {
		util.CheckError(err)
	}
	add := model.AddTB(tb)
	if add {
		c.String(200,"true")
	}
	c.String(200,"false")
}

func FindTsBxByUsername(c *gin.Context)  {
	query := c.Query("username")
	fmt.Println(query)
	username := model.GetTBByUsername(query)
	if username!=nil {
		c.JSON(200,model.CommonResult{
			Code:    1,
			Message: util.GetDataSuccess,
			Data:    username,
		})
	}else {
		c.JSON(200,model.CommonResult{
			Code:    0,
			Message: util.GetDataFail,
			Data:    nil,
		})
	}


}
