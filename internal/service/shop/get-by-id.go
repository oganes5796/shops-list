package shop

import (
	"context"

	"github.com/oganes5796/shops-list/internal/model"
)

func (r *shopService) GetByID(ctx context.Context, id int64) (*model.Shop, error) {
	return r.repo.GetByID(ctx, id)
}
