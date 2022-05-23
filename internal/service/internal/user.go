package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type UserDao struct {
	table string
	group string
	columns UserColumns
}

type UserColumns struct {
	Id string
	Name string
	Age string
	CreatedAt string
	UpdatedAt string
}

func NewUserDao() *UserDao {
	return &UserDao{
		group: "default",
		table: "user",
		columns: UserColumns{
			Id: "id",
			Name: "name",
			Age: "age",
			CreatedAt: "created_at",
			UpdatedAt: "updated_at",
		},
	}
}

func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

func (dao *UserDao) Table() string {
	return dao.table
}

func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

func (dao *UserDao) Group() string {
	return dao.group
}

func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

func (dao *UserDao) Tx(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) error {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
