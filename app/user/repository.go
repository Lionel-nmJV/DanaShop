package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
}

func NewPostgres() Postgres {
	return Postgres{}
}

func (p Postgres) CreateUser(ctx context.Context, tx *sqlx.Tx, user User) (uuid.UUID, error) {

	userSQL := `INSERT INTO "users" ("email", "password", "created_at" ) VALUES ($1, $2, now()) RETURNING id`
	var userId uuid.UUID
	userResult := tx.QueryRowContext(ctx, userSQL, user.Email, user.Password)

	if err := userResult.Scan(&userId); err != nil {

		return uuid.UUID{}, err
	}

	return userId, nil
}

func (p Postgres) CreateMerchant(ctx context.Context, tx *sqlx.Tx, merchant Merchant) error {
	merchantSQL := `INSERT INTO "merchants" ("user_id", "name" , "created_at" ) VALUES ($1, $2, NOW())`
	_, err := tx.ExecContext(ctx, merchantSQL, merchant.UserId, merchant.Name)

	if err != nil {
		return err
	}

	return nil
}
