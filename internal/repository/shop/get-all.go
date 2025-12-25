package shop

import (
	"context"
	"fmt"

	"github.com/oganes5796/shops-list/internal/model"
)

func (r *shopRepository) GetAll(ctx context.Context) ([]model.Shop, error) {
	var lists []model.Shop

	query := fmt.Sprintf("SELECT id, title, address, operating_mode, created_at, updated_at FROM %s", tableName)
	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository:GetAll:Query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var shop model.Shop
		if err := rows.Scan(
			&shop.ID,
			&shop.Info.Title,
			&shop.Info.Address,
			&shop.Info.OperatingMode,
			&shop.CreatedAt,
			&shop.UpdatedAt); err != nil {
			return nil, fmt.Errorf("repository:GetAll:Scan: %w", err)
		}
		lists = append(lists, shop)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repository:GetAll:RowsErr: %w", err)
	}

	return lists, nil
}
