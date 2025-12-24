package shop

import (
	"context"

	"errors"

	"github.com/oganes5796/shops-list/internal/model"
)

func (r *shopService) Update(ctx context.Context, idShop int64, info *model.ShopInfo) error {
	err := r.repo.Update(ctx, idShop, info)
	if err != nil {
		if errors.Is(err, model.ErrShopNotFound) {
			return model.ErrShopNotFound
		}
		return err
	}

	return nil
}
