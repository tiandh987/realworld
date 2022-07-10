package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

// RealWorldService is a realworld service.
type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uu  *biz.UserUsecase
	log *log.Helper
}

// NewRealWorldService new a realworld service.
func NewRealWorldService(uu *biz.UserUsecase, logger log.Logger) *RealWorldService {
	return &RealWorldService{uu: uu, log: log.NewHelper(logger)}
}
