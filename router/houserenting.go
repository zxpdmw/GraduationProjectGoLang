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
		group.POST("/publish", publishHouse)
		group.POST("/edit", editHouse)
		group.GET("/delete", DeleteHouseById)
		group.GET("/get", getByUsername)
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

//@Summary Sale
//@Tags 房屋租售模块
//@Description 获取全部的售房信息
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /houserentsale/sale [get]
func Sale(c *gin.Context) {
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

//@Summary RentSale
//@Tags 房屋租售模块
//@Description 获取全部的房屋信息
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /houserentsale/all [get]
func RentSale(c *gin.Context) {
	data, err := model.RentSaleAll()
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.RentSaleFail,
			Data:    data,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: util.RentSaleSuccess,
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
func DeleteHouseById(c *gin.Context) {
	query := c.Query("id")
	atoi, _ := strconv.Atoi(query)
	err := model.DeleteHouseRentSale(atoi)
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
func publishHouse(c *gin.Context) {
	s := model.HouseRentSale{}
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
	}
	err := model.PublishHouse(s)
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

//@Summary editHouse
//@Tags 房屋租售模块
//@Description 修改房屋信息
//@Param editHouse body model.EditHouse true "修改房屋信息结构体"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /houserentsale/edit [post]
func editHouse(c *gin.Context) {
	s := model.EditHouse{}
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
	}
	err := model.EditHouseInfo(s)
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

func getByUsername(c *gin.Context) {
	query := c.Query("username")
	username, err := model.GetByUsername(query)
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
		Data:    username,
	})
}
