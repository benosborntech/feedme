package dal

import (
	"context"
	"database/sql"

	"github.com/benosborntech/feedme/common/types"
)

func QueryBusinessById(ctx context.Context, db *sql.DB, businessId int) (*types.Business, error) {
	var business types.Business

	row := db.QueryRowContext(ctx, "SELECT * FROM businesses WHERE id = ? LIMIT 1", businessId)
	if err := row.Scan(&business.Id, &business.Name, &business.Description, &business.Latitude, &business.Longitude, &business.CreatedBy, &business.UpdatedAt, &business.CreatedAt); err != nil {
		return nil, err
	}

	return &business, nil
}

func CreateBusiness(ctx context.Context, db *sql.DB, business *types.Business) (*types.Business, error) {
	var outBusiness types.Business

	row := db.QueryRowContext(
		ctx,
		"INSERT INTO businesses (name, description, latitude, longitude, created_by) VALUES (?, ?, ?) RETURNING id, name, description, latitude, longitude, created_by, updated_at, created_at",
		business.Name, business.Description, business.Latitude, business.Longitude, business.CreatedBy,
	)
	if err := row.Scan(&business.Id, &business.Name, &business.Description, &business.Latitude, &business.Longitude, &business.CreatedBy, &business.UpdatedAt, &business.CreatedAt); err != nil {
		return nil, err
	}

	return &outBusiness, nil
}
