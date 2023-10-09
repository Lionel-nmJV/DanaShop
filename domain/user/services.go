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
	GetUserByEmail(ctx context.Context, db *sqlx.DB, email string) (user User, err error)
}

type MerchantRepository interface {
	CreateMerchant(ctx context.Context, tx *sqlx.Tx, merchant Merchant) (err error)
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

func (u UserService) login(ctx context.Context, userLogin User) (res LoginResponse, err error) {

	userDb, err := u.repo.GetUserByEmail(ctx, u.db, userLogin.Email)
	if err != nil {
		return LoginResponse{}, err
	}

	ok, err := userDb.ValidatePasswordFromPlainText(userLogin)
	if !ok {
		return LoginResponse{}, err
	}

	AccessToken, errToken := userDb.CreateToken()
	if err != nil {
		return LoginResponse{}, errToken
	}

	return LoginResponse{AccessToken}, nil

}
