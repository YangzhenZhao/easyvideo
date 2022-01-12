package service

import "context"

type TagService interface {
	GetVideoTags(ctx context.Context, videoId int32) []string
}

type tagService struct{}

func (s tagService) GetVideoTags(ctx context.Context, videoId int32) []string {
	return []string{}
}

func New() TagService {
	return tagService{}
}
