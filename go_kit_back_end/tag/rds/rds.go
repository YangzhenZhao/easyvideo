package rds

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/config"
)

func InitMysql(config config.RDSConfig) error {
	DB, err := gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	TagDao.InitConf(DB)
	VideoTagDao.InitConf(DB)
	return nil
}

func GetTagNamesByVideoID(ctx context.Context, videoID int32) []string {
	tagIDs := VideoTagDao.GetTagIDsByVideoID(ctx, videoID)
	tagNames := make([]string, len(tagIDs))
	for i, tagID := range tagIDs {
		tagNames[i] = TagDao.GetTagNameByID(ctx, tagID)
	}
	return tagNames
}
