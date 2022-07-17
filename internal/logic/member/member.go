// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/11 17:28
// @Package admin

package member

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/a330202207/psychology-healthy-api/internal/consts"
	"github.com/a330202207/psychology-healthy-api/internal/dao"
	"github.com/a330202207/psychology-healthy-api/internal/model"
	"github.com/a330202207/psychology-healthy-api/internal/service"
)

type sMember struct {
}

func init() {
	service.RegisterMember(New())
}

func New() *sMember {
	return &sMember{}
}

// Edit 新增/修改
func (s *sMember) Edit(ctx context.Context, in *model.MemberInput) (err error) {
	ok, err := dao.SysMember.IsUniqueName(ctx, in.ID, in.UserName)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if !ok {
		err = gerror.New("帐号已存在")
		return err
	}

	passWd, err := service.Auth().EncryptPass(in.Passwd)
	if err != nil {
		return err
	}

	in.Passwd = string(passWd)

	tx, err := g.DB().Begin(ctx)
	if err != nil {
		err = errors.New("操作失败")
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
		if _, err = dao.SysMember.Ctx(ctx).Where("id", in.ID).OmitEmpty().Update(in); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}
	} else {
		in.Passwd = string(passWd)
		in.ID, err = dao.SysMember.Ctx(ctx).OmitEmpty().InsertAndGetId(in)
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}
	}

	// 更新角色
	if err = service.Rule().UpdateRuleByIds(ctx, in.ID, in.RuleIds); err != nil {
		return
	}

	return err
}
