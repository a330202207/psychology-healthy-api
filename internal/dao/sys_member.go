// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/a330202207/psychology-healthy-api/internal/consts"
	"github.com/a330202207/psychology-healthy-api/internal/dao/internal"
	"github.com/a330202207/psychology-healthy-api/internal/model/entity"
)

// internalSysMemberDao is internal type for wrapping internal DAO implements.
type internalSysMemberDao = *internal.SysMemberDao

// sysMemberDao is the data access object for table sys_member.
// You can define custom methods on it to extend its functionality as you wish.
type sysMemberDao struct {
	internalSysMemberDao
}

var (
	// SysMember is globally public accessible object for table sys_member operations.
	SysMember = sysMemberDao{
		internal.NewSysMemberDao(),
	}
)

// IsUniqueName . 用户名是否唯一
func (d *sysMemberDao) IsUniqueName(ctx context.Context, id int64, name string) (bool, error) {
	var data *entity.SysMember
	m := d.Ctx(ctx).Where("username", name)

	if id > 0 {
		m = m.WhereNot("id", id)
	}

	if err := m.Scan(&data); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return false, err
	}

	if data == nil {
		return true, nil
	}

	return false, nil
}

// IsUniqueMobile 电话是否唯一
func (d *sysMemberDao) IsUniqueMobile(ctx context.Context, id int64, mobile string) (bool, error) {
	var data *entity.SysMember
	m := d.Ctx(ctx).Where("mobile", mobile)

	if id > 0 {
		m = m.WhereNot("id", id)
	}

	if err := m.Scan(&data); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return false, err
	}

	if data == nil {
		return true, nil
	}

	return false, nil
}

// IsUniqueEmail 邮箱是否唯一
func (d *sysMemberDao) IsUniqueEmail(ctx context.Context, id int64, email string) (bool, error) {
	var data *entity.SysMember
	m := d.Ctx(ctx).Where("email", email)

	if id > 0 {
		m = m.WhereNot("id", id)
	}

	if err := m.Scan(&data); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return false, err
	}

	if data == nil {
		return true, nil
	}

	return false, nil
}
