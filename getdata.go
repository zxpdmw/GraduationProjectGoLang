package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"graduationproject/model"
	"io"
	"net/http"
)

type Notice struct {
	Id          string `json:"id"`
	Ctime       string `json:"ctime"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Source      string `json:"source"`
	PicUrl      string `json:"picUrl"`
	Url         string `json:"url"`
}

type Reponse struct {
	Code     int      `json:"code"`
	Msg      string   `json:"msg"`
	Newslist []Notice `json:"newslist"`
}

func main() {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", "zxpdmw", "Zxpdmw520", "rm-2zeqer8186x8o6hi9vo.mysql.rds.aliyuncs.com", 3306, "graduationproject")
	Db, _ := gorm.Open(mysql.Open(conn), &gorm.Config{})
	get, _ := http.Get("http://api.tianapi.com/esports/index?key=e7a558928691361882a5962cd136c3eb&num=10&rand=1&page=3")
	all, _ := io.ReadAll(get.Body)
	get.Body.Close()
	fmt.Println(string(all))
	s := Reponse{}
	json.Unmarshal([]byte(all), &s)
	fmt.Println(s.Code)
	fmt.Println(s.Msg)
	for _, notice := range s.Newslist {
		no := model.Notice{
			Publisher:   "张惟宇",
			Title:       notice.Title,
			Content:     notice.Description,
			PublishTime: notice.Ctime,
		}
		err := Db.Table("t_notice").Create(&no).Error
		fmt.Println(err)
	}
}
