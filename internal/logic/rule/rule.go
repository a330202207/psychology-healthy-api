// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/14 16:02
// @Package logic

package rule

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/a330202207/psychology-healthy-api/internal/dao"
	"github.com/a330202207/psychology-healthy-api/internal/model"
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
func (s *sRule) UpdateRuleByIds(ctx context.Context, memberId int64, ruleIds []int64, tx *gdb.TX) (err error) {
	var ruleData []map[string]interface{}

	if _, err = dao.SysMemberRole.Ctx(ctx).TX(tx).Where("member_id", memberId).Delete(); err != nil {
		err = gerror.Wrap(err, "del rule err!")
		return err
	}

	for _, val := range ruleIds {
		var data = make(map[string]interface{})
		data["member_id"] = memberId
		data["rule_id"] = val

		ruleData = append(ruleData, data)
	}

	if _, err = dao.SysMemberRole.Ctx(ctx).TX(tx).Insert(ruleData); err != nil {
		err = gerror.Wrap(err, "batch insert rule err")
		return err
	}

	return
}

// Edit 添加/编辑角色
func (s *sRule) Edit(ctx context.Context, in *model.RuleEditInput) (err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-rule-Edit")
	defer span.End()
	var logger = gconv.String(ctx.Value("logger"))

	ok, err := dao.SysRole.IsUniqueName(ctx, in.Name)
	if err != nil {
		g.Log(logger).Error(ctx, "service Rule EditIsUniqueName error:", err.Error())
		err = errors.New("操作失败[001]")
		return err
	}

	if ok {
		g.Log(logger).Error(ctx, "service Rule name exist")
		err = errors.New("角色名称已存在")
		return err
	}

	tx, err := g.DB().Begin(ctx)
	if err != nil {
		g.Log(logger).Error(ctx, "service Rule Del Transaction error:", err.Error())
		err = errors.New("操作失败[001]")
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}

	}()

	if in.ID > 0 {
		if _, err = dao.SysRole.Ctx(ctx).TX(tx).Where("id", in.ID).OmitEmpty().Update(in); err != nil {
			g.Log(logger).Error(ctx, "service Rule Edit update mysql error:", err.Error())
			err = errors.New("操作失败[002]")
			return
		}
	} else {
		in.ID, err = dao.SysRole.Ctx(ctx).TX(tx).OmitEmpty().InsertAndGetId(in)
		if err != nil {
			g.Log(logger).Error(ctx, "service Rule Edit insert mysql error:", err.Error())
			err = errors.New("操作失败[003]")
			return
		}
	}

	return
}

// Del .删除角色
func (s *sRule) Del(ctx context.Context, in *model.RuleBaseInput) (err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-rule-Del")
	defer span.End()
	var logger = gconv.String(ctx.Value("logger"))

	tx, err := g.DB().Begin(ctx)
	if err != nil {
		g.Log(logger).Error(ctx, "service Rule Del Transaction error:", err.Error())
		err = errors.New("操作失败[001]")
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	if _, err = dao.SysRole.Ctx(ctx).TX(tx).Delete("id", in.ID); err != nil {
		g.Log(logger).Error(ctx, "service Rule Del delete mysql error:", err.Error())
		err = errors.New("操作失败[002]")
		return
	}

	if _, err = dao.SysRoleMenu.Ctx(ctx).TX(tx).Delete("role_id", in.ID); err != nil {
		g.Log(logger).Error(ctx, "service RuleMenu Del delete mysql error:", err.Error())
		err = errors.New("操作失败[003]")
		return
	}

	return
}
