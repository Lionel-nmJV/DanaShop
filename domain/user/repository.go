package user

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type postgres struct {
}

func newPostgres() postgres {
	return postgres{}
}

func (p postgres) createUser(ctx context.Context, tx *sqlx.Tx, user User) (uuid.UUID, error) {

	userSQL := `INSERT INTO "users" ("email", "password", "created_at", "updated_at","roles" ) VALUES ($1, $2, now(), now(),'merchant') RETURNING id`
	var userId uuid.UUID
	if err := tx.QueryRowContext(ctx, userSQL, user.Email, user.Password).
		Scan(&userId); err != nil {
		if pgErr, ok := err.(*pq.Error); ok {

			if pgErr.Code == "23505" {
				return uuid.UUID{}, newCustomError(40901, 409, "email already taken")
			} else {
				return uuid.UUID{}, newCustomError(50001, 500, "repository error")
			}
		}
	}
	return userId, nil
}

func (p postgres) createMerchant(ctx context.Context, tx *sqlx.Tx, merchant Merchant) error {
	merchantSQL := `INSERT INTO "merchants" ("user_id", "name" , "created_at" ,"updated_at") VALUES ($1, $2, NOW(), NOW())`
	_, err := tx.ExecContext(ctx, merchantSQL, merchant.UserId, merchant.Name)

	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				return newCustomError(40902, 409, "merchant name already taken")
			} else {
				return newCustomError(50001, 500, "repository error")
			}
		}
	}
	return nil
}

func (p postgres) getUserByEmail(ctx context.Context, db *sqlx.DB, email string) (User, error) {
	userSQL := `SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1 and roles='merchant'`
	var user User
	err := db.QueryRowContext(ctx, userSQL, email).
		Scan(&user.Id, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		return User{}, newCustomError(40101, 401, "email or password invalid")
	case err != nil:
		return User{}, newCustomError(50001, 500, "repository error")
	}
	return user, nil
}
