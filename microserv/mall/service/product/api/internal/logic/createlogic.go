package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"plan_go/microserv/mall/service/product/api/internal/svc"
	"plan_go/microserv/mall/service/product/api/internal/types"
	"plan_go/microserv/mall/service/product/rpc/product"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateLogic {
	return CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req types.CreateRequest) (resp *types.CreateResponse, err error) {
	res, err := l.svcCtx.ProductRpc.Create(l.ctx, &product.CreateRequest{
		Name:   req.Name,
		Desc:   req.Desc,
		Stock:  req.Stock,
		Amount: req.Amount,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{
		Id: res.Id,
	}, nil
}