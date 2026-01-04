package orders

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/oganes5796/shops-list/internal/orders/model"
	"github.com/oganes5796/shops-list/internal/orders/repository"
)

const (
	tableName = "orders"
)

type ordersRepository struct {
	conn *pgx.Conn
}

func NewOrdersRepository(conn *pgx.Conn) repository.OrderRepository {
	return &ordersRepository{
		conn: conn,
	}
}

func (r *ordersRepository) Create(ctx context.Context, order *model.OrderInfo) (uuid.UUID, error) {
	query := fmt.Sprintf(`
        INSERT INTO %s (
            id,
            user_id,
            shop_id,
            cart,
            summary_volume,
            price,
            status
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `, tableName)

	_, err := r.conn.Exec(ctx, query,
		order.ID,
		order.UserID,
		order.ShopID,
		order.Cart,
		order.SummaryVolume,
		order.Price,
		order.Status,
	)

	if err != nil {
		return uuid.Nil, fmt.Errorf("repo create order: %w", err)
	}

	return order.ID, nil
}

func (r *ordersRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	query := fmt.Sprintf(`
		SELECT
			id,
			user_id,
			shop_id,
			cart,
			summary_volume,
			price,
			status,
			created_at,
			updated_at
		FROM %s
		WHERE id = $1
	`, tableName)

	var order model.Order

	err := r.conn.QueryRow(ctx, query, id).Scan(
		&order.Info.ID,
		&order.Info.UserID,
		&order.Info.ShopID,
		&order.Info.Cart,
		&order.Info.SummaryVolume,
		&order.Info.Price,
		&order.Info.Status,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrOrderNotFound
		}
		return nil, fmt.Errorf("get order by id: %w", err)
	}

	return &order, nil
}

func (r *ordersRepository) Update(ctx context.Context, order *model.OrderInfo) error {
	query := fmt.Sprintf(`UPDATE %s SET status = $1::order_status WHERE id = $2`, tableName)

	res, err := r.conn.Exec(ctx, query, string(order.Status), order.ID)
	if err != nil {
		return fmt.Errorf("update order: %w", err)
	}

	if res.RowsAffected() == 0 {
		return ErrOrderNotFound
	}

	return nil
}
