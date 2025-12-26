package shop

import (
	"context"

	"errors"

	"github.com/oganes5796/shops-list/internal/model"
	repoShop "github.com/oganes5796/shops-list/internal/repository/shop"
)

func (s *shopService) Update(ctx context.Context, idShop int64, info *model.ShopInfo) error {
	err := s.repo.Update(ctx, idShop, info)
	if err != nil {
		if errors.Is(err, repoShop.ErrShopNotFound) {
			return ErrShopNotFound
		}
		return err
	}

	return nil
}
