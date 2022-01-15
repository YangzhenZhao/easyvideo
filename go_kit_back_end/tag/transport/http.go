package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/endpoint"
	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/golang/protobuf/ptypes/empty"
)

func NewHTTPHandler(endpoints endpoint.Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/get_video_tags", kitHttp.NewServer(
		endpoints.GetVideoTagsEndpoint,
		decodeGetVideoTagsRequest,
		encodeGetVideoTagsResponse,
	))
	m.Handle("/tags", kitHttp.NewServer(
		endpoints.GetAllTagsEndpoint,
		decodeGetAllTagsRequest,
		encodeGetAllTagsResponse,
	))
	return m
}

func decodeGetVideoTagsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetVideoTagsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func encodeGetVideoTagsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

func decodeGetAllTagsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return &empty.Empty{}, nil
}

func encodeGetAllTagsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
	resp := response.(endpoint.GetAllTagsResponse)
	err = json.NewEncoder(w).Encode(resp.Tags)
	return
}
