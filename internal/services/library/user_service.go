package library

import (
	"context"

	"github.com/iButcat/quickgogo/internal/models"
)

type UserService interface {
	Create(ctx context.Context, user *models.User) error
}
