package model

import "time"

type Notice struct {
	ID int `json:"id" gorm:"primarykey"`
	Publisher string `json:"publisher"`
	Title string `json:"title"`
	Content string `json:"content"`
	PublishTime time.Time `json:"publish_time"`
}

func (n Notice) TableName()string  {
	return "t_notice"
}
