package routers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/YangzhenZhao/easyvideo/gin_back_end/models"
	"github.com/YangzhenZhao/easyvideo/gin_back_end/settings"
	"github.com/gin-gonic/gin"
)

type VideoViewItem struct {
	Name      string   `json:"name"`
	BytesSize uint64   `json:"bytesSize"`
	Tags      []string `json:"tags"`
}

func allVideos(c *gin.Context) {
	var videos []models.Video
	if result := settings.DB.Find(&videos); result.Error != nil {
		fmt.Println(result.Error)
	}

	resList := []VideoViewItem{}
	for _, video := range videos {
		var videoTags []models.VideoTag
		settings.DB.Where("video_id = ?", video.ID).Find(&videoTags)
		tagList := []string{}
		for _, videoTag := range videoTags {
			var tag models.Tag
			settings.DB.Where("id = ?", videoTag.TagId).Find(&tag)
			tagList = append(tagList, tag.TagName)
		}
		resList = append(resList, VideoViewItem{
			Name:      video.Name,
			BytesSize: video.BytesSize,
			Tags:      tagList,
		})
	}

	c.JSON(200, resList)
}

func allTags(c *gin.Context) {
	resList := []string{}
	var tags []models.Tag
	settings.DB.Find(&tags)
	for _, tag := range tags {
		resList = append(resList, tag.TagName)
	}
	c.JSON(200, resList)
}

type VideoName struct {
	VideoName string `uri:"video_name"`
}

func coverPicture(c *gin.Context) {
	var videoName VideoName
	if err := c.ShouldBindUri(&videoName); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	var video models.Video
	settings.DB.Where("name = ?", videoName.VideoName).First(&video)
	c.File(video.CoverPictruePath)
}

func download(c *gin.Context) {
	var videoName VideoName
	if err := c.ShouldBindUri(&videoName); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	var video models.Video
	settings.DB.Where("name = ?", videoName.VideoName).First(&video)
	c.File(video.VideoPath)
}

type TagNames struct {
	TagNames string `uri:"tag_names"`
}

type void struct{}

func getIntersection(set map[uint]void, videoTags []models.VideoTag) map[uint]void {
	resSet := make(map[uint]void)
	for _, videoTag := range videoTags {
		if _, exists := set[videoTag.VideoId]; exists {
			resSet[videoTag.VideoId] = void{}
		}
	}
	return resSet
}

func tagsVideos(c *gin.Context) {
	var tagNames TagNames
	if err := c.ShouldBindUri(&tagNames); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	tagList := strings.Split(tagNames.TagNames, ",")

	set := make(map[uint]void)
	for i, tagName := range tagList {
		var tag models.Tag
		settings.DB.Where("tag_name = ?", tagName).First(&tag)
		var videoTags []models.VideoTag
		settings.DB.Where("tag_id = ?", tag.ID).Find(&videoTags)
		if i == 0 {
			for _, videoTag := range videoTags {
				set[videoTag.VideoId] = void{}
			}
		} else {
			set = getIntersection(set, videoTags)
		}
	}

	resList := []VideoViewItem{}
	for videoId, _ := range set {
		var video models.Video
		settings.DB.Where("id = ?", videoId).First(&video)
		var videoTags []models.VideoTag
		settings.DB.Where("video_id = ?", videoId).Find(&videoTags)
		tagList := []string{}
		for _, videoTag := range videoTags {
			var tag models.Tag
			settings.DB.Where("id = ?", videoTag.TagId).First(&tag)
			tagList = append(tagList, tag.TagName)
		}
		resList = append(resList, VideoViewItem{
			Name:      video.Name,
			BytesSize: video.BytesSize,
			Tags:      tagList,
		})
	}
	c.JSON(200, resList)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			// 可将将* 替换为指定的域名
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

func InitRouters() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())

	r.GET("/videos", allVideos)
	r.GET("/tags", allTags)
	r.GET("/cover_picture/:video_name", coverPicture)
	r.GET("/download/:video_name", download)
	r.GET("/tags_videos/:tag_names", tagsVideos)
	return r
}
