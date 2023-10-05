package user

import (
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CustomError struct {
	ErrorCode  int
	StatusCode int
	Message    string
}

func (e *CustomError) Error() string {
	return e.Message
}

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
		return u, &CustomError{
			ErrorCode:  40001,
			StatusCode: 400,
			Message:    "invalid request",
		}
	}

	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(emailPattern)

	if !(regex.MatchString(req.Email)) {
		return u, &CustomError{
			ErrorCode:  40002,
			StatusCode: 400,
			Message:    "email is not valid",
		}
	}

	if req.Password == "" {
		return u, &CustomError{
			ErrorCode:  40001,
			StatusCode: 400,
			Message:    "invalid request",
		}
	}

	if len(req.Password) < 8 {
		return u, &CustomError{
			ErrorCode:  40003,
			StatusCode: 400,
			Message:    "password is not valid",
		}
	}

	u.Email = req.Email
	u.Password = req.Password
	return u, nil

}

func (m Merchant) FromRegisterToMerchant(req Register) (Merchant, error) {
	if req.MerchantName == "" {
		return m, &CustomError{
			ErrorCode:  40001,
			StatusCode: 400,
			Message:    "invalid request",
		}
	}
	m.Name = req.MerchantName

	return m, nil

}
