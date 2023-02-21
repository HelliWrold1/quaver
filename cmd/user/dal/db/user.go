package db

import (
	"context"
	"github.com/HelliWrold1/quaver/dal/model"
)

func CreateUser(ctx context.Context, u *model.User) error {
	existUser(ctx, u)
	q := Q.User.WithContext(ctx)
	return q.Create(u)
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
