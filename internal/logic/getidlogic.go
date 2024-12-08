package logic

import (
	"context"

	"snowflake_id_server/internal/svc"
	"snowflake_id_server/snowflake_id_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIDLogic {
	return &GetIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetIDLogic) GetID(in *snowflake_id_server.Request) (*snowflake_id_server.GetIDResp, error) {
	id, err := l.svcCtx.FlakegenNode.GetID()
	if err != nil {
		return nil, err
	}
	resp := &snowflake_id_server.GetIDResp{
		RespID: in.ReqID,
		ID:     id,
	}
	return resp, nil
}
