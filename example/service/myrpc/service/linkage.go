package service

import (
	"context"
	"test/service/myrpc/proto"
	"test/service/operate"
)
	
// 联动组列表
func (s Service) LinkageGet(ctx context.Context, in *proto.LinkageGetReq) (*proto.LinkageGetResp, error){
	return operate.LinkageGet(ctx, in)
}

// 增加联动组
func (s Service) LinkagePost(ctx context.Context, in *proto.LinkagePostReq) (*proto.LinkagePostResp, error){
	return operate.LinkagePost(ctx, in)
}

// 删除联动组
func (s Service) LinkageDel(ctx context.Context, in *proto.LinkageDelReq) (*proto.LinkageDelResp, error){
	return operate.LinkageDel(ctx, in)
}

// 修改联动组
func (s Service) LinkagePut(ctx context.Context, in *proto.LinkagePutReq) (*proto.LinkagePutResp, error){
	return operate.LinkagePut(ctx, in)
}

