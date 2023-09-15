package repositories

import (
	"context"
	"database/sql"

	"github.com/wildanfaz/go-challenge/internal/app/entities"
)

type Auth interface {
	Register(ctx context.Context, user entities.User) error
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
}

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) Auth {
	return &AuthRepo{db: db}
}

func (a *AuthRepo) Register(ctx context.Context, user entities.User) error {
	q := `INSERT INTO users(email, password) VALUES(?,?)`

	if _, err := a.db.ExecContext(ctx, q, user.Email, user.Password); err != nil {
		return err
	}

	return nil
}

func (a *AuthRepo) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	var (
		user entities.User
	)

	q := `SELECT id, email, password, balance, created_at, updated_at FROM users WHERE email=? LIMIT 1`

	rows, err := a.db.QueryContext(ctx, q, email)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&user.ID, &user.Email, &user.Password,
			&user.Balance, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
	}

	return &user, nil
}
