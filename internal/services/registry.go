package registry

import (
	"github.com/iButcat/quickgogo/internal/services/library"
	"go.uber.org/zap"
)

// Registry contains everything we need to share accross APIs
type Registry struct {
	AuthService library.AuthService
	Logger      *zap.Logger
}

func NewRegistry(logger *zap.Logger) library.Library {
	return &Registry{}
}

func (r *Registry) GetAuthService() library.AuthService {
	return r.AuthService
}
