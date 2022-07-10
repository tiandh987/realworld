package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Username string
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
}

type ProfileRepo interface {
}

type UserUsecase struct {
	ur UserRepo
	pr ProfileRepo
	tr TagRepo

	log *log.Helper
}

func NewUserUsecase(ur UserRepo, pr ProfileRepo, tr TagRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		ur:  ur,
		pr:  pr,
		tr:  tr,
		log: log.NewHelper(logger),
	}
}

func (uu *UserUsecase) Register(ctx context.Context, u *User) error {
	if err := uu.ur.CreateUser(ctx, u); err != nil {

	}
	return nil
}
