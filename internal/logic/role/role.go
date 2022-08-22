// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/14 16:02
// @Package logic

package role

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

type sRole struct {
}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

// Edit 添加/编辑角色
func (s *sRole) Edit(ctx context.Context, in *model.RoleEditInput) (err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-role-Edit")
	defer span.End()
	var logger = gconv.String(ctx.Value("logger"))

	ok, err := dao.SysRole.IsUniqueName(ctx, in.Name)
	if err != nil {
		g.Log(logger).Error(ctx, "service Role EditIsUniqueName error:", err.Error())
		err = gerror.NewCode(gcode.New(500, "", nil), "操作失败[01]")
		return err
	}

	if ok {
		g.Log(logger).Error(ctx, "service Role name exist")
		err = gerror.NewCode(gcode.New(500, "", nil), "角色名称已存在")
		return err
	}

	tx, err := g.DB().Begin(ctx)
	if err != nil {
		g.Log(logger).Error(ctx, "service Role Del Transaction error:", err.Error())
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
			g.Log(logger).Error(ctx, "service Role Edit update mysql error:", err.Error())
			err = gerror.NewCode(gcode.New(500, "", nil), "操作失败[03]")
			return
		}
	} else {
		in.ID, err = dao.SysRole.Ctx(ctx).TX(tx).OmitEmpty().InsertAndGetId(in)
		if err != nil {
			g.Log(logger).Error(ctx, "service Role Edit insert mysql error:", err.Error())
			err = gerror.NewCode(gcode.New(500, "", nil), "操作失败[04]")
			return
		}
	}

	if err = dao.SysRoleMenu.UpdateRoleMenuByIds(ctx, in.ID, in.MenuIds, tx); err != nil {
		g.Log(logger).Error(ctx, "service Role Edit update UpdateRoleMenuByIds error:", err.Error())
		err = gerror.NewCode(gcode.New(500, "", nil), "操作失败[05]")
		return
	}

	return
}

// Del .删除角色
func (s *sRole) Del(ctx context.Context, in *model.RoleBaseInput) (err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-role-Del")
	defer span.End()
	var logger = gconv.String(ctx.Value("logger"))

	tx, err := g.DB().Begin(ctx)
	if err != nil {
		g.Log(logger).Error(ctx, "service Role Del Transaction error:", err.Error())
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
		g.Log(logger).Error(ctx, "service Role Del delete mysql error:", err.Error())
		err = gerror.NewCode(gcode.New(500, "", nil), "操作失败[02]")
		return
	}

	if _, err = dao.SysRoleMenu.Ctx(ctx).TX(tx).Delete("role_id", in.ID); err != nil {
		g.Log(logger).Error(ctx, "service RoleMenu Del delete mysql error:", err.Error())
		err = gerror.NewCode(gcode.New(500, "", nil), "操作失败[03]")
		return
	}

	return
}

// List 获取角色列表
func (s *sRole) List(ctx context.Context, in *model.RoleListInput) (*model.RoleListOut, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-role-list")
	defer span.End()

	var (
		logger = gconv.String(ctx.Value("logger"))
		out    = &model.RoleListOut{}
		list   []*model.RoleItem
	)

	m := dao.SysRole.Ctx(ctx)
	if in.ID > 0 {
		m = m.Where("id = ?", in.ID)
	}

	if in.Name != "" {
		m = m.Where("name = ?", in.Name)
	}

	if in.Status > 0 {
		m = m.Where("status = ?", in.Status)
	}

	count, err := m.Count()
	if err != nil {
		g.Log(logger).Error(ctx, "service Role list count error:", err.Error())
		err = gerror.NewCode(gcode.New(500, "", nil), "获取角色数据失败[01]")
		return nil, err
	}

	if count > 0 {
		if err = m.Page(in.Page, in.PageSize).Order("sort asc,id asc").Scan(&list); err != nil {
			g.Log(logger).Error(ctx, "service Role list scan error:", err.Error())
			err = gerror.NewCode(gcode.New(500, "", nil), "获取角色数据失败[02]")
			return nil, err
		}
	} else {
		list = []*model.RoleItem{}
	}

	out.List = list
	out.PageBaseInfo.Page = in.Page
	out.PageBaseInfo.PageSize = in.PageSize
	out.PageInfo.Total = count

	return out, nil
}

// GetAll 获取所有角色
func (s sRole) GetAll(ctx context.Context) (out []*model.RoleItem, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-role-list")
	defer span.End()
	var logger = gconv.String(ctx.Value("logger"))

	if err = dao.SysRole.Ctx(ctx).Where("status = 10").Order("sort asc,id asc").Scan(&out); err != nil {
		g.Log(logger).Error(ctx, "service GetAll list scan error:", err.Error())
		err = gerror.NewCode(gcode.New(500, "", nil), "获取角色数据失败")
		return
	}
	return
}
