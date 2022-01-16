package service

import (
	"context"

	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/video/external/tag"
	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/video/rds"
)

type VideoMetadata struct {
	Name      string   `json:"name"`
	BytesSize int64    `json:"bytesSize"`
	Tags      []string `json:"tags"`
}

type VideoService interface {
	GetAllVideosMetadata(ctx context.Context) ([]VideoMetadata, error)
	GetCoverPicturePath(ctx context.Context, videoName string) (string, error)
}

type videoService struct{}

func (s videoService) GetAllVideosMetadata(ctx context.Context) ([]VideoMetadata, error) {
	videosData, err := rds.VideoDao.GetAllVideosData(ctx)
	if err != nil {
		return nil, err
	}
	resList := []VideoMetadata{}
	for _, videoData := range videosData {
		tagNames := tag.Client.GetVideoTags(ctx, videoData.ID)
		resList = append(resList, VideoMetadata{
			Name:      videoData.Name,
			BytesSize: int64(videoData.BytesSize),
			Tags:      tagNames,
		})
	}

	return resList, nil
}

func (s videoService) GetCoverPicturePath(ctx context.Context, videoName string) (string, error) {
	return rds.VideoDao.GetCoverPicturePathByVideoName(ctx, videoName)
}

func New() VideoService {
	return videoService{}
}
