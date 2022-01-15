package rds

import (
	"context"

	"gorm.io/gorm"
)

type IVideoTagDao interface {
	InitConf(db *gorm.DB)
	GetTagIDsByVideoID(ctx context.Context, videoID int32) []int32
}

type videoTagImpl struct {
	DB *gorm.DB
}

var VideoTagDao IVideoTagDao = &videoTagImpl{}

func (dao *videoTagImpl) InitConf(db *gorm.DB) {
	dao.DB = db
}

func (dao *videoTagImpl) GetTagIDsByVideoID(ctx context.Context, videoID int32) []int32 {
	var videoTags []VideoTag
	dao.DB.Where("video_id = ?", videoID).Find(&videoTags)
	tagIDs := []int32{}
	for _, videoTag := range videoTags {
		tagIDs = append(tagIDs, videoTag.TagId)
	}
	return tagIDs
}
