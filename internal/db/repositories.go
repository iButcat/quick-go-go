package db

import (
	"context"

	"github.com/iButcat/quickgogo/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
}
