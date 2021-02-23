package model
//房屋租售结构体
type HouseRent struct {
	ID int `json:"id" gorm:"primarykey"`
	Publisher string `json:"publisher"`  //发布者
	Status string `json:"status"`        //房屋类型 租还是售
	Content string `json:"content"`		//房屋信息
	Address string `json:"address"`		//房屋地址
	Phone string `json:"phone"`			//联系方式
}

func (receiver HouseRent) TableName()string  {
	return "t_house_rent"
}
