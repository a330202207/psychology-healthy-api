// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"

	"github.com/a330202207/psychology-healthy-api/internal/model"
)

type IRule interface {
	Edit(ctx context.Context, in *model.RuleEditInput) (err error)
	Del(ctx context.Context, in *model.RuleBaseInput) (err error)
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
