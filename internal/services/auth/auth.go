package auth

import (
	"context"

	"github.com/iButcat/quickgogo/internal/models"
	"github.com/iButcat/quickgogo/internal/services/library"
	"golang.org/x/oauth2"
)

type authService struct {
	authConfig *oauth2.Config
}

func NewAuthService(auth *oauth2.Config) library.AuthService {
	return &authService{
		authConfig: auth,
	}
}

func (a *authService) GetAuthURL(ctx context.Context, state string) (string, error) {
	return a.authConfig.AuthCodeURL(state), nil
}

func (a *authService) ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	return nil, nil
}

func (a *authService) GetUserInfo(ctx context.Context, token *oauth2.Token) (*models.User, error) {
	return nil, nil
}
