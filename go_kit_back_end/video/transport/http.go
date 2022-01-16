package transport

import (
	"context"
	"encoding/json"
	"net/http"

	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/golang/protobuf/ptypes/empty"

	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/video/endpoint"
)

func NewHTTPHandler(endpoints endpoint.Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/videos", kitHttp.NewServer(
		endpoints.GetAllVideosMetadataEndpoint,
		decodecGetAllVideosMetadataRequest,
		encodeGetAllVideosMetadataResponse,
	))
	return m
}

func decodecGetAllVideosMetadataRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return &empty.Empty{}, nil
}

func encodeGetAllVideosMetadataResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
	resp := response.(endpoint.GetAllVideosMetadataResponse)
	err = json.NewEncoder(w).Encode(resp.VideosMetadata)
	return
}
