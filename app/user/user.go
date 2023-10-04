package user

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailEmpty     = errors.New("email is empty")
	ErrEmailInvalid   = errors.New("email is invalid")
	ErrPasswordEmpty  = errors.New("password is empty")
	ErrPasswordLength = errors.New("password must be at least 8 characters")
	ErrMerchantEmpty  = errors.New("merchant name is empty")
)

type User struct {
	Id        uuid.UUID `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Roles     string    `db:"roles"`
	ImageUrl  string    `db:"image_url"`
}

type Merchant struct {
	UserId uuid.UUID `db:"user_id"`
	Name   string    `db:"name"`
}

func NewUser() User {
	return User{}
}

func (u *User) HashPassword() (err error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashed)
	return
}

func (u User) FromRegisterToUser(req Register) (User, error) {
	if req.Email == "" {
		return u, ErrEmailEmpty
	}

	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(emailPattern)

	fmt.Println(req.Email)

	if !(regex.MatchString(req.Email)) {
		return u, ErrEmailInvalid
	}

	if req.Password == "" {
		return u, ErrPasswordEmpty
	}

	if len(req.Password) < 8 {
		return u, ErrPasswordLength
	}

	u.Email = req.Email
	u.Password = req.Password
	return u, nil

}

func (m Merchant) FromRegisterToMerchant(req Register) (Merchant, error) {
	if req.MerchantName == "" {
		return m, ErrMerchantEmpty
	}
	m.Name = req.MerchantName

	return m, nil

}
