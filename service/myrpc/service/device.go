package service

import (
	"context"
	"test/service/myrpc/proto"
	"test/service/operate"
)
	
// 设备注册
func (s Service) DeviceReg(ctx context.Context, in *proto.DeviceRegReq) (*proto.DeviceResp, error){
	return operate.DeviceReg(ctx, in)
}

// 设备列表
func (s Service) DeviceGet(ctx context.Context, in *proto.DeviceGetReq) (*proto.DeviceGetResp, error){
	return operate.DeviceGet(ctx, in)
}

