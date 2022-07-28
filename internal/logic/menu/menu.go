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
		err = errors.New("操作失败[001]")
		return err
	}

	if ok {
		g.Log(logger).Error(ctx, "service Menu name exist")
		err = errors.New("菜单名称已存在")
		return err
	}

	tx, err := g.DB().Begin(ctx)
	if err != nil {
		g.Log(logger).Error(ctx, "service Menu Del Transaction error:", err.Error())
		err = errors.New("操作失败[002]")
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

	if _, err = dao.SysMenu.Ctx(ctx).TX(tx).Where("id", in.ID).Delete(); err != nil {
		g.Log(logger).Error(ctx, "service Menu Del error:", err.Error())
		err = errors.New("操作失败[002]")
		return err
	}

	if _, err = dao.SysRoleMenu.Ctx(ctx).TX(tx).Where("menu_id", in.ID).Delete(); err != nil {
		g.Log(logger).Error(ctx, "service RuleMenu Del error:", err.Error())
		err = errors.New("操作失败[003]")
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
func (s *sMenu) GetAll(ctx context.Context) (res []*model.MenuTree, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-menu-GetAll")
	defer span.End()

	var (
		logger = gconv.String(ctx.Value("logger"))
		list   []*entity.SysMenu
	)

	if err = dao.SysMenu.Ctx(ctx).Where("status = ?", 10).Where("is_visible = ?", 20).Scan(&list); err != nil {
		g.Log(logger).Error(ctx, "service Menu GetAll error:", err.Error())
		return
	}

	res = dao.SysMenu.GenTreeList(ctx, 0, list)

	return
}
