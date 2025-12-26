package shop

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/oganes5796/shops-list/internal/model"
)

func (r *shopRepository) GetByID(ctx context.Context, idShop int64) (*model.Shop, error) {
	var shop model.Shop

	query := fmt.Sprintf("SELECT id, title, address, operating_mode, created_at, updated_at FROM %s WHERE id=$1", tableName)
	row := r.conn.QueryRow(ctx, query, idShop)
	err := row.Scan(
		&shop.ID,
		&shop.Info.Title,
		&shop.Info.Address,
		&shop.Info.OperatingMode,
		&shop.CreatedAt,
		&shop.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrShopNotFound
		}

		return nil, fmt.Errorf("repository:shop:GetByID:row.Scan: %w", err)
	}

	return &shop, nil
}
