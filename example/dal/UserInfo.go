package dal

import (
	"github.com/woaijssss/godbx"
	"github.com/woaijssss/godbx/ttypes"
)

var UserInfoFields = struct {
	Id       string
	Name     string
	CreateAt string
	ModifyAt string
}{
	"id",
	"name",
	"create_at",
	"modify_at",
}

var UserInfoMeta = &godbx.TableMeta[UserInfo]{
	Table: "user_info",
	Columns: []string{
		"id",
		"name",
		"create_at",
		"modify_at",
	},
	AutoColumn: "id",
	LookupFieldFunc: func(columnName string, ins *UserInfo, point bool) any {
		if "id" == columnName {
			if point {
				return &ins.Id
			}
			return ins.Id
		}
		if "name" == columnName {
			if point {
				return &ins.Name
			}
			return ins.Name
		}
		if "create_at" == columnName {
			if point {
				return &ins.CreateAt
			}
			return ins.CreateAt
		}
		if "modify_at" == columnName {
			if point {
				return &ins.ModifyAt
			}
			return ins.ModifyAt
		}

		return nil
	},
}

var UserInfoDao godbx.QuickDao[UserInfo] = &struct {
	godbx.QuickDao[UserInfo]
}{
	godbx.NewBaseQuickDao(UserInfoMeta),
}

type UserInfo struct {
	Id       int64
	Name     string
	CreateAt ttypes.NormalDatetime
	ModifyAt ttypes.NilableDatetime
}
