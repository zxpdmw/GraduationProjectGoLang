package model

import "graduationproject/util"

//房屋租售结构体
type HouseRent struct {
	ID        int    `json:"id" gorm:"primarykey"`
	Publisher string `json:"publisher"` //发布者
	HRType    string `json:"hr_type"`   //房屋类型 租还是售
	Message   string `json:"content"`   //房屋信息
	Address   string `json:"address"`   //房屋地址
	Phone     string `json:"phone"`     //联系方式
}

func (HouseRent) TableName() string {
	return "t_house_rent"
}

func RentAll() ([]HouseRent, bool) {
	var hrs []HouseRent
	find := util.Db.Table("t_house_rent").Where("hr_type=?", "租").Find(&hrs)
	if find.RowsAffected != 0 {
		return hrs, true
	}
	return hrs, false
}

func SaleAll() ([]HouseRent, bool) {
	var hrs []HouseRent
	find := util.Db.Table("t_house_rent").Where("hr_type=?", "售").Find(&hrs)
	if find.RowsAffected != 0 {
		return hrs, true
	}
	return hrs, false
}

func RentSaleAll() ([]HouseRent, bool) {
	var hrs []HouseRent
	find := util.Db.Table("t_house_rent").Find(&hrs)
	if find.RowsAffected != 0 {
		return hrs, true
	}
	return hrs, false
}

func Delete(id int) bool {
	tx := util.Db.Table("t_house_rent").Model(&HouseRent{}).Where("id=?", id).Delete(&HouseRent{})
	if tx.RowsAffected == 1 {
		return true
	}
	return false
}
