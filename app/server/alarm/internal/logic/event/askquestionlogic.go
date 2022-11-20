package eventlogic

import (
	"context"

	"zlei/app/server/alarm/internal/svc"
	"zlei/app/server/alarm/proto"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type AskQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAskQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AskQuestionLogic {
	return &AskQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AskQuestionLogic) AskQuestion(in *proto.EventReq) (*proto.EventResp, error) {
	names := lo.Uniq([]string{"Samuel", "John", "Samuel"})
	logx.Info(names)
	return &proto.EventResp{}, nil
}
