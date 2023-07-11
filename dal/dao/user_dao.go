package dao

import (
	"WebService/dal/do"
	"context"
)

type UserDao interface {
	GetUserByID(ctx context.Context, userID string) (*do.User, error)
	SaveUser(ctx context.Context, user *do.User) error
}
