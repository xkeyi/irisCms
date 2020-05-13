package service

import (
	"github.com/go-xorm/xorm"
	"code/irisCms/model"
	"github.com/kataras/iris"
)

/**
 * 订单服务接口
 */
type OrderService interface {
	GetCount() (int64, error)
	GetOrderList(offset, limit int) []model.OrderDetail
}

/**
 * 订单服务
 */
type orderService struct {
	Engine *xorm.Engine
}

func (orderService *orderService) GetOrderList(offset, limit int) []model.OrderDetail {
	orderList := make([]model.OrderDetail, 0)

	iris.New().Logger().Info(orderList)
	return orderList
}

/**
 * 实例化OrderService服务对象
 */
func NewOrderService (db *xorm.Engine) OrderService {
	return &orderService{Engine: db}
}

/**
 * 获取订单总数量
 */
func (orderService *orderService) GetCount() (int64, error) {
	count, err := orderService.Engine.Where(" del_flag = 0 ").Count(new(model.UserOrder))
	if err != nil {
		return 0, err
	}

	return count, nil
}