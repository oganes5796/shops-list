package shop

import (
	"context"
	"errors"

	"github.com/oganes5796/shops-list/internal/shops-list/model"
	repoShop "github.com/oganes5796/shops-list/internal/shops-list/repository/shop"
)

func (s *shopService) GetByID(ctx context.Context, idShop int64) (*model.Shop, error) {
	shop, err := s.repo.GetByID(ctx, idShop)
	if err != nil {
		if errors.Is(err, repoShop.ErrShopNotFound) {
			return nil, ErrShopNotFound
		}
		return nil, err
	}

	return shop, nil
}
