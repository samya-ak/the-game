package model

import (
	"context"
	mw "thegame/middleware"

	"github.com/google/uuid"
)

func GetIntPointer(value int) *int {
	return &value
}

func (u *User) Create(ctx context.Context, user *User) (*User, error) {
	db := mw.GetDbFromContext(ctx)

	tx := db.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	gameState := &GameState{
		ID:          uuid.NewString(),
		GamesPlayed: GetIntPointer(0),
		Score:       GetIntPointer(0),
		UserID:      user.ID,
	}

	if err := tx.Create(&gameState).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) GetAll(ctx context.Context) ([]*User, error) {
	var users []*User
	db := mw.GetDbFromContext(ctx)
	if err := db.Preload("GameState").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u User) UpdateGameState(ctx context.Context, input *UserGameState) (*GameState, error) {
	state := &GameState{UserID: input.UserID}
	db := mw.GetDbFromContext(ctx)

	if err := db.First(state).Error; err != nil {
		return nil, err
	}
	state.GamesPlayed = input.GamesPlayed
	state.Score = input.Score
	if err := db.Save(state).Error; err != nil {
		return nil, err
	}
	return state, nil
}
