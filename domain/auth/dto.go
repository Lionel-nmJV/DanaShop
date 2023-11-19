package auth

type register struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	MerchantName string `json:"merchant_name"`
}

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
