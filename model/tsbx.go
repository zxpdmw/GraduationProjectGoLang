package model

import (
	"graduationproject/util"
)

type TSBX struct {
	ID       int    `json:"id" gorm:"primarykey"`
	CRType   string `json:"cr_type"`
	Status   string `json:"status"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Message  string `json:"message"`
	Username string `json:"username"`
}

func (TSBX) TableName() string {
	return "t_tsbx"
}

func AddTB(tb TSBX) bool {
	create := util.Db.Create(&tb)
	if create.RowsAffected == 1 {
		return true
	}
	return false
}

func GetTBByUsername(username string) []TSBX {
	var tbs []TSBX
	find := util.Db.Where("username=?", username).Find(&tbs)
	if find.RowsAffected != 0 {
		return tbs
	} else {
		return nil
	}
}
