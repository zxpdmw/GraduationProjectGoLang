package model

import "graduationproject/util"

type Property struct {
	ID      int     `json:"id" gorm:"primarykey"`
	HouseID string  `json:"house_id"` //房屋ID
	Amount  float32 `json:"amount"`   //物业费账户
	Address string  `json:"address"`  //房屋地址
}

func (Property) TableName() string {
	return "t_property"
}

//获取每间房屋剩余的物业费
func GetPropertyByHouseId(houseId string) (bool, float32) {
	var r float32
	find := util.Db.Table("t_property").Model(&User{}).Select("amount").Where("house_id=?", houseId).Find(&r)
	if find.RowsAffected == 1 {
		return true, r
	} else {
		return false, 0.0
	}
}

//缴纳物业费
func PayProperty(fee float32, houseId string) bool {
	var r float32
	find := util.Db.Table("t_property").Select("amount").Where("house_id=?", houseId).Find(&r)
	if find.RowsAffected == 1 {
		update := util.Db.Table("t_property").Where("house_id=?", houseId).Update("amount", r+fee)
		if update.RowsAffected == 1 {
			return true
		}
		return false
	}
	return false
}
