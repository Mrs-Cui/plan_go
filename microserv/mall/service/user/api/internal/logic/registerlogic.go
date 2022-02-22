package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	svc2 "plan_go/plan_go/microserv/mall/service/user/api/internal/svc"
	types2 "plan_go/plan_go/microserv/mall/service/user/api/internal/types"
	userclient2 "plan_go/plan_go/microserv/mall/service/user/rpc/userclient"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc2.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc2.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types2.RegisterRequest) (resp *types2.RegisterResponse, err error) {
	res, err := l.svcCtx.UserRpc.Register(l.ctx, &userclient2.RegisterRequest{
		Name:     req.Name,
		Gender:   req.Gender,
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types2.RegisterResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil
}