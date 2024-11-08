package library

import (
	"context"

	"github.com/iButcat/quickgogo/internal/models"
	"golang.org/x/oauth2"
)

type AuthService interface {
	GetAuthURL(ctx context.Context, state string) (string, error)
	ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error)
	GetUserInfo(ctx context.Context, token *oauth2.Token) (*models.User, error)
}
