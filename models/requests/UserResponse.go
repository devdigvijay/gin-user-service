package requests

type CreateUserResponse struct {
	Id       string `json:"Id"`
	UserName string `json:"UserName"`
	IsActive bool   `json:"IsActive"`
}
