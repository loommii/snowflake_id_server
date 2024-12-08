package logic

import (
	"context"
	"time"

	"snowflake_id_server/internal/svc"
	"snowflake_id_server/snowflake_id_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *snowflake_id_server.Request) (*snowflake_id_server.PingResp, error) {
	resp := &snowflake_id_server.PingResp{
		RespID:  in.ReqID,
		Message: "ok",
		Time:    time.Now().Local().String(),
	}
	return resp, nil
}
