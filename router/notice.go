package router

import (
	"github.com/gin-gonic/gin"
	"graduationproject/model"
	"graduationproject/util"
	"net/http"
)

func NoticeRouters(e *gin.Engine) {
	group := e.Group("/notice")
	{
		group.GET("/recommend", recommendNotice)
		group.GET("/detail", detailNotice)
		group.GET("/delete", deleteNotice)
		group.POST("/publish", publishNotice)
		group.POST("/info", editNotice)
	}

}

//@Summary recommendNotice
//@Tags 社区公告模块
//@Description 获取推荐公告，按发布时间最新排序
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /notice/recommend [get]
func recommendNotice(c *gin.Context) {
	data, err := model.RecommendNotice()
	if err != nil {
		c.JSON(http.StatusOK, util.Response{
			Code:    555,
			Message: util.RecommendNoticeFail,
			Data:    data,
		})
		return
	}

	c.JSON(http.StatusOK, util.Response{
		Code:    666,
		Message: util.RecommendNoticeSuccess,
		Data:    data,
	})
}

//@Summary detailNotice
//@Tags 社区公告模块
//@Description 获取公告的详细信息
//@Param title query string true "公告标题"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /notice/detail [get]
func detailNotice(c *gin.Context) {
	query := c.Query("title")
	data, err := model.DetailNotice(query)
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.DetailNoticeFail,
			Data:    data,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: util.DetailNoticeSuccess,
		Data:    data,
	})
}

//@Summary deleteNotice
//@Tags 社区公告模块
//@Description 删除公告根据标题
//@Param title query string true "公告标题"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /notice/delete [get]
func deleteNotice(c *gin.Context) {
	query := c.Query("title")
	err := model.DeleteNotice(query)
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})

		return
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: util.RequestSuccess,
		Data:    nil,
	})
}

//@Summary publishNotice
//@Tags 社区公告模块
//@Description 发布社区公告
//@Param notice body model.Notice true "Notice结构体"
//@Success 200 {object} util.Response
//@Failure 500 {object} util.Response
//@Router /notice/publish [post]
func publishNotice(c *gin.Context) {
	var n model.Notice
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
	}
	err := model.PublishNotice(n.Title, n.Content, n.Publisher)
	if err != nil {
		c.JSON(200, util.Response{
			Code:    555,
			Message: util.RequestFail,
			Data:    nil,
		})
		return
	}
	c.JSON(200, util.Response{
		Code:    666,
		Message: util.RequestSuccess,
		Data:    nil,
	})
}

func editNotice(c *gin.Context) {

}
