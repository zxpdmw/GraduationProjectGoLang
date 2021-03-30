package model

import (
	"graduationproject/util"
)

//管理员结构体
type Admin struct {
	Id       int    `json:"id" gorm:"primarykey"`
	Username string `json:"username"` //管理员登录账号
	Password string `json:"password"` //管理员登录密码
}

func (receiver Admin) TableName() string {
	return "t_admin"
}

func AdminLogin(username, password string) (err error) {
	var c int64
	err = util.Db.Table("t_admin").Where("username=? And password=?", username, password).Count(&c).Error
	return
}

func AdminRegister(username, password string) (err error) {
	var a = Admin{
		Username: username,
		Password: password,
	}
	err = util.Db.Create(&a).Error
	if err != nil {
		return
	}
	return
}
