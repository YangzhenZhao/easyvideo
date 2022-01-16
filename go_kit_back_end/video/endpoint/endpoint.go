package endpoint

import (
	"context"
	"fmt"

	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/video/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/golang/protobuf/ptypes/empty"
)

type Endpoints struct {
	GetAllVideosMetadataEndpoint           endpoint.Endpoint
	GetCoverPicturePathByVideoNameEndpoint endpoint.Endpoint
}

func New(svc service.VideoService) Endpoints {
	return Endpoints{
		GetAllVideosMetadataEndpoint:           MakeGetAllVideosMetadataEmdpoint(svc),
		GetCoverPicturePathByVideoNameEndpoint: MakeGetCoverPicturePathByVideoNameEndpoint(svc),
	}
}

func MakeGetAllVideosMetadataEmdpoint(s service.VideoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		videosMetadata, err := s.GetAllVideosMetadata(ctx)
		return GetAllVideosMetadataResponse{VideosMetadata: videosMetadata}, err
	}
}

func MakeGetCoverPicturePathByVideoNameEndpoint(s service.VideoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetCoverPicturePathByVideoNameRequest)
		coverPicturePath, err := s.GetCoverPicturePath(ctx, req.VideoName)
		return GetCoverPicturePathByVideoNameResponse{CoverPicturePath: coverPicturePath}, err
	}
}

func (s Endpoints) GetAllVideosMetadata(ctx context.Context) []service.VideoMetadata {
	resp, err := s.GetAllVideosMetadataEndpoint(ctx, empty.Empty{})
	if err != nil {
		fmt.Printf("err:%+v\n", err)
		return []service.VideoMetadata{}
	}
	response := resp.(GetAllVideosMetadataResponse)
	return response.VideosMetadata
}

func (s Endpoints) GetCoverPicturePathByVideoName(ctx context.Context, videoName string) (string, error) {
	resp, err := s.GetCoverPicturePathByVideoNameEndpoint(ctx, GetCoverPicturePathByVideoNameRequest{VideoName: videoName})
	if err != nil {
		fmt.Printf("err:%+v\n", err)
		return "", nil
	}
	response := resp.(GetCoverPicturePathByVideoNameResponse)
	return response.CoverPicturePath, nil
}

type GetAllVideosMetadataResponse struct {
	VideosMetadata []service.VideoMetadata `json:"videosMetadata"`
}

type GetCoverPicturePathByVideoNameRequest struct {
	VideoName string `json:"videoName"`
}

type GetCoverPicturePathByVideoNameResponse struct {
	CoverPicturePath string `json:"coverPicturePath"`
}
