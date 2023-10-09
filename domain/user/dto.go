package user

type Register struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	MerchantName string `json:"merchant_name"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
