package model

import (
	"gorm.io/gorm"
	"graduationproject/util"
)

type Property struct {
	ID      int     `json:"id" gorm:"primarykey"`
	HouseID string  `json:"house_id"` //房屋ID
	Balance float32 `json:"balance"`  //物业费账户
}

func (Property) TableName() string {
	return "t_property"
}

//获取每间房屋剩余的物业费
func GetPropertyByHouseId(houseId string) (data float32, err error) {
	find := util.Db.Table("t_property").Model(&User{}).Select("balance").Where("house_id=?", houseId).Find(&data).Error
	if find != nil {
		return
	}
	return
}

//缴纳物业费
func PayProperty(property float32, houseId string) (err error) {
	var r float32
	err = util.Db.Table("t_property").Select("balance").Where("house_id=?", houseId).Find(&r).Error
	if err != nil {
		return
	}
	err = util.Db.Table("t_property").Where("house_id=?", houseId).Update("balance", r+property).Error
	if err != nil {
		return
	}
	return
}

func UserBindHouseID(property Property) (err error) {
	err = util.Db.Table("t_property").Create(&property).Error
	if err != nil {
		return err
	}
	return err
}

func CronProperty() {
	util.Db.Table("t_property").Where("1=?", 1).
		Update("balance", gorm.Expr("balance-?", 5))
}
