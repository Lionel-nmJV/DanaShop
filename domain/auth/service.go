package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type repository interface {
	userRepository
	merchantRepository
}

type userRepository interface {
	createUser(ctx context.Context, tx *sqlx.Tx, user User) (userId uuid.UUID, err error)
	getUserByEmail(ctx context.Context, db *sqlx.DB, email string) (user User, err error)
}

type merchantRepository interface {
	createMerchant(ctx context.Context, tx *sqlx.Tx, merchant Merchant) (err error)
	getMerchantIdByUserId(ctx context.Context, db *sqlx.DB, userId uuid.UUID) (uuid.UUID, error)
}

type UserService struct {
	repo repository
	db   *sqlx.DB
}

func newService(repo repository, db *sqlx.DB) UserService {
	return UserService{
		repo: repo,
		db:   db,
	}
}

func (u UserService) Register(ctx context.Context, user User, merchant Merchant) (err error) {
	if err = user.hashPassword(); err != nil {
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

	userId, err := u.repo.createUser(ctx, tx, user)
	if err != nil {
		return err
	}

	newMerchant := Merchant{
		Name:   merchant.Name,
		UserId: userId,
	}

	if err = u.repo.createMerchant(ctx, tx, newMerchant); err != nil {
		return
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return
}

func (u UserService) login(ctx context.Context, userLogin User) (res LoginResponse, err error) {

	userDb, err := u.repo.getUserByEmail(ctx, u.db, userLogin.Email)
	if err != nil {
		return LoginResponse{}, err
	}

	ok, err := userDb.ValidatePasswordFromPlainText(userLogin)
	if !ok {
		return LoginResponse{}, err
	}

	merchantId, err := u.repo.getMerchantIdByUserId(ctx, u.db, userDb.Id)
	if err != nil {
		return LoginResponse{}, err
	}

	AccessToken, errToken := userDb.CreateToken(merchantId)
	if err != nil {
		return LoginResponse{}, errToken
	}

	return LoginResponse{AccessToken}, nil

}
