package router

import (
	"github.com/gin-gonic/gin"
	"graduationproject/model"
	"graduationproject/util"
	"strconv"
)

func HouseRentingRouters(e *gin.Engine) {
	group := e.Group("/houserentsale")
	{
		group.GET("/rent", Rent)
		group.GET("/sale", Sale)
		group.GET("/all", RentSale)
	}
}

//@Summary Rent
//@Tags 房屋租售模块
//@Description 获取全部租房信息
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /houserentsale/rent [get]
func Rent(c *gin.Context) {
	all, b := model.RentAll()
	if b {
		c.JSON(200, util.Response{
			Code:    666,
			Message: util.RentSuccess,
			Data:    all,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    555,
		Message: util.RentFail,
		Data:    all,
	})
}

//@Summary Sale
//@Tags 房屋租售模块
//@Description 获取全部的售房信息
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /houserentsale/sale [get]
func Sale(c *gin.Context) {
	all, b := model.SaleAll()
	if b {
		c.JSON(200, util.Response{
			Code:    666,
			Message: util.SaleSuccess,
			Data:    all,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    555,
		Message: util.SaleFail,
		Data:    all,
	})
}

//@Summary RentSale
//@Tags 房屋租售模块
//@Description 获取全部的房屋信息
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /houserentsale/all [get]
func RentSale(c *gin.Context) {
	all, b := model.RentSaleAll()
	if b {
		c.JSON(200, util.Response{
			Code:    666,
			Message: util.RentSaleSuccess,
			Data:    all,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    555,
		Message: util.RentSaleFail,
		Data:    all,
	})
}

//@Summary DeleteHouseRentSale
//@Tags 房屋租售模块
//@Description 根据ID删除房屋
//@Param id query string true "房屋主键ID"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /houserentsale [get]
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
