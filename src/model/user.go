package model

import (
	"context"
	"encoding/json"
	mw "thegame/middleware"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	GameState  *GameState `json:"gameState"`
	Friends    string     `json:"friends" gorm:"type:text"`
	FriendsArr []string   `gorm:"-"`
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	friends, err := json.Marshal(u.FriendsArr)
	if err != nil {
		return err
	}
	u.Friends = string(friends)
	return nil
}

func (u *User) AfterFind(tx *gorm.DB) error {
	err := json.Unmarshal([]byte(u.Friends), &u.FriendsArr)
	return err
}

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

func (u *User) GetGameState(ctx context.Context, user *User) (*GameState, error) {
	state := &GameState{UserID: user.ID}
	db := mw.GetDbFromContext(ctx)
	if err := db.First(state).Error; err != nil {
		return nil, err
	}
	return state, nil
}

func (u *User) AddFriends(ctx context.Context, friends []string) ([]*Friend, error) {
	db := mw.GetDbFromContext(ctx)
	u.FriendsArr = friends

	if err := db.Save(u).Error; err != nil {
		return nil, err
	}
	fs, err := u.GetFriends(ctx)
	if err != nil {
		return nil, err
	}
	return fs, nil
}

func (u *User) GetFriends(ctx context.Context) ([]*Friend, error) {
	db := mw.GetDbFromContext(ctx)
	user := &User{ID: u.ID}
	if err := db.First(user).Error; err != nil {
		return nil, err
	}
	var users []*User
	if err := db.Where("id in (?)", user.Friends).Find(&users).Error; err != nil {
		return nil, err
	}
	return usersToFriends(users), nil
}

func usersToFriends(users []*User) []*Friend {
	friends := make([]*Friend, 0)

	for _, user := range users {
		friend := &Friend{
			ID:        user.ID,
			Name:      user.Name,
			Highscore: *user.GameState.Score,
		}
		friends = append(friends, friend)
	}
	return friends
}
