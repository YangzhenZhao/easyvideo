package rds

import (
	"context"

	"gorm.io/gorm"
)

type ITagDao interface {
	InitConf(db *gorm.DB)
	GetAllTagNames(ctx context.Context) []string
	GetTagNameByID(ctx context.Context, tagID int32) string
}

type tagDaoImpl struct {
	DB *gorm.DB
}

var TagDao ITagDao = &tagDaoImpl{}

func (dao *tagDaoImpl) InitConf(db *gorm.DB) {
	dao.DB = db
}

func (dao *tagDaoImpl) GetAllTagNames(ctx context.Context) []string {
	resList := []string{}
	var tags []Tag
	dao.DB.Find(&tags)
	for _, tag := range tags {
		resList = append(resList, tag.TagName)
	}
	return resList
}

func (dao *tagDaoImpl) GetTagNameByID(ctx context.Context, tagID int32) string {
	var tag Tag
	dao.DB.Where("id = ?", tagID).Find(&tag)
	return tag.TagName
}
