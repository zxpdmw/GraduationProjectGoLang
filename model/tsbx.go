package model

import "graduationproject/util"

//投诉报修结构体
type TSBX struct {
	ID       int    `json:"id" gorm:"primarykey"`
	CRType   string `json:"cr_type"` //投诉 报修
	Status   string `json:"status"`  //未处理 处理中 已处理
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Message  string `json:"message"`
	Username string `json:"username"`
}

func (TSBX) TableName() string {
	return "t_tsbx"
}

//添加投诉报修
func AddTB(tb TSBX) bool {
	create := util.Db.Create(&tb)
	if create.RowsAffected == 1 {
		return true
	}
	return false
}

//获取user 添加的投诉保修
func GetTBByUsername(username string) []TSBX {
	var tbs []TSBX
	find := util.Db.Where("username=?", username).Find(&tbs)
	if find.RowsAffected != 0 {
		return tbs
	} else {
		return nil
	}
}

//管理员获取全部的投诉报修
func GetAllCR() ([]TSBX, bool) {
	var tbs []TSBX
	find := util.Db.Table("t_tsbx").Find(&tbs)
	if find.RowsAffected != 0 {
		return tbs, true
	}
	return nil, false
}
