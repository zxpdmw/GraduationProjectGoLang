package model

import "graduationproject/util"

type Property struct {
	ID int `json:"id" gorm:"primarykey"`
	HouseID string `json:"house_id"`
	Amount float32 `json:"amount"`
	Address string `json:"address"`
}

type result struct {
	PropertyCosts float32
}

func ( Property) TableName()string  {
	return "t_property"
}

func GetPropertyByHouseId(houseid string) (bool,float32) {
	var r result
	find := util.Db.Model(&User{}).Select("property_costs").Where("house_id=?", houseid).Find(&r)
	if find.RowsAffected==1 {
		return true,r.PropertyCosts
	}else {
		return false,0.0
	}
}