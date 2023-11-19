package auth

import (
	"os"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func newUser() User {
	return User{}
}

func (u *User) hashPassword() (err error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashed)
	return
}

func (u User) fromRegisterToUser(req register) (User, error) {
	if req.Email == "" {
		return u, newCustomError(40001, 400, "invalid request")
	}

	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(emailPattern)

	if !(regex.MatchString(req.Email)) {
		return u, newCustomError(40002, 400, "email is not valid")
	}

	if len(req.Password) < 8 {
		return u, newCustomError(40003, 400, "password is not valid")
	}

	u.Email = req.Email
	u.Password = req.Password
	return u, nil

}

func (m Merchant) FromRegisterToMerchant(req register) (Merchant, error) {
	if req.MerchantName == "" {
		return m, newCustomError(40001, 400, "invalid request")
	}
	m.Name = req.MerchantName

	return m, nil

}

func (u User) FromLogin(req login) (User, error) {
	if req.Email == "" {
		return u, newCustomError(40001, 400, "invalid request")
	}

	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(emailPattern)

	if !(regex.MatchString(req.Email)) {
		return u, newCustomError(40002, 400, "email is not valid")
	}

	if len(req.Password) < 8 {
		return u, newCustomError(40003, 400, "password is not valid")
	}

	u.Email = req.Email
	u.Password = req.Password
	return u, nil

}

func (u User) ValidatePasswordFromPlainText(userLogin User) (ok bool, err error) {
	errCompare := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(userLogin.Password))
	if errCompare != nil {
		return ok, newCustomError(40102, 401, "email or password invalid")
	}
	ok = true
	return
}

func (u User) CreateToken(merchantId uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"merchant_id": merchantId,
		"user_id":     u.Id,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
