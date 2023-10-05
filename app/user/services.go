package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	UserRepository
	MerchantRepository
}

type UserRepository interface {
	CreateUser(ctx context.Context, tx *sqlx.Tx, user User) (userId uuid.UUID, err error)
}

type MerchantRepository interface {
	CreateMerchant(ctx context.Context, tx *sqlx.Tx, merchant Merchant) (err error)
}

type Transaction interface {
	Begin(ctx context.Context) error
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
	CommitOrRollback(ctx context.Context, err error) error
}

type UserService struct {
	repo Repository
	db   *sqlx.DB
}

func NewService(repo Repository, db *sqlx.DB) UserService {
	return UserService{
		repo: repo,
		db:   db,
	}
}

func (u UserService) Register(ctx context.Context, user User, merchant Merchant) (err error) {
	if err = user.HashPassword(); err != nil {
		return
	}

	tx, err := u.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	userId, err := u.repo.CreateUser(ctx, tx, user)
	if err != nil {
		return err
	}

	newMerchant := Merchant{
		Name:   merchant.Name,
		UserId: userId,
	}

	if err = u.repo.CreateMerchant(ctx, tx, newMerchant); err != nil {
		return
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return
}

// func (a authService) login(ctx context.Context, req Auth) (item Auth, err error) {
// 	auth, err := a.repo.findByEmail(ctx, req.Email)
// 	if err != nil {
// 		return
// 	}

// 	ok, err := auth.ValidatePasswordFromPlainText(req.Password)
// 	if err != nil {
// 		return req, err
// 	}

// 	if !ok {
// 		return req, ErrInvalidPassword
// 	}

// 	return auth, nil

// }
