package user

import (
	"b2b/entities"
	"context"
)

type Repo interface {
	Create(ctx context.Context, member entities.User) (*entities.User, error)
	Get(ctx context.Context, filter entities.User) (*entities.User, error)
	Update(ctx context.Context, userId string, updateModel entities.User) error
}
