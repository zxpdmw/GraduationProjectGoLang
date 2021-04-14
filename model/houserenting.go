package model

import (
	"graduationproject/util"
	"strconv"
)

//房屋租售结构体
type HouseRentSale struct {
	ID       int    `json:"id" gorm:"primarykey"`
	Username string `json:"username"` //发布者
	HRType   string `json:"hr_type"`  //房屋类型 租还是售
	Message  string `json:"message"`  //房屋信息
	Address  string `json:"address"`  //房屋地址
	Phone    string `json:"phone"`    //联系方式
	Price    string `json:"price"`    //房屋价格
}

type EditHouse struct {
	ID       string `json:"id"`
	Username string `json:"username"` //发布者
	HRType   string `json:"hr_type"`  //房屋类型 租还是售
	Message  string `json:"message"`  //房屋信息
	Address  string `json:"address"`  //房屋地址
	Phone    string `json:"phone"`    //联系方式
	Price    string `json:"price"`    //房屋价格
}

func (HouseRentSale) TableName() string {
	return "t_house_rent"
}

//获取全部租房信息
func RentAll() (data []HouseRentSale, err error) {
	err = util.Db.Table("t_house_renting").Where("hr_type=?", "rent").Find(&data).Error
	return
}

//获取全部售房信息
func SaleAll() (data []HouseRentSale, err error) {
	err = util.Db.Table("t_house_renting").Where("hr_type=?", "sale").Find(&data).Error
	return
}

//获取全部租售房信息
func RentSaleAll() (data []HouseRentSale, err error) {
	err = util.Db.Table("t_house_renting").Find(&data).Error
	return
}

//删除租售房信息
func DeleteHouseRentSale(id string) (err error) {
	err = util.Db.Table("t_house_renting").Where("id=?", id).Delete(&HouseRentSale{}).Error
	return
}

//发布房屋信息
func PublishHouse(rent HouseRentSale) (err error) {
	err = util.Db.Table("t_house_renting").Create(&rent).Error
	return
}

//修改房屋信息
func EditHouseInfo(rent EditHouse) (err error) {
	atoi, _ := strconv.Atoi(rent.ID)
	var hr = HouseRentSale{
		ID:       atoi,
		Username: rent.Username,
		HRType:   rent.HRType,
		Message:  rent.Message,
		Address:  rent.Address,
		Phone:    rent.Phone,
		Price:    rent.Price,
	}
	err = util.Db.Table("t_house_renting").Updates(&hr).Error
	return
}

func GetByUsername(username string) (house []HouseRentSale, err error) {
	err = util.Db.Table("t_house_renting").Where("username=?", username).Find(&house).Error
	return
}

func EditHousePrice(id string, price string) (err error) {
	err = util.Db.Table("t_house_renting").Where("id=?", id).Update("price", price).Error
	return
}
