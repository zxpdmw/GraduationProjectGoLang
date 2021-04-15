package router

import (
	"github.com/gin-gonic/gin"
	"graduationproject/model"
	"graduationproject/util"
)

func HouseSaleRouters(e *gin.Engine) {
	group := e.Group("/sale")
	{
		group.GET("/all", sale)
		group.POST("/publish", publishSaleHouse)
		group.GET("/delete", deleteSaleHouseById)
		group.GET("/editprice", editHouseSalePrice)
	}
}

//@Summary Sale
//@Tags 房屋租售模块
//@Description 获取全部的售房信息
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /houserentsale/sale [get]
func sale(c *gin.Context) {
	data, err := model.SaleAll()
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.SaleFail,
			Data:    data,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: util.SaleSuccess,
		Data:    data,
	})
}

//@Summary DeleteHouseRentSale
//@Tags 房屋租售模块
//@Description 根据ID删除房屋
//@Param id query string true "房屋主键ID"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /houserentsale/delete [get]
func deleteSaleHouseById(c *gin.Context) {
	id := c.Query("id")
	err := model.DeleteHouseSale(id)
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.DeleteHouseFail,
			Data:    nil,
		})
		return
	}

	c.JSON(200, util.Response{
		Code:    666,
		Message: util.DeleteHouseSuccess,
		Data:    nil,
	})
}

//@Summary publishHouse
//@Tags 房屋租售模块
//@Description 发布房屋租售信息
//@Param houseRent body model.HouseRentSale true "房屋租售结构体"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /houserentsale/publish [post]
func publishSaleHouse(c *gin.Context) {
	s := model.HouseSaleRent{}
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
	}
	err := model.PublishHouseSale(s)
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

//@Summary editHousePrice
//@Tags 房屋租赁模块
//@Description 修改房屋价格
//@Param id query string true "房屋唯一ID"
//@Param price query string true "新的房屋价格"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /houserentsale/editprice [get]
func editHouseSalePrice(c *gin.Context) {
	id := c.Query("id")
	price := c.Query("price")
	err := model.EditHouseSalePrice(id, price)
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
