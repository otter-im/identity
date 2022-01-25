package app

import (
	"context"
	"fmt"
	"github.com/go-redis/cache/v8"
	"github.com/google/uuid"
	"time"
)

type User struct {
	tableName struct{} `pg:",alias:u"`

	Id           uuid.UUID `pg:"id"`
	Username     string    `json:"username" pg:"username"`
	Hash         []byte    `json:"-" pg:"hash"`
	Salt         []byte    `json:"-" pg:"salt"`
	CreationDate time.Time `json:"creation_date" pg:"creation_date"`
}

type UserService interface {
	SelectUser(ctx context.Context, id uuid.UUID) (*User, error)
	SelectUserByUsername(ctx context.Context, username string) (*User, error)
}

type userDto struct{}

func (u *userDto) SelectUser(ctx context.Context, id uuid.UUID) (*User, error) {
	user := new(User)
	if err := RedisCache().Once(&cache.Item{
		Ctx:   ctx,
		Key:   fmt.Sprintf("user:%s", id.String()),
		Value: user,
		TTL:   15 * time.Minute,
		Do: func(item *cache.Item) (interface{}, error) {
			return selectUser(ctx, id)
		},
	}); err != nil {
		return nil, err
	}
	return user, nil
}

func selectUser(ctx context.Context, id uuid.UUID) (*User, error) {
	user := new(User)
	if err := Postgres().
		ModelContext(ctx, user).
		Where("id = ?", id).
		Select(); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userDto) SelectUserByUsername(ctx context.Context, username string) (*User, error) {
	user := new(User)
	if err := Postgres().
		ModelContext(ctx, user).
		Where("username = ?", username).
		Select(); err != nil {
		return nil, err
	}
	return user, nil
}
