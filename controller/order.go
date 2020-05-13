package controller

import (
	"code/irisCms/service"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"code/irisCms/utils"
	//"strconv"
)

type OrderController struct {
	Ctx     iris.Context
	Service service.OrderService
	Session *sessions.Session
}

/**
 * 查询订单记录总数
 */
func (orderController *OrderController) GetCount() mvc.Result {
	iris.New().Logger().Info(" 查询订单记录总数 ")

	count, err := orderController.Service.GetCount()
	if err != nil {
		return mvc.Response{
			Object:map[string]interface{}{
				"status": utils.RECODE_FAIL,
				"count":  0,
			},
		}
	}

	return mvc.Response{
		Object:map[string]interface{}{
			"status": utils.RECODE_OK,
			"count": count,
		},
	}
}