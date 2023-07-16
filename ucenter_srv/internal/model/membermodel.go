package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ MemberModel = (*customMemberModel)(nil)

type (
	// MemberModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMemberModel.
	MemberModel interface {
		memberModel
	}

	customMemberModel struct {
		*defaultMemberModel
	}
)

const SendCodeCacheKey = "MEMBER::"

// NewMemberModel returns a model for the database table.
func NewMemberModel(conn sqlx.SqlConn) MemberModel {
	return &customMemberModel{
		defaultMemberModel: newMemberModel(conn),
	}
}
