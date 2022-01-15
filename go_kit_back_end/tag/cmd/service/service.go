package service

import (
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/config"
	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/endpoint"
	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/pb"
	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/rds"
	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/service"
	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/transport"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

func Run() {
	rds.InitMysql(config.RDSConfig{
		Dsn: "root:woaini123@tcp(127.0.0.1:3306)/easyvideo?charset=utf8mb4&parseTime=True&loc=Local",
	})
	var (
		service     = service.New()
		endpoints   = endpoint.New(service)
		httpHandler = transport.NewHTTPHandler(endpoints)
		grpcServer  = transport.NewGRPCServer(endpoints)
	)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()

		httpListener, _ := net.Listen("tcp", ":8081")
		http.Serve(httpListener, httpHandler)
		fmt.Println("http finish")
	}()
	go func() {
		defer wg.Done()

		grpcListener, _ := net.Listen("tcp", ":8082")
		baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
		pb.RegisterTagServiceServer(baseServer, grpcServer)
		baseServer.Serve(grpcListener)
		fmt.Println("grpc finish")
	}()

	wg.Wait()
}
