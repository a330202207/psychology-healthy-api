// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/11 17:27
// @Package v1

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AdminAddReq .
type AdminAddReq struct {
	g.Meta  `path:"/admin/add" method:"post" tag:"adminService" summary:"添加管理员" tags:"管理员管理"`
	Account string `json:"account" v:"required#用户名不能为空" dc:"用户名"`
	Passwd  string `json:"passwd" v:"required|length:4,20#密码不能为空|密码长度应当在4到20之间" dc:"密码"`
	Type    uint   `json:"type" dc:"账户类型"`
	Status  uint   `json:"status" dc:"状态"`
}

type AdminAddRes struct {
}
