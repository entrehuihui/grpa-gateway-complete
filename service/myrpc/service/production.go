package service

import (
	"context"
)
	
// 列表
func (s Service) ProductionGet(ctx context.Context, in *proto.ProductionGetReq) (*proto.ProductionGetResp, error){
	return operate.ProductionGet(ctx, in)
}

// 修改
func (s Service) ProductionPut(ctx context.Context, in *proto.ProductionPutReq) (*proto.ProductionPutResp, error){
	return operate.ProductionPut(ctx, in)
}

// 新增
func (s Service) ProductionPost(ctx context.Context, in *proto.ProductionPostReq) (*proto.ProductionPostResp, error){
	return operate.ProductionPost(ctx, in)
}


func (s Service) ProductionDel(ctx context.Context, in *proto.ProductionDelReq) (*proto.ProductionDelResp, error){
	return operate.ProductionDel(ctx, in)
}

