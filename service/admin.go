package service

import (
	"code/irisCms/model"
	"github.com/go-xorm/xorm"
)

/**
 * 管理员服务
 * 标准的开发模式将每个实体的提供的功能以接口标准的形式定义,供控制层进行调用。
 *
 */
type AdminService interface {
	//通过管理员用户名+密码 获取管理员实体 如果查询到，返回管理员实体，并返回true
	//否则 返回 nil ，false
	GetByAdminNameAndPassword(username, password string) (model.Admin, bool)
	GetByAdminId(adminId int64) (model.Admin, bool)
	//获取管理员总数
	GetAdminCount() (int64, error)
}

func NewAdminService(db *xorm.Engine) AdminService {
	return &adminService{
		engine: db,
	}
}

/**
 * 管理员的服务实现结构体
 */
type adminService struct {
	engine *xorm.Engine
}

/**
 * 查询管理员总数
 */
func (ac *adminService) GetAdminCount() (int64, error) {
	count, err := ac.engine.Count(new(model.Admin))

	if err != nil {
		panic(err.Error())
		return 0, err
	}

	return count, err
}

/**
 * 通过用户名和密码查询管理员
 */
func (ac *adminService) GetByAdminNameAndPassword(username, password string) (model.Admin, bool) {
	var admin model.Admin

	ac.engine.Where(" admin_name = ? and pwd = ? ", username, password).Get(&admin)

	return admin, admin.Id != 0
}

/**
 * 查询管理员信息
 */
func (ac *adminService) GetByAdminId(adminId int64) (model.Admin, bool) {
	var admin model.Admin

	ac.engine.Id(adminId).Get(&admin)

	return admin, admin.Id != 0
}