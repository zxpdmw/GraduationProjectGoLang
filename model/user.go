package model

import (
	"graduationproject/util"
)

type User struct {
	ID       int    `json:"id" gorm:"primarykey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	HouseId  string `json:"house_id"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	//PropertyCosts float32 `json:"property_costs"`
}

func (User) TableName() string {
	return "t_user"
}

func Login(name string, password string) bool {
	var u User
	util.Db.Where("username=? and password=? ", name, password).Find(&u)
	if u.Username != name || u.Password != password {
		return false
	}
	return true
}

func Register(name string, password string, nickname string, houseid string) bool {
	var u = User{
		Username: name,
		Password: password,
		Nickname: nickname,
		HouseId:  houseid,
	}
	create := util.Db.Create(&u)
	if create.RowsAffected == 1 {
		id, f := GetPropertyByHouseId(houseid)
		if id {
			util.Rdb.Set(houseid, f, 0)
			return true
		} else {
			return false
		}
		return true
	} else {
		return false
	}
}

func EditInfo(address string, phone string) bool {
	return false
}
