package dao_impl

import (
	"WebService/dal/do"
	"WebService/dal/driver/rdb"
	"context"
	"encoding/json"
)

type UserDaoImpl struct {
}

func (d *UserDaoImpl) GetUserByID(ctx context.Context, userID string) (*do.User, error) {
	v, err := rdb.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	if v == "" {
		return nil, nil
	}
	var user do.User
	err = json.Unmarshal([]byte(v), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (d *UserDaoImpl) SaveUser(ctx context.Context, user *do.User) error {
	if user == nil {
		return nil
	}
	jsonStr, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return rdb.Set(ctx, user.ID, string(jsonStr), 0)
}
