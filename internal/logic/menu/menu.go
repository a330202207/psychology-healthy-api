// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/14 16:02
// @Package menu

package menu

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/a330202207/psychology-healthy-api/internal/dao"
	"github.com/a330202207/psychology-healthy-api/internal/model"
	"github.com/a330202207/psychology-healthy-api/internal/model/entity"
	"github.com/a330202207/psychology-healthy-api/internal/service"
)

type sMenu struct {
}

func init() {
	service.RegisterMenu(New())
}

func New() *sMenu {
	return &sMenu{}
}

// Edit 添加/编辑菜单
func (s *sMenu) Edit(ctx context.Context, in *model.MenuEditInput) (err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-menu-Edit")
	defer span.End()
	var logger = gconv.String(ctx.Value("logger"))

	ok, err := dao.SysMenu.IsUniqueName(ctx, in.Name)
	if err != nil {
		g.Log(logger).Error(ctx, "service Menu EditIsUniqueName error:", err.Error())
		err = gerror.NewCode(gcode.New(500, "", nil), "操作失败(01)")
		return err
	}

	if ok {
		g.Log(logger).Error(ctx, "service Menu name exist")
		err = gerror.NewCode(gcode.New(500, "", nil), "菜单名称已存在")
		return err
	}

	tx, err := g.DB().Begin(ctx)
	if err != nil {
		g.Log(logger).Error(ctx, "service Menu Del Transaction error:", err.Error())
		err = gerror.NewCode(gcode.New(500, "", nil), "操作失败[02]")
		return err
	}

	if in.ID > 0 {
		if _, err = dao.SysRole.Ctx(ctx).TX(tx).Where("id", in.ID).OmitEmpty().Update(in); err != nil {
			g.Log(logger).Error(ctx, "service Menu Edit update mysql error:", err.Error())
			err = errors.New("操作失败[002]")
			return
		}
	} else {
		in.ID, err = dao.SysMenu.Ctx(ctx).TX(tx).OmitEmpty().InsertAndGetId()
		if err != nil {
			g.Log(logger).Error(ctx, "service Menu Edit insert mysql error[01]:", err.Error())
			err = errors.New("操作失败[003]")
			return
		}

		if in.ID != 1 {
			g.Log(logger).Error(ctx, "service Menu Edit insert mysql error[02],ID:", in.ID)
			err = errors.New("操作失败[004]")
			return
		}
	}

	return
}

// Del 删除菜单
func (s *sMenu) Del(ctx context.Context, in *model.MenuBaseInput) (err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-menu-Del")
	defer span.End()
	var logger = gconv.String(ctx.Value("logger"))

	tx, err := g.DB().Begin(ctx)
	if err != nil {
		g.Log(logger).Error(ctx, "service Menu Del Transaction error:", err.Error())
		err = gerror.NewCode(gcode.New(500, "", nil), "操作失败[01]")
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	if _, err = dao.SysMenu.Ctx(ctx).TX(tx).Where("id", in.ID).Delete(); err != nil {
		g.Log(logger).Error(ctx, "service Menu Del error:", err.Error())
		err = gerror.NewCode(gcode.New(500, "", nil), "操作失败[02]")
		return err
	}

	if _, err = dao.SysRoleMenu.Ctx(ctx).TX(tx).Where("menu_id", in.ID).Delete(); err != nil {
		g.Log(logger).Error(ctx, "service RoleMenu Del error:", err.Error())
		err = gerror.NewCode(gcode.New(500, "", nil), "操作失败[03]")
		return err
	}

	return
}

// Get 菜单详情
func (s *sMenu) Get(ctx context.Context, in *model.MenuBaseInput) (res []*model.MenuGetOut, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-menu-Get")
	defer span.End()
	return
}

// GetAll 菜单列表
func (s *sMenu) GetAll(ctx context.Context) ([]*model.MenuTree, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-menu-GetAll")
	defer span.End()

	var (
		logger = gconv.String(ctx.Value("logger"))
		out    = make([]*model.MenuTree, 0)
		list   []*entity.SysMenu
	)

	list, err := dao.SysMenu.GetList(ctx)
	if err != nil {
		g.Log(logger).Error(ctx, "service Menu GetAll error:", err.Error())
		err = gerror.NewCode(gcode.New(500, "", nil), "获取列表失败[01]")
		return nil, err
	}

	if len(list) > 0 {
		out, err = dao.SysMenu.GenTreeList(0, list)
		if err != nil {
			g.Log(logger).Error(ctx, "service Menu GenTreeList error:", err.Error())
			err = gerror.NewCode(gcode.New(500, "", nil), "获取列表失败[02]")
			return nil, err
		}
	}

	return out, nil
}

// List 菜单列表
func (s *sMenu) List(ctx context.Context, in *model.MenuListInput) (*model.MenuListOut, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-menu-List")
	defer span.End()
	var (
		logger = gconv.String(ctx.Value("logger"))
		out    = &model.MenuListOut{}
		list   []*model.MenuItem
	)

	m := dao.SysMenu.Ctx(ctx)

	if len(in.Name) > 0 {
		m = m.Where("name like ?", "%"+in.Name+"%")
	}

	if in.Status > 0 {
		m = m.Where("status = ?", in.Status)
	}

	if in.Type > 0 {
		m = m.Where("type = ?", in.Type)
	}

	count, err := m.Count()
	if err != nil {
		g.Log(logger).Error(ctx, "service Menu list count error:", err.Error())
		err = gerror.NewCode(gcode.New(500, "", nil), "获取菜单数据失败[01]")
		return nil, err

	}

	if count > 0 {
		if err = m.Page(in.Page, in.PageSize).Order("sort asc,id asc").Scan(&list); err != nil {
			g.Log(logger).Error(ctx, "service Menu list scan error:", err.Error())
			err = gerror.NewCode(gcode.New(500, "", nil), "获取菜单数据失败[02]")
			return nil, err
		}
	} else {
		list = []*model.MenuItem{}
	}

	out.List = list
	out.PageBaseInfo.Page = in.Page
	out.PageBaseInfo.PageSize = in.PageSize
	out.PageInfo.Total = count

	return out, nil
}
