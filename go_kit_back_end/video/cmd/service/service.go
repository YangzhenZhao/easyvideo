package service

import (
	"net"
	"net/http"

	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/video/config"
	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/video/endpoint"
	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/video/external/tag"
	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/video/rds"
	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/video/service"
	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/video/transport"
)

func Run() {
	rds.InitMysql(config.RDSConfig{
		Dsn:        "root:woaini123@tcp(127.0.0.1:3306)/easyvideo?charset=utf8mb4&parseTime=True&loc=Local",
		StorageDir: "/Users/nocilantro/Downloads/storage",
	})
	tag.Init()
	var (
		service     = service.New()
		endpoints   = endpoint.New(service)
		httpHandler = transport.NewHTTPHandler(endpoints)
	)

	httpListener, _ := net.Listen("tcp", ":8000")
	http.Serve(httpListener, httpHandler)
}
