package shop

import "context"

func (r *shopService) Delete(ctx context.Context, idShop int64) error {
	return r.repo.Delete(ctx, idShop)
}
