// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/14 16:02
// @Package logic

package rule

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/a330202207/psychology-healthy-api/internal/dao"
	"github.com/a330202207/psychology-healthy-api/internal/service"
)

type sRule struct {
}

func init() {
	service.RegisterRule(New())
}

func New() *sRule {
	return &sRule{}
}

// UpdateRuleByIds 更新用户角色
func (s *sRule) UpdateRuleByIds(ctx context.Context, memberId int64, ruleIds []int64) (err error) {
	var (
		ruleData []map[string]interface{}
	)

	if _, err = dao.SysMemberRole.Ctx(ctx).Where("member_id", memberId).Delete(); err != nil {
		err = gerror.Wrap(err, "del rule err!")
		return err
	}

	for _, val := range ruleIds {
		var data = make(map[string]interface{})
		data["member_id"] = memberId
		data["rule_id"] = val

		ruleData = append(ruleData, data)
	}

	if _, err = dao.SysMemberRole.Ctx(ctx).Insert(ruleData); err != nil {
		err = gerror.Wrap(err, "batch insert rule err")
		return err
	}

	return
}
