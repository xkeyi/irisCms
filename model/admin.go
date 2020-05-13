package model

import (
	"time"
)

// 定义管理员结构体
type Admin struct {
	Id         int64    `xorm:"pk autoincr" json:"id"`
	AdminName  string    `xorm:"varchar(32)" json:"admin_name"`
	CreateTime time.Time `xorm:"DateTime" json:"create_time"`
	Status     int64     `xorm:"default 0" json:"status"`
	Avatar     string    `xorm:"varchar(255)" json:"avatar"`
	Pwd        string    `xorm:"varchar(255)" json:"pwd"`      //管理员密码
	CityName   string    `xorm:"varchar(12)" json:"city_name"` //管理员所在城市名称
	CityId     int64     `xorm:"index" json:"city_id"`
	City       *City     `xorm:"- <- ->"` //所对应的城市结构体（基础表结构体）
}

/**
 * 从Admin数据库实体转换为前端请求的resp的json格式
 */
func (this *Admin) AdminToRespDesc() interface{} {
	respDesc := map[string]interface{}{
		"user_name":   this.AdminName,
		"id":          this.Id,
		"create_time": this.CreateTime,
		"status":      this.Status,
		"avatar":      this.Avatar,
		"city":        this.CityName,
		"admin":       "管理员",
	}
	return respDesc
}