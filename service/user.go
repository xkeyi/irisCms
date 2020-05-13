package service

import (
	"code/irisCms/model"
	"github.com/go-xorm/xorm"
)

/**
 * 用户模块功能服务接口
 */
type UserService interface {
	GetUserDailyStaticCount(datetime string) int64
	GetUserTotalCount() (int64, error)
	GetUserList(offset, limit int) []*model.User
}

/**
 * 实例化用户服务结构实体对象
 */
func NewUserService(engine *xorm.Engine) UserService {
	return &userService{
		Engine: engine,
	}
}

/**
 * 用户服务实现结构体
 */
type userService struct {
	Engine *xorm.Engine
}

func (us userService) GetUserDailyStaticCount(datetime string) int64 {
	result, err := us.Engine.Count(new(model.User))
	if err != nil {
		panic(err.Error())
	}

	return result
}

func (us userService) GetUserTotalCount() (int64, error) {
	//查询del_flag 为0 的用户的总数量；del_flag:0 正常状态；del_flag:1 用户注销或者被删除
	count, err := us.Engine.Where(" del_flag = ? ", 0).Count(new(model.User))
	if err != nil {
		panic(err.Error())
		return 0, err
	}

	return count, nil
}

func (us userService) GetUserList(offset, limit int) []*model.User {
	panic("implement me")
}
