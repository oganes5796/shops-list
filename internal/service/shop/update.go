package shop

import (
	"context"

	"github.com/oganes5796/shops-list/internal/model"
)

func (r *shopService) Update(ctx context.Context, idShop int64, info *model.ShopInfo) error {
	return r.repo.Update(ctx, idShop, info)
}
