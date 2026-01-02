package shop

import (
	"context"

	"github.com/oganes5796/shops-list/internal/shops-list/model"
)

func (s *shopService) GetAll(ctx context.Context) ([]model.Shop, error) {
	shops, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return shops, nil
}
