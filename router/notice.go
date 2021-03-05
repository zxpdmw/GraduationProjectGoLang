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

func recommendNotice(c *gin.Context) {
	notice, notices := model.RecommendNotice()
	if notice {
		c.JSON(http.StatusOK, util.CommonResult{
			Code:    666,
			Message: util.RecommendNoticeSuccess,
			Data:    notices,
		})
		return
	}
	c.JSON(http.StatusOK, util.CommonResult{
		Code:    555,
		Message: util.RecommendNoticeFail,
		Data:    notices,
	})
}

func detailNotice(c *gin.Context) {
	query := c.Query("title")
	notice, n := model.DetailNotice(query)
	if notice {
		c.JSON(200, util.CommonResult{
			Code:    666,
			Message: util.DetailNoticeSuccess,
			Data:    n,
		})
		return
	}
	c.JSON(200, util.CommonResult{
		Code:    555,
		Message: util.DetailNoticeFail,
		Data:    n,
	})
}

func deleteNotice(c *gin.Context) {
	query := c.Query("title")
	util.Db.Where("title=?", query).Delete(&model.Notice{})
}

func publishNotice(c *gin.Context) {
	var n model.Notice
	if err := c.ShouldBindJSON(&n); err != nil {
		util.CheckError(err)
	} else {
		notice := model.PublishNotice(n.Title, n.Content, n.Publisher)
		if notice {
			c.JSON(200, gin.H{
				"code":    666,
				"message": util.PublishNoticeSuccess,
			})
			return
		}
		c.JSON(200, gin.H{
			"code":    555,
			"message": util.PublishNoticeFail,
		})
	}
	c.JSON(200, gin.H{
		"code":    555,
		"message": util.RequestFail,
	})
}

func editNotice(c *gin.Context) {

}
