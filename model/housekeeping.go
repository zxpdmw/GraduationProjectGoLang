package model

//家政服务结构体
type HouseKeeping struct {
	ID      int    `json:"id" gorm:"primarykey"`
	HKType  string `json:"hk_type"` //家政服务类型
	Address string `json:"address"` //地址
	Phone   string `json:"phone"`   //联系方式
	Status  string `json:"status"`  //服务状态

}

func (HouseKeeping) TableName() string {
	return "t_house_keeping"
}
