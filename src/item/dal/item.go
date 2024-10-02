package dal

import (
	"context"
	"database/sql"
	"time"

	"github.com/benosborntech/feedme/common/types"
)

func QueryItemById(ctx context.Context, db *sql.DB, itemId int) (*types.Item, error) {
	var item types.Item

	row := db.QueryRowContext(ctx, "SELECT * FROM items WHERE id = ? LIMIT 1", itemId)
	if err := row.Scan(&item.Id, &item.Location, &item.ItemType, &item.Quantity, &item.ExpiresAt, &item.CreatedBy, &item.BusinessId, &item.UpdatedAt, &item.CreatedAt); err != nil {
		return nil, err
	}

	return &item, nil
}

func CreateItem(ctx context.Context, db *sql.DB, item *types.Item) (*types.Item, error) {
	var outItem types.Item

	row := db.QueryRowContext(
		ctx,
		"INSERT INTO items (location, item_type, quantity, expires_at, created_by, business_id) VALUES (?, ?, ?) RETURNING id, location, item_type, quantity, expires_at, created_by, business_id, updated_at, created_at",
		item.Location, item.ItemType, item.Quantity, item.ExpiresAt, item.CreatedBy, item.BusinessId,
	)
	if err := row.Scan(&item.Id, &item.Location, &item.ItemType, &item.Quantity, &item.ExpiresAt, &item.CreatedBy, &item.BusinessId, &item.UpdatedAt, &item.CreatedAt); err != nil {
		return nil, err
	}

	return &outItem, nil
}

func QueryItemFromUserIdAndTimestamp(ctx context.Context, db *sql.DB, idFrom int, timeFrom time.Time) ([]*types.Item, error) {
	rows, err := db.QueryContext(ctx, "SELECT * FROM items WHERE id >= ? AND created_at > ? ORDER BY created_at DESC", idFrom, timeFrom)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []*types.Item{}

	for rows.Next() {
		var item types.Item
		if err := rows.Scan(&item.Id, &item.Location, &item.ItemType, &item.Quantity, &item.ExpiresAt, &item.CreatedBy, &item.BusinessId, &item.UpdatedAt, &item.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	return items, nil
}
