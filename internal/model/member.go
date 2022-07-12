// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/11 17:41
// @Package model

package model

// MemberAddInput .
type MemberAddInput struct {
	Account string // 登陆用户名
	Passwd  string // 登陆密码
	Type    uint   // 账户类型
	Status  uint   // 状态:10-开启-20关闭
}

type MemberAddOutput struct {
}
