package shop

import (
	"context"
	"errors"

	repoShop "github.com/oganes5796/shops-list/internal/shops-list/repository/shop"
)

func (s *shopService) Delete(ctx context.Context, idShop int64) error {
	err := s.repo.Delete(ctx, idShop)
	if err != nil {
		if errors.Is(err, repoShop.ErrShopNotFound) {
			return ErrShopNotFound
		}
		return err
	}

	return nil
}
