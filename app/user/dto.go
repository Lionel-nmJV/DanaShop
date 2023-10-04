package user

type Register struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	MerchantName string `json:"merchant_name"`
}
