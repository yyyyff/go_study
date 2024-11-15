package user

import "context"

// 面向对象
// user.Service
// 接口定义，一定要考虑兼容性,接口的参数不能变
// 中间件参数,取消/Trace/... 怎么产生怎么传递？
// 用context传递非业务需求
type Service interface {
	// 用户创建
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// 查询用户列表
	QueryUser(context.Context, *QueryUserRequest) (*UserSet, error)
	// 查询用户详情
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)

	// 作业：
	// 用户修改
	// 用户删除
}

type CreateUserRequest struct {
	Username string
	Password string
	Role     string
	Label    map[string]string
}

type QueryUserRequest struct {
	// 分页大小，一页多少个
	PageSize int
	//当前页，查询哪一页的数据
	PageNumber int
	// 查询指定用户的 用户列表[]
	Username string
}

type DescribeUserRequest struct {
	UserId int
}
type UserSet struct {
	// 总共多少个
	Total int64
	// 当前查询的数据清单
	Items []*User
}
