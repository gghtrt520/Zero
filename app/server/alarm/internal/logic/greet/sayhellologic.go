package greetlogic

import (
	"context"

	"zlei/app/server/alarm/internal/svc"
	"zlei/app/server/alarm/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type SayHelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloLogic {
	return &SayHelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayHelloLogic) SayHello(in *proto.HelloReq) (*proto.HelloResp, error) {
	// todo: add your logic here and delete this line

	return &proto.HelloResp{}, nil
}
