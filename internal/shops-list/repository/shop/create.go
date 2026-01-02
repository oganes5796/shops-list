package shop

import (
	"context"
	"fmt"

	"github.com/oganes5796/shops-list/internal/shops-list/model"
)

func (r *shopRepository) Create(ctx context.Context, info *model.ShopInfo) (int64, error) {
	var id int64
	query := fmt.Sprintf("INSERT INTO %s (title, address, operating_mode) VALUES ($1, $2, $3) RETURNING id", tableName)
	row := r.conn.QueryRow(ctx, query, info.Title, info.Address, info.OperatingMode)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("repository:shop:Create:row.Scan")
	}
	return id, nil
}
