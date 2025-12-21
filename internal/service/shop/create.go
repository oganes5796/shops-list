package shop

import (
	"context"

	"github.com/oganes5796/shops-list/internal/model"
)

func (r *shopService) Create(ctx context.Context, info *model.ShopInfo) (int64, error) {
	return r.repo.Create(ctx, info)
}
