package model

import (
	"graduationproject/util"
	"time"
)

//公告结构体
type Notice struct {
	ID          int       `json:"id" gorm:"primarykey"`
	Publisher   string    `json:"publisher"`    //公告发布者
	Title       string    `json:"title"`        //公告发布标题
	Content     string    `json:"content"`      //公告内容
	PublishTime time.Time `json:"publish_time"` //公告发布时间
}

func (Notice) TableName() string {
	return "t_notice"
}

//获取推荐公告 按时间最新排序获取
func RecommendNotice() (bool, []Notice) {
	var n []Notice
	find := util.Db.Order("publish_time desc").Find(&n)
	if find.RowsAffected != 0 {
		return true, n
	}
	return false, n
}

//获取公告详情
func DetailNotice(title string) (bool, Notice) {
	var n Notice
	find := util.Db.Where("title=?", title).Find(&n)
	if find.RowsAffected == 1 {
		return true, n
	}
	return false, n
}

//删除公告
func DeleteNotice(title string) bool {
	tx := util.Db.Where("title=?", title).Delete(&Notice{})
	if tx.RowsAffected == 1 {
		return true
	}
	return false
}

//发布公告
func PublishNotice(title string, content string, publisher string) bool {

	var n = Notice{
		Publisher:   publisher,
		Title:       title,
		Content:     content,
		PublishTime: time.Now(),
	}
	create := util.Db.Create(&n)
	if create.RowsAffected == 1 {
		return true
	}
	return false
}

func EditNotice() {

}

func GetAllNotice() ([]Notice, bool) {
	var ns []Notice
	find := util.Db.Find(&ns)
	if find.RowsAffected == 0 {
		return nil, false
	}
	return ns, true
}
