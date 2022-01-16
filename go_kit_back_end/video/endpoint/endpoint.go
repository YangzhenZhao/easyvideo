package endpoint

import (
	"context"
	"fmt"

	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/video/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/golang/protobuf/ptypes/empty"
)

type Endpoints struct {
	GetAllVideosMetadataEndpoint endpoint.Endpoint
}

func New(svc service.VideoService) Endpoints {
	return Endpoints{
		GetAllVideosMetadataEndpoint: MakeGetAllVideosMetadataEmdpoint(svc),
	}
}

func MakeGetAllVideosMetadataEmdpoint(s service.VideoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		videosMetadata, err := s.GetAllVideosMetadata(ctx)
		return GetAllVideosMetadataResponse{VideosMetadata: videosMetadata}, err
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

type GetAllVideosMetadataResponse struct {
	VideosMetadata []service.VideoMetadata `json:"videosMetadata"`
}
