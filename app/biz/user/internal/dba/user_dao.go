package dba

import (
	"context"
	"fmt"
	"github.com/go-faster/errors"
	"gorm.io/gorm"
	"lend/app/biz/user/dm"
	"lend/gen/db"
	"lend/gen/model"
)

// Repository 抽象查询接口
type UserRepository struct {
	db *gorm.DB
	dq *db.Query
}

func NewUserRepository(db *gorm.DB, query *db.Query) *UserRepository {
	return &UserRepository{db: db, dq: query}
}

func (repo UserRepository) FindByName(ctx context.Context, name string) ([]*model.UserInfo, error) {
	return repo.dq.UserInfo.WithContext(ctx).Where(repo.dq.UserInfo.Name.Eq(name)).Find()

}

func (repo UserRepository) FindById(ctx context.Context, id int32) (*model.UserInfo, error) {
	return repo.dq.UserInfo.WithContext(ctx).Where(repo.dq.UserInfo.ID.Eq(id)).FirstOrCreate()
}

func (repo UserRepository) DisableUser(ctx context.Context, id int32) error {
	update, err := repo.dq.UserInfo.WithContext(ctx).Where(repo.dq.UserInfo.ID.Eq(id)).Update(repo.dq.UserInfo.IsDisable, true)
	if err != nil {
		return err
	}
	if update.RowsAffected <= 0 || update.Error != nil {
		return errors.New(fmt.Sprintf("update rows %d error detail: %s", update.RowsAffected, update.Error))
	}
	return nil
}

func (repo UserRepository) FindByForeignerId(ctx context.Context, foreignCode string, foreignerId string) (*dm.UserInfoF, error) {
	uF := &dm.UserInfoF{}
	UQ := repo.dq.UserInfo
	UFQ := repo.dq.UserForeignerID

	e := UQ.WithContext(ctx).
		Select(UQ.ALL, UFQ.ALL).
		LeftJoin(UFQ, UFQ.UserID.EqCol(UQ.ID)).
		Where(UFQ.ForeignerID.Eq(foreignerId)).
		Where(UFQ.ForeignCode.Eq(foreignCode)).
		Scan(uF)

	return uF, e
}
