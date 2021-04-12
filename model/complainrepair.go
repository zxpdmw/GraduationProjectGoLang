package model

import "graduationproject/util"

//投诉报修结构体
type ComplainRepair struct {
	ID       int    `json:"id" gorm:"primarykey"`
	CRType   string `json:"cr_type"`  //投诉 报修
	Status   string `json:"status"`   //未处理 处理中 已处理
	Address  string `json:"address"`  //投诉报修者地址
	Phone    string `json:"phone"`    //投诉报修者电话
	Message  string `json:"message"`  //投诉报修详细信息
	Username string `json:"username"` //投诉报修者账户
}

func (ComplainRepair) TableName() string {
	return "t_complain_repair"
}

//修改后台投诉报修处理的状态
func EditComplainRepairStatus(id int) (err error) {
	err = util.Db.Table("t_complain_repair").Where("id=?", id).Update("status", "已处理").Error
	return
}

//添加投诉报修
func AddCR(cr ComplainRepair) (data []ComplainRepair, err error) {
	cr.Status = "未处理"
	err = util.Db.Table("t_complain_repair").Create(&cr).Error
	err = util.Db.Table("t_complain_repair").Where("username=?", cr.Username).Find(&data).Error
	return
}

//获取user添加的投诉保修
func GetCRByUsername(username string) (data []ComplainRepair, err error) {
	err = util.Db.Table("t_complain_repair").Where("username=?", username).Find(&data).Error
	return
}

//管理员获取全部的投诉报修
func GetAllCR() (data []ComplainRepair, err error) {
	err = util.Db.Table("t_complain_repair").Find(&data).Error
	return
}

func DeleteCRById(id string) (err error) {
	err = util.Db.Table("t_complain_repair").Where("id=?", id).Delete(&ComplainRepair{}).Error
	return
}
