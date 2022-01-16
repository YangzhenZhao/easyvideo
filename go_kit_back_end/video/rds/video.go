package rds

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type IVideoDao interface {
	InitConf(db *gorm.DB, storageDir string)
	GetAllVideosData(ctx context.Context) ([]Video, error)
}

type videoDaoImpl struct {
	DB         *gorm.DB
	StorageDir string
}

var VideoDao IVideoDao = &videoDaoImpl{}

func (dao *videoDaoImpl) InitConf(db *gorm.DB, storageDir string) {
	dao.DB = db
	dao.StorageDir = storageDir
}

func (dao *videoDaoImpl) GetAllVideosData(ctx context.Context) ([]Video, error) {
	var videos []Video
	if result := dao.DB.Find(&videos); result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	return videos, nil
}
