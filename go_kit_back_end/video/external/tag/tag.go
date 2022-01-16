package tag

import (
	"time"

	tagService "github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/service"
	client "github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/transport"
	"google.golang.org/grpc"
)

var (
	Client tagService.TagService
)

func Init() {
	conn, _ := grpc.Dial("127.0.0.1:8082", grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	Client = client.NewGRPCClient(conn)
}
