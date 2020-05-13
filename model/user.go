package model

import (
	"time"
	"code/irisCms/utils"
)

/**
 * 用户信息结构体,用于生成用户信息表
 */
type User struct {
	Id           int64    `xorm:"pk autoincr" json:"id"`
	UserName     string   `xorm:"varchar(12)" json:"username"`
	RegisterTime time.Time `json:"register_time"`
	Mobile       string    `xorm:"varchar(11)" json:"mobile"`
	IsActive     int64     `json:"is_active"`
	Balance      int64     `json:"balance"`   //用户的账户余额（简单起见，使用int类型）
	Avatar       string    `xorm:"varchar(255)" json:"avatar"`
	Pwd          string    `json:"password"`
	DelFlag      int64     `json:"del_flag"`
	CityName     string    `xorm:"varchar(24)" json:"city_name"`
	city         *City    `xorm:"- <- ->"`
}


/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (user *User) UserToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":           user.Id,
		"user_id":      user.Id,
		"username":     user.UserName,
		"city":         user.CityName,
		"registe_time": utils.FormatDatetime(user.RegisterTime),
		"mobile":       user.Mobile,
		"is_active":    user.IsActive,
		"balance":      user.Balance,
		"avatar":       user.Avatar,
	}

	return respInfo
}