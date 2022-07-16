// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
)

type IRule interface {
	UpdateRuleByIds(ctx context.Context, memberId int64, ruleIds []int64) (err error)
}

var localRule IRule

func Rule() IRule {
	if localRule == nil {
		panic("implement not found for interface IRule, forgot register?")
	}
	return localRule
}

func RegisterRule(i IRule) {
	localRule = i
}
