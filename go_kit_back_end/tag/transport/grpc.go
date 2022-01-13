package transport

import (
	"context"

	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/endpoint"
	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/pb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	getVideoTags grpctransport.Handler
	pb.UnsafeTagServiceServer
}

func (s *grpcServer) GetVideoTags(ctx context.Context, req *pb.GetVideoTagsRequest) (*pb.GetVideoTagsResponse, error) {
	_, rep, err := s.getVideoTags.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetVideoTagsResponse), nil
}

func (grpcServer) mustEmbedUnimplementedTagServiceServer() {}

func NewGRPCServer(endpoints endpoint.Endpoints) pb.TagServiceServer {
	return &grpcServer{
		getVideoTags: grpctransport.NewServer(
			endpoints.GetVideoTagsEndpoint,
			decodeGRPCGetVideoTagsRequest,
			encodeGRPCGetVideoTagsResponse,
		),
	}
}

func decodeGRPCGetVideoTagsRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetVideoTagsRequest)
	return endpoint.GetVideoTagsRequest{VideoID: req.VideoID}, nil
}

func encodeGRPCGetVideoTagsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetVideoTagsResponse)
	return &pb.GetVideoTagsResponse{Tags: resp.Tags}, nil
}
