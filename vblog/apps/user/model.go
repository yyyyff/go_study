package user

// 存放需要入库的数据结构 （java中叫PO）

// 用户创建成功后返回一个User对象
type User struct {
	Id        int
	CreatedAt int64
	UpdatedAt int64

	*CreateUserRequest
}
