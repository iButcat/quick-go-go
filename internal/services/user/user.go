package user

import (
	"context"

	"github.com/iButcat/quickgogo/internal/models"
	"github.com/iButcat/quickgogo/internal/services/library"
)

type userService struct {
}

func NewUserService() library.UserService {
	return &userService{}
}

func (u *userService) Create(ctx context.Context, user *models.User) error {
	return nil
}
