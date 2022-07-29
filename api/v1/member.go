// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/11 17:27
// @Package v1

package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/a330202207/psychology-healthy-api/internal/model"
)

// MemberEditReq .
type MemberEditReq struct {
	g.Meta     `path:"/admin/add" method:"post" tag:"memberService" summary:"添加管理员" tags:"管理员管理"`
	ID         int64   `json:"id"`
	RuleIds    []int64 `json:"ruleIds" v:"required#角色不能为空" dc:"角色ID"`
	UserName   string  `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Passwd     string  `json:"passwd" v:"required|length:4,20#密码不能为空|密码长度应当在4到20之间" dc:"密码"`
	Type       uint    `json:"type" dc:"账户类型:1-普通管理员，10-超级管理员"`
	Status     uint    `json:"status" dc:"状态:10-正常，20-关闭，30-待验证"`
	NickName   string  `json:"nickName" dc:"昵称"`
	Avatar     string  `json:"avatar" dc:"头像"`
	Sex        uint    `json:"sex" dc:"性别:0-未知，1-男，2-女"`
	Email      string  `json:"email" dc:"邮箱"`
	ProvinceID uint    `json:"provinceId" dc:"省"`
	CityID     uint    `json:"cityId" dc:"城市"`
	AreaID     uint    `json:"areaId" dc:"地区"`
	Address    string  `json:"address" dc:"地址"`
	Mobile     string  `json:"mobile" dc:"手机号码"`
	Birthday   string  `json:"birthday" dc:"生日"`
}

// MemberEditRes .
type MemberEditRes struct {
	g.Meta `mime:"application/json"`
}

// MemberUpdatePassWdReq .
type MemberUpdatePassWdReq struct {
	g.Meta `path:"/admin/updatePassWd" method:"post" tag:"memberService" summary:"修改密码" tags:"修改密码"`
	ID     int64  `json:"id"`
	Passwd string `json:"passwd" v:"required|length:4,20#密码不能为空|密码长度应当在4到20之间" dc:"密码"`
}

// MemberUpdatePassWdRes .
type MemberUpdatePassWdRes struct {
	g.Meta `mime:"application/json"`
}

// MemberResetPassWdReq .
type MemberResetPassWdReq struct {
	g.Meta     `path:"/admin/resetPassWd" method:"post" tag:"memberService" summary:"重置密码" tags:"重置密码"`
	Type       uint   `json:"type" dc:"类型:10-手机，20-邮箱"`
	Account    string `json:"account" dc:"手机(邮箱)"`
	VerifyCode string `v:"required#验证码不能为空" json:"verifyCode" dc:"验证码"`
	Passwd     string `v:"required|length:4,20#密码不能为空|密码长度应当在4到20之间" json:"passwd" dc:"密码"`
}

// MemberResetPassWdRes .
type MemberResetPassWdRes struct {
	g.Meta `mime:"application/json"`
}

// MemberDelReq .
type MemberDelReq struct {
	g.Meta `path:"/admin/del" method:"post" tag:"memberService" summary:"删除管理员" tags:"删除管理员"`
	ID     int64 `v:"required#管理员ID不能为空" json:"id"`
}

// MemberDelRes .
type MemberDelRes struct {
	g.Meta `mime:"application/json"`
}

// MemberListReq .
type MemberListReq struct {
	g.Meta   `path:"/admin/list" method:"get" tag:"memberService" summary:"管理员列表" tags:"管理员列表"`
	UserName string `json:"username" dc:"用户名"`
	Type     uint   `json:"type" dc:"账户类型:1-普通管理员，10-超级管理员"`
	Status   uint   `json:"status" dc:"状态:10-正常，20-关闭，30-待验证"`
	NickName string `json:"nickName" dc:"昵称"`
	Mobile   string `json:"mobile" dc:"手机号码"`
	Sex      uint   `json:"sex" dc:"性别:0-未知，1-男，2-女"`
	Email    string `json:"email" dc:"邮箱"`
}

// MemberListRes .
type MemberListRes struct {
	g.Meta   `mime:"application/json"`
	List     []*model.MemberListOut `json:"list"`
	PageInfo *PageInfo              `json:"pageInfo"`
}

// MemberGetReq .
type MemberGetReq struct {
	g.Meta `path:"/admin/get" method:"get" tag:"memberService" summary:"获取管理员信息" tags:"获取管理员信息"`
	ID     int64 `v:"required#管理员ID不能为空" json:"id"`
}

// MemberGetRes .
type MemberGetRes struct {
	g.Meta     `mime:"application/json"`
	ID         int64   `json:"id"`
	RuleIds    []int64 `json:"ruleIds" dc:"角色ID"`
	UserName   string  `json:"username" dc:"用户名"`
	Type       uint    `json:"type" dc:"账户类型:1-普通管理员，10-超级管理员"`
	Status     uint    `json:"status" dc:"状态:10-正常，20-关闭，30-待验证"`
	NickName   string  `json:"nickName" dc:"昵称"`
	Avatar     string  `json:"avatar" dc:"头像"`
	Sex        uint    `json:"sex" dc:"性别:0-未知，1-男，2-女"`
	Email      string  `json:"email" dc:"邮箱"`
	ProvinceID uint    `json:"provinceId" dc:"省"`
	CityID     uint    `json:"cityId" dc:"城市"`
	AreaID     uint    `json:"areaId" dc:"地区"`
	Address    string  `json:"address" dc:"地址"`
	Mobile     string  `json:"mobile" dc:"手机号码"`
	Birthday   string  `json:"birthday" dc:"生日"`
}
