package service

import (
	"test/service/myrpc/proto"
)

// Service .
type Service struct {
	// ##继承
	proto.UnimplementedLoginServer
	// ##继承
}

// NewService .
func NewService() *Service {
	s := new(Service)
	return s
}
