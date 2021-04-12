package model

import (
	"graduationproject/util"
)

//家政服务结构体
type HouseKeeping struct {
	ID       int    `json:"id" gorm:"primarykey"`
	HKType   string `json:"hk_type"`                //家政服务类型
	Address  string `json:"address"`                //地址
	Phone    string `json:"phone"`                  //联系方式
	Status   string `json:"status" gorm:"default "` //服务状态
	Username string `json:"username"`               //用户账户

}

func (HouseKeeping) TableName() string {
	return "t_house_keeping"
}

func AddHouseKeeping(keeping HouseKeeping) (ks []HouseKeeping, err error) {
	keeping.Status = "未处理"
	err = util.Db.Table("t_house_keeping").Create(&keeping).Error
	err = util.Db.Table("t_house_keeping").Where("username=?", keeping.Username).Find(&ks).Error
	return
}

func GetAllHouseKeeping() (data []HouseKeeping, err error) {
	err = util.Db.Find(&data).Error
	return
}

func EditHouseKeepingStatus(id int) (err error) {
	err = util.Db.Table("t_house_keeping").Where("id=?", id).Update("status", "已处理").Error
	return
}

func GetHouseKeepingByUsername(username string) (data []HouseKeeping, err error) {
	err = util.Db.Table("t_house_keeping").Where("username=?", username).Find(&data).Error
	return
}

func DeleteHouseKeeping(id string) (err error) {
	err = util.Db.Table("t_house_keeping").Where("id=?", id).Delete(&HouseKeeping{}).Error
	return
}
