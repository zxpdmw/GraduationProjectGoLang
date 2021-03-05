package model

import "graduationproject/util"

type Property struct {
	ID      int     `json:"id" gorm:"primarykey"`
	HouseID string  `json:"house_id"`
	Amount  float32 `json:"amount"`
	Address string  `json:"address"`
}

type result struct {
	Amount float32
}

func (Property) TableName() string {
	return "t_property"
}

func GetPropertyByHouseId(houseId string) (bool, float32) {
	var r result
	find := util.Db.Table("t_property").Model(&User{}).Select("amount").Where("house_id=?", houseId).Find(&r)
	if find.RowsAffected == 1 {
		return true, r.Amount
	} else {
		return false, 0.0
	}
}

func PayProperty(fee float32, houseId string) bool {
	var r result
	find := util.Db.Table("t_property").Select("amount").Where("house_id=?", houseId).Find(&r)
	if find.RowsAffected == 1 {
		update := util.Db.Table("t_property").Where("house_id=?", houseId).Update("amount", r.Amount+fee)
		if update.RowsAffected == 1 {
			return true
		}
		return false
	}
	return false
}
