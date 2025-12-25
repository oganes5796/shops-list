package shop

import (
	"context"
	"fmt"

	"github.com/oganes5796/shops-list/internal/model"
)

func (r *shopRepository) Update(ctx context.Context, idShop int64, info *model.ShopInfo) error {
	query := fmt.Sprintf(`
		UPDATE %s 
        SET title = $1, address = $2, operating_mode = $3, updated_at = NOW() 
        WHERE id = $4`,
		tableName,
	)
	result, err := r.conn.Exec(ctx, query,
		info.Title, info.Address, info.OperatingMode, idShop)
	if err != nil {
		return fmt.Errorf("repository:Update:exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return ErrShopNotFound
	}

	return nil
}
