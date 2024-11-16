package impl

import (
	"context"

	"github.com/yyyyff/go_study/tree/main/vblog/apps/user"
)

var _ user.Service = (*UserServiceImpl)(nil)

// 实现user.Service

func (i *UserServiceImpl) CreateUser(ctx context.Context, in *user.CreateUserRequest) (*user.User, error) {
	return nil, nil
}

// QueryUser 查询用户列表
func (i *UserServiceImpl) QueryUser(ctx context.Context, in *user.QueryUserRequest) (*user.UserSet, error) {
	return nil, nil
}

// DescribeUser 查询用户详情
func (i *UserServiceImpl) DescribeUser(ctx context.Context, in *user.DescribeUserRequest) (*user.User, error) {
	return nil, nil
}
