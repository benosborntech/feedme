package dal

import (
	"context"
	"database/sql"
	"log"

	"github.com/benosborntech/feedme/common/types"
)

func GetUserByUserId(ctx context.Context, db *sql.DB, userId int) (*types.User, error) {
	rows, err := db.QueryContext(ctx, "SELECT * FROM users WHERE id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*types.User{}

	for rows.Next() {
		var user types.User
		if err := rows.Scan(&user.Id, &user.Email, &user.Name, &user.UpdatedAt, &user.CreatedAt); err != nil {
			log.Printf("failed to parse item, err=%v", err)

			continue
		}
		users = append(users, &user)
	}

	if len(users) == 0 {
		log.Printf("no users to retrieve")

		return nil, nil
	}

	return users[0], nil
}

func CreateUser(ctx context.Context, db *sql.DB, user *types.User) (*types.User, error) {
	if user == nil {
		return nil, nil
	}

	_, err := db.ExecContext(ctx, "INSERT INTO users (id, email, name) VALUES (?, ?, ?)", user.Id, user.Email, user.Name)
	if err != nil {
		return nil, err
	}

	return GetUserByUserId(ctx, db, user.Id)
}
