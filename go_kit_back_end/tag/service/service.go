package service

import (
	"context"

	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/rds"
)

type TagService interface {
	GetVideoTags(ctx context.Context, videoID int32) []string
	GetAllTags(ctx context.Context) []string
}

type tagService struct{}

func (s tagService) GetVideoTags(ctx context.Context, videoID int32) []string {
	return rds.GetTagNamesByVideoID(ctx, videoID)
}

func (s tagService) GetAllTags(ctx context.Context) []string {
	return rds.TagDao.GetAllTagNames(ctx)
}

func New() TagService {
	return tagService{}
}
