// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/14 16:02
// @Package logic

package rule

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
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

// Edit 添加/编辑角色
func (s *sRule) Edit(ctx context.Context, in *model.RuleEditInput) (err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-rule-Edit")
	defer span.End()
	var logger = gconv.String(ctx.Value("logger"))

	ok, err := dao.SysRole.IsUniqueName(ctx, in.Name)
	if err != nil {
		g.Log(logger).Error(ctx, "service Rule EditIsUniqueName error:", err.Error())
		err = gerror.NewCode(gcode.New(500, "", nil), "操作失败[01]")
		return err
	}

	if ok {
		g.Log(logger).Error(ctx, "service Rule name exist")
		err = gerror.NewCode(gcode.New(500, "", nil), "角色名称已存在")
		return err
	}

	tx, err := g.DB().Begin(ctx)
	if err != nil {
		g.Log(logger).Error(ctx, "service Rule Del Transaction error:", err.Error())
		err = gerror.NewCode(gcode.New(500, "", nil), "操作失败[02]")
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
			err = gerror.NewCode(gcode.New(500, "", nil), "操作失败[03]")
			return
		}
	} else {
		in.ID, err = dao.SysRole.Ctx(ctx).TX(tx).OmitEmpty().InsertAndGetId(in)
		if err != nil {
			g.Log(logger).Error(ctx, "service Rule Edit insert mysql error:", err.Error())
			err = gerror.NewCode(gcode.New(500, "", nil), "操作失败[04]")
			return
		}
	}

	if err = dao.SysRoleMenu.UpdateRoleMenuByIds(ctx, in.ID, in.MenuIds, tx); err != nil {
		g.Log(logger).Error(ctx, "service Rule Edit update UpdateRoleMenuByIds error:", err.Error())
		err = gerror.NewCode(gcode.New(500, "", nil), "操作失败[05]")
		return
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
		err = gerror.NewCode(gcode.New(500, "", nil), "操作失败[01]")
		return
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
		err = gerror.NewCode(gcode.New(500, "", nil), "操作失败[02]")
		return
	}

	if _, err = dao.SysRoleMenu.Ctx(ctx).TX(tx).Delete("role_id", in.ID); err != nil {
		g.Log(logger).Error(ctx, "service RuleMenu Del delete mysql error:", err.Error())
		err = gerror.NewCode(gcode.New(500, "", nil), "操作失败[03]")
		return
	}

	return
}

// List 获取角色列表
func (s *sRule) List(ctx context.Context, in *model.RuleListInput) (out *model.RuleListOut, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-rule-list")
	defer span.End()
	var logger = gconv.String(ctx.Value("logger"))
	m := dao.SysRole.Ctx(ctx)
	if in.ID > 0 {
		m.Where("id = ?", in.ID)
	}

	if in.Name != "" {
		m.Where("name = ?", in.Name)
	}

	if in.Status > 0 {
		m.Where("status = ?", in.Status)
	}
	out.Total, err = m.Count()
	if err != nil {
		g.Log(logger).Error(ctx, "service Rule list count error:", err.Error())
		err = gerror.NewCode(gcode.New(500, "", nil), "获取角色数据失败[01]")
		return
	}
	out.Page = in.Page
	out.PageSize = in.PageSize

	if err = m.Page(in.Page, in.PageSize).Order("sort asc,id asc").Scan(&out.List); err != nil {
		g.Log(logger).Error(ctx, "service Rule list scan error:", err.Error())
		err = gerror.NewCode(gcode.New(500, "", nil), "获取角色数据失败[02]")
		return
	}

	return
}

// GetAll 获取所有角色
func (s sRule) GetAll(ctx context.Context) (out []*model.RuleItem, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-rule-list")
	defer span.End()
	var logger = gconv.String(ctx.Value("logger"))

	if err = dao.SysRole.Ctx(ctx).Where("status = 10").Order("sort asc,id asc").Scan(&out); err != nil {
		g.Log(logger).Error(ctx, "service GetAll list scan error:", err.Error())
		err = gerror.NewCode(gcode.New(500, "", nil), "获取角色数据失败")
		return
	}
	return
}
