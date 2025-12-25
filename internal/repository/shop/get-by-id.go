package shop

import (
	"context"
	"fmt"

	"github.com/oganes5796/shops-list/internal/model"
)

const (
	tableName = "shops"
)

func (r *shopRepository) GetByID(ctx context.Context, id int64) (*model.Shop, error) {
	var shop model.Shop

	query := fmt.Sprintf("SELECT id, title, address, operating_mode, created_at, updated_at FROM %s WHERE id=$1", tableName)
	row := r.conn.QueryRow(ctx, query, id)
	err := row.Scan(
		&shop.ID,
		&shop.Info.Title,
		&shop.Info.Address,
		&shop.Info.OperatingMode,
		&shop.CreatedAt,
		&shop.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("repository:GetByID:row.Scan: %w", err)
	}

	return &shop, nil
}
