package service

import (
	"code/irisCms/model"
	"fmt"
	"github.com/go-xorm/xorm"
	"math/rand"
	"time"
)

/**
 * 统计功能模块接口标准
 */
type StatisService interface {
	GetUserDailyCount(date string) int64
	GetOrderDailyCount(date string) int64
	GetAdminDailyCount(date string) int64
}

/**
 * 统计功能服务实现结构体
 */
type statisService struct {
	Engine *xorm.Engine
}

func NewStatisService(engine *xorm.Engine) StatisService {
	return &statisService{
		Engine: engine,
	}
}

func (ss statisService) GetUserDailyCount(date string) int64 {
	if date == "NaN-NaN-NaN" { //当日增长数据请求
		date = time.Now().Format("2006-01-02")
	}

	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}

	endDate := startDate.AddDate(0, 0, 1)
	result, err := ss.Engine.Where(" register_time between ? and ? and del_flag = 0 ", startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05")).Count(model.User{})
	if err != nil {
		return 0
	}
	fmt.Println(result)
	//return result
	return int64(rand.Intn(100))
}

func (ss statisService) GetOrderDailyCount(date string) int64 {
	if date == "NaN-NaN-NaN" { //当日增长数据请求
		date = time.Now().Format("2006-01-02")
	}

	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}

	//2019-4-23
	// 2019-4-23 00:00:00  --- 2019-04-23 23:59:59
	//                         2019-04-24 00:00:00

	endDate := startDate.AddDate(0, 0, 1)
	result, err := ss.Engine.Where(" time between ? and ? and del_flag = 0 ", startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05")).Count(model.UserOrder{})
	if err != nil {
		return 0
	}

	fmt.Println(result)
	//return result
	return int64(rand.Intn(100))
}

func (ss statisService) GetAdminDailyCount(date string) int64 {
	if date == "NaN-NaN-NaN" {
		date = time.Now().Format("2006-01-02")
	}

	//查询日期date格式解析
	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}

	endDate := startDate.AddDate(0, 0, 1)

	result, err := ss.Engine.Where(" create_time between ? and ? and status = 0 ",
		startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05")).Count(model.Admin{})

	if err != nil {
		return 0
	}

	fmt.Println(result)
	//return result
	return int64(rand.Intn(100))
}
