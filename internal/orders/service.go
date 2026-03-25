package orders

import (
	"context"

	repo "github.com/Shenith404/go-ecom/internal/adapters/postgre/sqlc"
)

type svc struct {
	repo *repo.Queries
}

func NewService(r *repo.Queries) Service {
	return &svc{
		repo: r,
	}
}

func (s *svc) PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order ,error) {}