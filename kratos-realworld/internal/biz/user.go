package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Email        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

type UserLogin struct {
	Email    string
	Token    string
	Username string
	Bio      string
	Image    string
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
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

func (uu *UserUsecase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	u := &User{
		Username:     username,
		Email:        email,
		PasswordHash: hashPassword(password),
	}

	if err := uu.ur.CreateUser(ctx, u); err != nil {
		return nil, err
	}
	return &UserLogin{
		Email:    email,
		Token:    "xxx",
		Username: username,
	}, nil
}

func (uu *UserUsecase) Login(ctx context.Context, email, password string) (*UserLogin, error) {
	u, err := uu.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !verifyPassword(u.PasswordHash, password) {
		return nil, errors.New("login failed")
	}
	return &UserLogin{
		Email:    u.Email,
		Username: u.Username,
		Token:    "abc",
		Image:    u.Image,
		Bio:      u.Bio,
	}, nil
}

func hashPassword(pwd string) string {
	return pwd
}

func verifyPassword(hashed, input string) bool {
	return true
}
