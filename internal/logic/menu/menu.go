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
func (s *sMenu) Get(ctx context.Context, in *model.MenuBaseInput) (res *model.MenuGetOut, err error) {

	return

}
