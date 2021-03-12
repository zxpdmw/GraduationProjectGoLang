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
func GetPropertyByHouseId(houseId string) (data float32, err error) {
	find := util.Db.Table("t_property").Model(&User{}).Select("amount").Where("house_id=?", houseId).Find(&data).Error
	if find != nil {
		return
	}
	return
}

//缴纳物业费
func PayProperty(property float32, houseId string) (err error) {
	var r float32
	err = util.Db.Table("t_property").Select("amount").Where("house_id=?", houseId).Find(&r).Error
	if err != nil {
		return
	}
	err = util.Db.Table("t_property").Where("house_id=?", houseId).Update("amount", r+property).Error
	if err != nil {
		return
	}
	return
}
