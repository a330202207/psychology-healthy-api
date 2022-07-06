// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"psychology-healthy-api/internal/dao/internal"
)

// internalAdminDao is internal type for wrapping internal DAO implements.
type internalAdminDao = *internal.AdminDao

// adminDao is the data access object for table admin.
// You can define custom methods on it to extend its functionality as you wish.
type adminDao struct {
	internalAdminDao
}

var (
	// Admin is globally public accessible object for table admin operations.
	Admin = adminDao{
		internal.NewAdminDao(),
	}
)

// Fill with you ideas below.
