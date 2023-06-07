package operate

import (
	"context"
	"test/service/myrpc/proto"
)
	
// 设备组列表
func GroupGet(ctx context.Context, in *proto.GroupGetReq) (*proto.GroupGetResp, error){
	return nil,nil
}

// 增加设备组
func GroupPost(ctx context.Context, in *proto.GroupPostReq) (*proto.GroupPostResp, error){
	return nil,nil
}

// 删除设备组
func GroupDel(ctx context.Context, in *proto.GroupDelReq) (*proto.GroupDelResp, error){
	return nil,nil
}

// 修改设备组
func GroupPut(ctx context.Context, in *proto.GroupPutReq) (*proto.GroupPutResp, error){
	return nil,nil
}

