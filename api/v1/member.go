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

// MemberAddReq .
type MemberAddReq struct {
	g.Meta     `path:"/admin/add" method:"post" tag:"memberService" summary:"添加管理员" tags:"管理员管理"`
	Account    string `json:"account" v:"required#用户名不能为空" dc:"用户名"`
	Passwd     string `json:"passwd" v:"required|length:4,20#密码不能为空|密码长度应当在4到20之间" dc:"密码"`
	Type       uint   `json:"type" dc:"账户类型:1-普通管理员，10-超级管理员"`
	Status     uint   `json:"status" dc:"状态:10-正常，20-关闭，30-待验证"`
	NickName   string `json:"nickName" dc:"昵称"`
	Avatar     string `json:"avatar" dc:"头像"`
	Sex        uint   `json:"sex" dc:"性别:0-未知，1-男，2-女"`
	Email      string `json:"email" dc:"邮箱"`
	ProvinceID uint   `json:"provinceId" dc:"省"`
	CityID     uint   `json:"cityId" dc:"城市"`
	AreaID     uint   `json:"areaId" dc:"地区"`
	Address    string `json:"address" dc:"默认地址"`
	Mobile     string `json:"mobile" dc:"手机号码"`
	Birthday   string `json:"birthday" dc:"生日"`
}

type MemberAddRes struct {
}
