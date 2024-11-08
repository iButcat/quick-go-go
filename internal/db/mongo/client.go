package mongo

import (
	"github.com/iButcat/quickgogo/internal/db"
	"github.com/iButcat/quickgogo/internal/db/mongo/repositories"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func New(client *mongo.Client, logger *zap.Logger) *db.Store {
	return &db.Store{
		UserRepository: repositories.NewUserRepository(client, logger),
	}
}
