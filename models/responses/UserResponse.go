package responses

type CreateUserRequest struct {
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	UserName   string `json:"userName"`
	Age        int    `json:"Age"`
	Password   string `json:"Password"`
	IsActive   bool   `json:"isActive"`
}
