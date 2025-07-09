package response

type UserResponse struct {
	Id       int    `json:"Id"`
	Username string `json:"Username"`
	Email    string `json:"Email"`
}
