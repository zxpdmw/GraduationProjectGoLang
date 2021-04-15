package model

import "graduationproject/util"

type HouseSaleRent struct {
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
func SaleAll() (data []HouseSaleRent, err error) {
	err = util.Db.Table("t_house_sale").Find(&data).Error
	return
}

//删除租售房信息
func DeleteHouseSale(id string) (err error) {
	err = util.Db.Table("t_house_sale").Where("id=?", id).Delete(&HouseSaleRent{}).Error
	return
}

//发布房屋信息
func PublishHouseSale(rent HouseSaleRent) (err error) {
	rent.T = "sale"
	err = util.Db.Table("t_house_sale").Create(&rent).Error
	return
}

func EditHouseSalePrice(id string, price string) (err error) {
	err = util.Db.Table("t_house_sale").Where("id=?", id).Update("price", price).Error
	return
}

func GetSaleByUsername(username string) (house []HouseSaleRent, err error) {
	err = util.Db.Table("t_house_sale").Where("username=?", username).Find(&house).Error
	return
}
