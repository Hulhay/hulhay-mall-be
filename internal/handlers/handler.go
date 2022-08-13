package handlers

import (
	"context"
	configs "hulhay-mall/config"
	"hulhay-mall/internal/apis/operations/user"
	"hulhay-mall/internal/models"
	"hulhay-mall/internal/usecase"
)

type handler struct {
	useCase usecase.UseCase
}

type Handlers interface {
	GetHealtcheck() (*models.Health, error)

	CreateStore(ctx context.Context, identifier string, params *models.StoresRequest) error
	GetStores(ctx context.Context, identifier string) ([]*models.Stores, error)
	GetStoreByID(ctx context.Context, identifier string, storeID string) (*models.Stores, error)
	DeleteStoreByID(ctx context.Context, identifier string, storeID string) error
	UpdateStoreByID(ctx context.Context, identifier string, params *models.StoresRequest, storeID string) error

	Register(ctx context.Context, params *models.RegisterRequest) error
	Login(ctx context.Context, params *models.LoginRequest) (*models.LoginResponse, error)
	Logout(ctx context.Context, params *user.PatchLogoutParams) error
}

func NewHandler() Handlers {
	return &handler{
		useCase: configs.GetUsecase(),
	}
}
