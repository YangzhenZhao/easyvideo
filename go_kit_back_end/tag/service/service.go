package service

import (
	"context"
	"fmt"
)

type TagService interface {
	GetVideoTags(ctx context.Context, videoId int32) []string
}

type tagService struct{}

func (s tagService) GetVideoTags(ctx context.Context, videoId int32) []string {
	return []string{fmt.Sprintf("%d", videoId)}
}

func New() TagService {
	return tagService{}
}
