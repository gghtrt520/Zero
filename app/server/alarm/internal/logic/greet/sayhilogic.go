package greetlogic

import (
	"context"

	"zlei/app/server/alarm/internal/svc"
	"zlei/app/server/alarm/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type SayHiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHiLogic {
	return &SayHiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayHiLogic) SayHi(in *proto.HiReq) (*proto.HiResp, error) {
	// todo: add your logic here and delete this line

	return &proto.HiResp{}, nil
}
