package dal

import (
	"context"
	"database/sql"

	"github.com/benosborntech/feedme/common/types"
)

func GetUserByUserId(ctx context.Context, db *sql.DB, userId int) (*types.User, error) {
	var user types.User

	row := db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = ? LIMIT 1", userId)
	if err := row.Scan(&user.Id, &user.Email, &user.Name, &user.UpdatedAt, &user.CreatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(ctx context.Context, db *sql.DB, user *types.User) (*types.User, error) {
	if user == nil {
		return nil, nil
	}

	var outUser types.User

	row := db.QueryRowContext(ctx, "INSERT INTO users (id, email, name) VALUES (?, ?, ?) RETURNING id, email, name, updated_at, created_at", user.Id, user.Email, user.Name)
	if err := row.Scan(&user.Id, &user.Email, &user.Name, &user.UpdatedAt, &user.CreatedAt); err != nil {
		return nil, err
	}

	return &outUser, nil
}
