package service

import (
	"net"
	"net/http"

	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/endpoint"
	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/service"
	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/transport"
)

func Run() {
	var (
		service     = service.New()
		endpoints   = endpoint.New(service)
		httpHandler = transport.NewHTTPHandler(endpoints)
	)
	httpListener, _ := net.Listen("tcp", ":8081")
	http.Serve(httpListener, httpHandler)
}
