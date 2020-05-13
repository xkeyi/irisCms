package controller

import (
	"code/irisCms/service"
	"code/irisCms/utils"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

/**
 * 用户控制器结构体：用来实现处理用户模块的接口的请求，并返回给客户端
 */
type UserController struct {
	Ctx       iris.Context
	Service   service.UserService
	Session   sessions.Session
}

/**
 * 获取用户总数
 * 请求类型：Get
 * 请求Url：/v1/users/count
 */
func (uc *UserController) GetCount() mvc.Result {
	total, err := uc.Service.GetUserTotalCount()

	if err != nil {
		return mvc.Response{
			Object:map[string]interface{}{
				"status": utils.RECODE_FAIL,
				"count":  0,
			},
		}
	}

	//正常情况的返回值
	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"count":  total,
		},
	}
}