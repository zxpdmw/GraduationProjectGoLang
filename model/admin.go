package model

import (
	"graduationproject/util"
)

//管理员结构体
type Admin struct {
	Id       int    `json:"id" , gorm:"primarykey"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (receiver Admin) TableName() string {
	return "t_admin"
}

func AdminLogin(username, password string) bool {
	var c int64
	count := util.Db.Table("t_admin").Model(&Admin{}).Where("username=? And password=?", username, password).Count(&c)
	if count.RowsAffected == 1 && c == 1 {
		return true
	}
	return false
}

func AdminRegister(username, password string) bool {
	var a = Admin{
		Username: username,
		Password: password,
	}
	create := util.Db.Create(&a)
	if create.RowsAffected == 1 {
		return true
	}
	return false
}
