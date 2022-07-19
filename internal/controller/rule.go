// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/14 15:56
// @Package controller

package controller

import (
	"context"

	v1 "github.com/a330202207/psychology-healthy-api/api/v1"
)

var Rule = cRule{}

type cRule struct {
}

// 添加/编辑角色
func (c *cRule) Edit(ctx context.Context, input *v1.RuleEditReq) (res *v1.RuleEditRes, err error) {

	return
}

// 删除角色

// 角色列表
