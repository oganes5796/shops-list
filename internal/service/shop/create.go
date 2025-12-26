package shop

import (
	"context"

	"github.com/oganes5796/shops-list/internal/model"
)

func (s *shopService) Create(ctx context.Context, info *model.ShopInfo) (int64, error) {
	return s.repo.Create(ctx, info)
}
