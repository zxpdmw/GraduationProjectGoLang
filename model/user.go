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
	_, err = GetPropertyByHouseId(register.HouseId)
	return
}

func EditInfo(u User) (err error) {
	err = util.Db.Model(&User{}).Where("username=?", u.Username).Updates(&u).Error
	return
}

func EditPassword(username, password string) (err error) {
	err = util.Db.Table("t_user").Where("username=?", username).Update("password", password).Error
	return
}

func GetInfo(username string) (data UserInfo, err error) {
	err = util.Db.Table("t_user").
		Select("username,nickname,house_id,address,phone").
		Where("username=?", username).
		Find(&data).Error
	return
}

func CheckUserExist(username string) (sum int64, err error) {
	err = util.Db.Table("t_user").Where("username=?", username).Count(&sum).Error
	return
}

func GetHouseId(usernaem string) (houseId string, err error) {
	err = util.Db.Table("t_user").Select("house_id").Where("username=?", usernaem).Find(&houseId).Error
	return
}

func EditNickname(username, nickname string) (err error) {
	err = util.Db.Table("t_user").Where("username=?", username).Update("nickname", nickname).Error
	return
}

func EditAddress(username, address string) (err error) {
	err = util.Db.Table("t_user").Where("username=?", username).Update("address", address).Error
	return
}

func EditPhone(username, phone string) (err error) {
	err = util.Db.Table("t_user").Where("username=?", username).Update("phone", phone).Error
	return
}

func EditHouseID(username, house string) (err error) {
	err = util.Db.Table("t_user").Where("username=?", username).Update("house_id", house).Error
	return
}
