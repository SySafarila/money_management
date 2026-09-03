package dtos

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

func (RegisterRequest) TableName() string {
	return "users"
}

type LoginResponse struct {
	Token      string `json:"token"`
	ValidUntil int64  `json:"valid_until"`
}

type RegisterResponse struct {
	Username string `json:"username"`
}
