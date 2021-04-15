package router

import (
	"github.com/gin-gonic/gin"
	"graduationproject/model"
	"graduationproject/util"
)

func HouseRentRouters(e *gin.Engine) {
	group := e.Group("/rent")
	{
		group.GET("/all", Rent)
		group.POST("/publish", publisRenthHouse)
		group.GET("/delete", DeleteRentHouseById)
		group.GET("/editprice", editHouseRentPrice)
	}
}

//@Summary Rent
//@Tags 房屋租售模块
//@Description 获取全部租房信息
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /houserentsale/rent [get]
func Rent(c *gin.Context) {
	data, err := model.RentAll()
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.RentFail,
			Data:    data,
		})

		return
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: util.RentSuccess,
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
func DeleteRentHouseById(c *gin.Context) {
	id := c.Query("id")
	err := model.DeleteHouseRent(id)
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
func publisRenthHouse(c *gin.Context) {
	s := model.HouseSaleRent{}
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
	}
	err := model.PublishRentHouse(s)
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
func editHouseRentPrice(c *gin.Context) {
	id := c.Query("id")
	price := c.Query("price")
	err := model.EditHouseRentPrice(id, price)
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
