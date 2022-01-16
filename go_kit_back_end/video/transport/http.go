package transport

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"

	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/golang/protobuf/ptypes/empty"

	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/video/endpoint"
)

const (
	coverPictureRoutePath = "/cover_picture/"
)

func NewHTTPHandler(endpoints endpoint.Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/videos", kitHttp.NewServer(
		endpoints.GetAllVideosMetadataEndpoint,
		decodecGetAllVideosMetadataRequest,
		encodeGetAllVideosMetadataResponse,
	))
	m.Handle(coverPictureRoutePath, kitHttp.NewServer(
		endpoints.GetCoverPicturePathByVideoNameEndpoint,
		decodeGetCoverPicturePathByVideoNameRequest,
		encodeGetCoverPicturePathByVideoNameResponse,
	))
	return m
}

func decodeGetCoverPicturePathByVideoNameRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	videoName, err := url.QueryUnescape(r.RequestURI[len(coverPictureRoutePath):])
	if err != nil {
		return nil, err
	}
	req := endpoint.GetCoverPicturePathByVideoNameRequest{
		VideoName: videoName,
	}
	return req, nil
}

func encodeGetCoverPicturePathByVideoNameResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	resp := response.(endpoint.GetCoverPicturePathByVideoNameResponse)
	coverPicturePath := resp.CoverPicturePath
	file, err := os.Open(coverPicturePath)
	if err != nil {
		return
	}
	w.Header().Add("Content-Type", "application/octet-stream")
	fileName := path.Base(coverPicturePath)
	w.Header().Add("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	io.Copy(w, file)
	return
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
