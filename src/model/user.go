package model

import (
	"context"
	mw "thegame/middleware"
)

func (u *User) Create(ctx context.Context, user *User) (*User, error) {
	db := mw.GetDbFromContext(ctx)
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) GetAll(ctx context.Context) ([]*User, error) {
	var users []*User
	db := mw.GetDbFromContext(ctx)
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
