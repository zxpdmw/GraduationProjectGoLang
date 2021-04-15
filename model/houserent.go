package model

import (
	"graduationproject/util"
)

//房屋租售结构体
type HouseRent struct {
	ID          int    `json:"id" gorm:"primarykey"`
	Username    string `json:"username"` //发布者
	Message     string `json:"message"`  //房屋信息
	Address     string `json:"address"`  //房屋地址
	Phone       string `json:"phone"`    //联系方式
	Price       string `json:"price"`    //房屋价格
	Area        string `json:"area"`
	Orientation string `json:"orientation"`
	Floor       string `json:"floor"`
	Ruzhu       string `json:"ruzhu"`
	T           string `json:"t"`
}

//获取全部租房信息
func RentAll() (data []HouseSaleRent, err error) {
	err = util.Db.Table("t_house_rent").Find(&data).Error
	return
}

//删除租房信息
func DeleteHouseRent(id string) (err error) {
	err = util.Db.Table("t_house_rent").Where("id=?", id).Delete(&HouseSaleRent{}).Error
	return
}

//发布房屋信息
func PublishRentHouse(rent HouseSaleRent) (err error) {
	rent.T = "rent"
	err = util.Db.Table("t_house_rent").Create(&rent).Error
	return
}

func GetRentByUsername(username string) (house []HouseSaleRent, err error) {
	err = util.Db.Table("t_house_rent").Where("username=?", username).Find(&house).Error
	return
}

func EditHouseRentPrice(id string, price string) (err error) {
	err = util.Db.Table("t_house_rent").Where("id=?", id).Update("price", price).Error
	return
}
