package db

import (
	"context"
	"github.com/HelliWrold1/quaver/dal/model"
	"github.com/HelliWrold1/quaver/pkg/errno"
)

func CreateUser(ctx context.Context, u *model.User) (int64, error) {
	if existUser(ctx, u) {
		return 0, errno.UserAlreadyExistErr
	}
	q := Q.User.WithContext(ctx)
	err := q.Create(u)
	return u.ID, err
}

func QueryUserByName(ctx context.Context, username string) ([]*model.User, error) {
	q := Q.User.WithContext(ctx)
	return q.Where(Q.User.Username.Eq(username)).Find()
}

func QueryUserByID(ctx context.Context, id int64) ([]*model.User, error) {
	q := Q.User.WithContext(ctx)
	return q.Where(Q.User.ID.Eq(id)).Find()
}

func existUser(ctx context.Context, u *model.User) bool {
	q := Q.User.WithContext(ctx)
	count, _ := q.Where(Q.User.Username.Eq(u.Username)).Count()
	if count != 0 {
		return true
	}
	return false
}
