// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/18 15:05
// @Package model

package model

// RuleEditInput .
type RuleEditInput struct {
	ID      int64
	Name    string
	MenuIds []int64 // 菜单ID
	Status  uint    // 状态:10-正常，20-关闭，30-待验证
	Sort    uint    // 排序
}

// RuleBaseInput .
type RuleBaseInput struct {
	ID int64
}
