package routers

import (
	"github.com/YangzhenZhao/easyvideo/gin_back_end/models"
	"github.com/YangzhenZhao/easyvideo/gin_back_end/settings"
	"github.com/gin-gonic/gin"
)

func allVideos(c *gin.Context) {
	var video models.Video
	settings.DB.First(&video)
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func InitRouters() *gin.Engine {
	r := gin.Default()
	r.GET("/videos", allVideos)
	return r
}
