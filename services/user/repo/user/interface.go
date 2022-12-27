package user

import (
	"context"
	"user/entities"
)

type Repo interface {
	Create(ctx context.Context, member entities.User) (*entities.User, error)
	Get(ctx context.Context, filter entities.User) (*entities.User, error)
}
