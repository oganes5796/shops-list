package shop

import (
	"context"

	"github.com/oganes5796/shops-list/internal/model"
)

func (r *shopService) GetAll(ctx context.Context) ([]model.Shop, error) {
	return r.repo.GetAll(ctx)
}
