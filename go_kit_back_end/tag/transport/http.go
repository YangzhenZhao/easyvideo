package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/endpoint"
	kitHttp "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(endpoints endpoint.Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/get_video_tags", kitHttp.NewServer(
		endpoints.GetVideoTagsEndpoint,
		decodeGetVideoTagsRequest,
		encodeGetVideoTagsResponse,
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
