package endpoint

import (
	"context"
	"fmt"

	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/golang/protobuf/ptypes/empty"
)

type Endpoints struct {
	GetVideoTagsEndpoint endpoint.Endpoint
	GetAllTagsEndpoint   endpoint.Endpoint
}

func New(svc service.TagService) Endpoints {
	return Endpoints{
		GetVideoTagsEndpoint: MakeGetVideoTagsEndpoint(svc),
		GetAllTagsEndpoint:   MakeGetAllTagsEndpoint(svc),
	}
}

func MakeGetVideoTagsEndpoint(s service.TagService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetVideoTagsRequest)
		tags := s.GetVideoTags(ctx, req.VideoID)
		return GetVideoTagsResponse{Tags: tags}, nil
	}
}

func MakeGetAllTagsEndpoint(s service.TagService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		tags := s.GetAllTags(ctx)
		return GetAllTagsResponse{Tags: tags}, nil
	}
}

func (s Endpoints) GetVideoTags(ctx context.Context, videoID int32) []string {
	resp, err := s.GetVideoTagsEndpoint(ctx, GetVideoTagsRequest{VideoID: videoID})
	if err != nil {
		fmt.Printf("err:%+v\n", err)
		return []string{}
	}
	response := resp.(GetVideoTagsResponse)
	return response.Tags
}

func (s Endpoints) GetAllTags(ctx context.Context) []string {
	resp, err := s.GetAllTagsEndpoint(ctx, empty.Empty{})
	if err != nil {
		fmt.Printf("err:%+v\n", err)
		return []string{}
	}
	response := resp.(GetAllTagsResponse)
	return response.Tags
}

type GetVideoTagsRequest struct {
	VideoID int32 `json:"videoID"`
}

type GetVideoTagsResponse struct {
	Tags []string `json:"tags"`
}

type GetAllTagsResponse struct {
	Tags []string `json:"tags"`
}
