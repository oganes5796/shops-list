package shop

import (
	"context"
	"fmt"
)

func (r *shopRepository) Delete(ctx context.Context, idShop int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableName)
	result, err := r.conn.Exec(ctx, query, idShop)

	if err != nil {
		return fmt.Errorf("repository:shop:Delete:exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return ErrShopNotFound
	}

	return nil
}
