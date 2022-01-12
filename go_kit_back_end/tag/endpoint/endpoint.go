package endpoint

import (
	"context"

	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetVideoTags endpoint.Endpoint
}

func New(svc service.TagService) Endpoints {
	return Endpoints{
		GetVideoTags: MakeGetVideoTagsEndpoint(svc),
	}
}

func MakeGetVideoTagsEndpoint(s service.TagService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetVideoTagsRequest)
		tags := s.GetVideoTags(ctx, req.VideoID)
		return GetVideoTagsResponse{Tags: tags}, nil
	}
}

type GetVideoTagsRequest struct {
	VideoID int32 `json:"videoID"`
}

type GetVideoTagsResponse struct {
	Tags []string `json:"tags"`
}
