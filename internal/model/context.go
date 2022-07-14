// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/14 12:12
// @Package model

package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Context 请求上下文结构
type Context struct {
	ReqId      string         // 上下文ID
	Module     string         // 应用模块
	TakeUpTime int64          // 请求耗时 ms
	Request    *ghttp.Request // 当前Request管理对象
	User       *ContextUser   // 上下文用户信息
	Data       g.Map          // 自定KV变量，业务模块根据需要设置，不固定
}

// ContextUser .
type ContextUser struct {
	ID         int64  `json:"id"`          // 会员ID
	Username   string `json:"username"`    // 用户名
	Nickname   string `json:"nick_name"`   // 昵称
	Avatar     string `json:"avatar"`      // 头像
	Email      string `json:"email"`       // 邮箱
	Mobile     string `json:"mobile"`      // 手机
	VisitCount uint   `json:"visit_count"` // 访问次数
	Type       uint   `json:"Type"`        // 账户类型:1-普通管理员，10-超级管理员
	LastTime   int    `json:"last_time"`   // 最后一次登录时间
	LastIp     string `json:"last_ip"`     // 最后一次登录ip
	Exp        int64  `json:"exp"`         // 登录有效期截止时间戳
	Expires    int64  `json:"expires"`     // 登录有效期
	App        string `json:"app"`         // 登录应用
}
