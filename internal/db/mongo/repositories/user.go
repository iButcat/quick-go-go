package repositories

import (
	"context"

	"github.com/iButcat/quickgogo/internal/db"
	"github.com/iButcat/quickgogo/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type userRepository struct {
	client *mongo.Client
	logger *zap.Logger
}

func NewUserRepository(client *mongo.Client, logger *zap.Logger) db.UserRepository {
	return &userRepository{
		client: client,
		logger: logger,
	}
}

func (u *userRepository) Create(ctx context.Context, user *models.User) error {
	return nil
}
