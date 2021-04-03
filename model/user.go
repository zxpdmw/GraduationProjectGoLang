package model

import "graduationproject/util"

type User struct {
	ID       int    `json:"id" gorm:"primarykey"`
	Username string `json:"username"` //用户账号
	Password string `json:"password"` //用户密码
	Nickname string `json:"nickname"` //用户昵称
	HouseId  string `json:"house_id"` //房屋ID
	Address  string `json:"address"`  //用户地址
	Phone    string `json:"phone"`    //用户电话
}

type UserRegister struct {
	Username string `json:"username"` //用户账号
	Password string `json:"password"` //用户密码
	Nickname string `json:"nickname"` //用户昵称
	HouseId  string `json:"house_id"` //房屋ID
}

type UserInfo struct {
	Username string `json:"username"` //用户账号
	Nickname string `json:"nickname"` //用户昵称
	HouseId  string `json:"house_id"` //房屋ID
	Address  string `json:"address"`  //用户地址
	Phone    string `json:"phone"`    //用户电话
}

func (User) TableName() string {
	return "t_user"
}

func Login(name, password string) (user User, err error) {
	err = util.Db.Table("t_user").Where("username=?", name).Where("password=?", password).Find(&user).Error
	if err != nil {
		return
	}
	return
}

func Register(register UserRegister) (err error) {
	var u = User{
		Username: register.Username,
		Password: register.Password,
		Nickname: register.Nickname,
		HouseId:  register.HouseId,
	}
	err = util.Db.Create(&u).Error
	if err != nil {
		return
	}
	_, f := GetPropertyByHouseId(register.HouseId)
	if f != nil {
		return
	}
	//util.Rdb.Set(register.HouseId, id, 0)
	return
}

func EditInfo(u User) (err error) {
	err = util.Db.Model(&User{}).Where("username=?", u.Username).Updates(&u).Error
	if err != nil {
		return
	}
	return
}

func EditPassword(u User) (err error) {
	err = util.Db.Model(&User{}).Where("username=?", u.Username).Update("password", u.Password).Error
	if err != nil {
		return
	}
	return
}

func GetInfo(username string) (data UserInfo, err error) {
	err = util.Db.Table("t_user").
		Select("username,nickname,house_id,address,phone").
		Where("username=?", username).
		Find(&data).Error
	if err != nil {
		return
	}
	return
}

func CheckUserExist(username string) (sum int64, err error) {
	err = util.Db.Table("t_user").Where("username=?", username).Count(&sum).Error
	if err != nil {
		return
	}
	return
}

func GetHouseId(usernaem string) (houseId string, err error) {
	err = util.Db.Table("t_user").Select("house_id").Where("username=?", usernaem).Find(&houseId).Error
	if err != nil {
		return
	}
	return
}
