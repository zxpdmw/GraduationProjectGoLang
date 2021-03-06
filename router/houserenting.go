package router

import (
	"github.com/gin-gonic/gin"
	"graduationproject/model"
	"graduationproject/util"
	"strconv"
)

func HouseRentingRouters(e *gin.Engine) {
	group := e.Group("/houserent")
	{
		group.GET("/rent", Rent)
		group.GET("/sale", Sale)
		group.GET("/rentsale", RentSale)
	}
}

func Rent(c *gin.Context) {
	all, b := model.RentAll()
	if b {
		c.JSON(200, util.CommonResult{
			Code:    666,
			Message: util.RentSuccess,
			Data:    all,
		})
		return
	}
	c.JSON(200, util.CommonResult{
		Code:    555,
		Message: util.RentFail,
		Data:    all,
	})
}

func Sale(c *gin.Context) {
	all, b := model.SaleAll()
	if b {
		c.JSON(200, util.CommonResult{
			Code:    666,
			Message: util.SaleSuccess,
			Data:    all,
		})
		return
	}
	c.JSON(200, util.CommonResult{
		Code:    555,
		Message: util.SaleFail,
		Data:    all,
	})
}

func RentSale(c *gin.Context) {
	all, b := model.RentSaleAll()
	if b {
		c.JSON(200, util.CommonResult{
			Code:    666,
			Message: util.RentSaleSuccess,
			Data:    all,
		})
		return
	}
	c.JSON(200, util.CommonResult{
		Code:    555,
		Message: util.RentSaleFail,
		Data:    all,
	})
}

func Delete(c *gin.Context) {
	query := c.Query("id")
	atoi, _ := strconv.Atoi(query)
	b := model.Delete(atoi)
	if b {
		c.JSON(200, gin.H{
			"code":    666,
			"message": util.DeleteHouseSuccess,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    555,
		"message": util.DeleteHouseFail,
	})
}
