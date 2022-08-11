// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/11 17:41
// @Package model

package model

import (
	"github.com/a330202207/psychology-healthy-api/internal/model/entity"
)

// MemberEditInput .
type MemberEditInput struct {
	ID         int64
	RoleIds    []int64 // 角色ID
	UserName   string  // 登陆用户名
	Passwd     string  // 登陆密码
	Type       uint    // 账户类型:1-普通管理员，10-超级管理员"
	Status     uint    // 状态:10-正常，20-关闭，30-待验证
	NickName   string  // 昵称
	Avatar     string  // 头像
	Sex        uint    // 性别:0-未知，1-男，2-女
	Email      string  // 邮箱
	ProvinceID uint    // 省
	CityID     uint    // 城市
	AreaID     uint    // 地区
	Address    string  // 地址
	Mobile     string  // 手机
	Birthday   string  // 生日
}

// MemberUpdatePassWdInput .
type MemberUpdatePassWdInput struct {
	ID     int64
	Passwd string
}

// MemberBaseInput .
type MemberBaseInput struct {
	ID int64
}

// MemberResetPassWdInput .
type MemberResetPassWdInput struct {
	Type       uint
	Account    string
	VerifyCode string
	Passwd     string
}

// MemberListInput .
type MemberListInput struct {
}

// MemberListOut .
type MemberListOut struct {
	*entity.SysMember
}
