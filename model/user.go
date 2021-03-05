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

type UserInfo struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	HouseId  string `json:"house_id"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
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

func Register(name string, password string, nickname string, houseId string) bool {
	var u = User{
		Username: name,
		Password: password,
		Nickname: nickname,
		HouseId:  houseId,
	}
	create := util.Db.Create(&u)
	if create.RowsAffected == 1 {
		id, f := GetPropertyByHouseId(houseId)
		if id {
			util.Rdb.Set(houseId, f, 0)
			return true
		} else {
			return false
		}
		return true
	} else {
		return false
	}
}

func EditInfo(u User) bool {
	update := util.Db.Model(&User{}).Where("username=?", u.Username).Update("nickname", u.Nickname).Update("address", u.Address).Update("phone", u.Phone).Update("house_id", u.HouseId)
	if update.RowsAffected == 1 {
		return true
	}
	return false
}

func EditPassword(u User) bool {
	update := util.Db.Model(&User{}).Where("username=?", u.Username).Update("password", u.Password)
	if update.RowsAffected == 1 {
		return true
	}
	return false
}

func GetInfo(username string) (bool, UserInfo) {
	var ui UserInfo
	find := util.Db.Table("t_user").Select("username,nickname,house_id,address,phone").Where("username=?", username).Find(&ui)
	if find.RowsAffected == 1 {
		return true, ui
	}
	return false, ui
}
