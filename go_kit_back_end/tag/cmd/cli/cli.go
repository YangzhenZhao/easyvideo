package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/tag/transport"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8082", grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	svc := transport.NewGRPCClient(conn)
	// tags := svc.GetVideoTags(context.Background(), 11)
	tags := svc.GetVideoTags(context.Background(), 2333)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	// 	os.Exit(1)
	// }
	fmt.Fprintf(os.Stdout, "%+v\n", tags)
}
