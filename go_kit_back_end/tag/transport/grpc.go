package transport

import (
	"context"

	"google.golang.org/grpc"

	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/endpoint"
	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/pb"
	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/service"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/golang/protobuf/ptypes/empty"
)

type grpcServer struct {
	getVideoTags grpctransport.Handler
	getAllTags   grpctransport.Handler
	pb.UnsafeTagServiceServer
}

func (s *grpcServer) GetVideoTags(ctx context.Context, req *pb.GetVideoTagsRequest) (*pb.GetVideoTagsResponse, error) {
	_, rep, err := s.getVideoTags.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetVideoTagsResponse), nil
}

func (s *grpcServer) GetAllTags(ctx context.Context, req *empty.Empty) (*pb.GetAllTagsResponse, error) {
	_, rep, err := s.getAllTags.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllTagsResponse), nil
}

func NewGRPCServer(endpoints endpoint.Endpoints) pb.TagServiceServer {
	return &grpcServer{
		getVideoTags: grpctransport.NewServer(
			endpoints.GetVideoTagsEndpoint,
			decodeGRPCGetVideoTagsRequest,
			encodeGRPCGetVideoTagsResponse,
		),
		getAllTags: grpctransport.NewServer(
			endpoints.GetAllTagsEndpoint,
			decodeGRPCGetAllTagsRequest,
			encodeGRPCGetAllTagsResponse,
		),
	}
}

func NewGRPCClient(conn *grpc.ClientConn) service.TagService {
	return endpoint.Endpoints{
		GetVideoTagsEndpoint: grpctransport.NewClient(
			conn,
			"tag.TagService",
			"GetVideoTags",
			encodeGRPCGetVideoTagsRequest,
			decodeGRPCGetVideoTagsResponse,
			pb.GetVideoTagsResponse{},
		).Endpoint(),
		GetAllTagsEndpoint: grpctransport.NewClient(
			conn,
			"tag.TagService",
			"GetAllTags",
			encodeGRPCGetAllTagsRequest,
			decodeGRPCGetAllTagsResponse,
			pb.GetAllTagsResponse{},
		).Endpoint(),
	}
}

func decodeGRPCGetVideoTagsRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetVideoTagsRequest)
	return endpoint.GetVideoTagsRequest{VideoID: req.VideoID}, nil
}

func encodeGRPCGetVideoTagsResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetVideoTagsResponse)
	return &pb.GetVideoTagsResponse{Tags: resp.Tags}, nil
}

func decodeGRPCGetAllTagsRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return &empty.Empty{}, nil
}

func encodeGRPCGetAllTagsResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetVideoTagsResponse)
	return &pb.GetAllTagsResponse{Tags: resp.Tags}, nil
}

func encodeGRPCGetVideoTagsRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetVideoTagsRequest)
	return &pb.GetVideoTagsRequest{VideoID: req.VideoID}, nil
}

func decodeGRPCGetVideoTagsResponse(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	resp := grpcReq.(*pb.GetVideoTagsResponse)
	return endpoint.GetVideoTagsResponse{Tags: resp.Tags}, nil
}

func encodeGRPCGetAllTagsRequest(ctx context.Context, request interface{}) (interface{}, error) {
	return &empty.Empty{}, nil
}

func decodeGRPCGetAllTagsResponse(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	resp := grpcReq.(*pb.GetAllTagsResponse)
	return endpoint.GetAllTagsResponse{Tags: resp.Tags}, nil
}
