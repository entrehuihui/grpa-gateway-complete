package service

import (
	"context"
	"test/service/myrpc/proto"
	"test/service/operate"
)
	
// 设备组列表
func (s Service) GroupGet(ctx context.Context, in *proto.GroupGetReq) (*proto.GroupGetResp, error){
	return operate.GroupGet(ctx, in)
}

// 增加设备组
func (s Service) GroupPost(ctx context.Context, in *proto.GroupPostReq) (*proto.GroupPostResp, error){
	return operate.GroupPost(ctx, in)
}

// 删除设备组
func (s Service) GroupDel(ctx context.Context, in *proto.GroupDelReq) (*proto.GroupDelResp, error){
	return operate.GroupDel(ctx, in)
}

// 修改设备组
func (s Service) GroupPut(ctx context.Context, in *proto.GroupPutReq) (*proto.GroupPutResp, error){
	return operate.GroupPut(ctx, in)
}

