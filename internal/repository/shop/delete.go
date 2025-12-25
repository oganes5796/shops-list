package shop

import (
	"context"
	"fmt"
)

const tableName = "shops"

func (r *shopRepository) Delete(ctx context.Context, idShop int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableName)
	_, err := r.conn.Exec(ctx, query, idShop)
	if err != nil {
		return fmt.Errorf("repository:Delete:exec: %w", err)
	}
	return nil
}
