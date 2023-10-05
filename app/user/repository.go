package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Postgres struct {
}

func NewPostgres() Postgres {
	return Postgres{}
}

func (p Postgres) CreateUser(ctx context.Context, tx *sqlx.Tx, user User) (uuid.UUID, error) {

	userSQL := `INSERT INTO "users" ("email", "password", "created_at" ) VALUES ($1, $2, now()) RETURNING id`
	var userId uuid.UUID
	if err := tx.QueryRowContext(ctx, userSQL, user.Email, user.Password).
		Scan(&userId); err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			// Check if it's a duplicate key violation error
			if pgErr.Code == "23505" {
				return uuid.UUID{}, &CustomError{
					ErrorCode:  40901,
					StatusCode: 409,
					Message:    "email already exists",
				}
			} else {
				return uuid.UUID{}, &CustomError{
					ErrorCode:  50001,
					StatusCode: 500,
					Message:    "repository error",
				}
			}
		}
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
