package main

import (
	"flag"
	"fmt"

	"zlei/app/server/alarm/internal/config"
	eventServer "zlei/app/server/alarm/internal/server/event"
	greetServer "zlei/app/server/alarm/internal/server/greet"
	"zlei/app/server/alarm/internal/svc"
	"zlei/app/server/alarm/proto"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/alarm.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		proto.RegisterGreetServer(grpcServer, greetServer.NewGreetServer(ctx))
		proto.RegisterEventServer(grpcServer, eventServer.NewEventServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
