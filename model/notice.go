package model

import (
	"graduationproject/util"
	"time"
)

//公告结构体
type Notice struct {
	ID          int    `json:"id" gorm:"primarykey"`
	Publisher   string `json:"publisher"`    //公告发布者
	Title       string `json:"title"`        //公告发布标题
	Content     string `json:"content"`      //公告内容
	PublishTime string `json:"publish_time"` //公告发布时间
}

func (Notice) TableName() string {
	return "t_notice"
}

//获取推荐公告 按时间最新排序获取
func RecommendNotice() (data []Notice, err error) {
	find := util.Db.Table("t_notice").Order("publish_time desc").Find(&data).Error
	if find != nil {
		return
	}
	return
}

//获取公告详情
func DetailNotice(title string) (data Notice, err error) {
	find := util.Db.Table("t_notice").Where("title=?", title).Find(&data).Error
	if find != nil {
		return
	}
	return
}

//删除公告
func DeleteNotice(title string) (err error) {
	tx := util.Db.Table("t_notice").Where("title=?", title).Delete(&Notice{}).Error
	if tx != nil {
		return
	}
	return
}

//发布公告
func PublishNotice(title, content, publisher string) (err error) {

	var n = Notice{
		Publisher:   publisher,
		Title:       title,
		Content:     content,
		PublishTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	err = util.Db.Table("t_notice").Create(&n).Error
	if err != nil {
		return
	}
	return
}

func EditNotice() {

}

func GetAllNotice() (data []Notice, err error) {
	find := util.Db.Table("t_notice").Find(&data).Error
	if find != nil {
		return
	}
	return
}
