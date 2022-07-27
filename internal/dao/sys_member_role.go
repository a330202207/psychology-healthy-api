// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/a330202207/psychology-healthy-api/internal/dao/internal"
)

// internalSysMemberRoleDao is internal type for wrapping internal DAO implements.
type internalSysMemberRoleDao = *internal.SysMemberRoleDao

// sysMemberRoleDao is the data access object for table sys_member_role.
// You can define custom methods on it to extend its functionality as you wish.
type sysMemberRoleDao struct {
	internalSysMemberRoleDao
}

var (
	// SysMemberRole is globally public accessible object for table sys_member_role operations.
	SysMemberRole = sysMemberRoleDao{
		internal.NewSysMemberRoleDao(),
	}
)

// Fill with you ideas below.

// UpdateMemberRoleByIds 更新用户角色
func (d *sysMemberRoleDao) UpdateMemberRoleByIds(ctx context.Context, memberId int64, ruleIds []int64, tx *gdb.TX) (err error) {
	if _, err = d.Ctx(ctx).TX(tx).Where("member_id", memberId).Delete(); err != nil {
		return
	}

	var ruleData []map[string]interface{}

	for _, val := range ruleIds {
		var data = make(map[string]interface{})
		data["member_id"] = memberId
		data["rule_id"] = val

		ruleData = append(ruleData, data)
	}

	if _, err = d.Ctx(ctx).TX(tx).Insert(ruleData); err != nil {
		return
	}

	return
}
