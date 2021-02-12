package user

// RegisterInput ...
type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
}

// LoginInput ...
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UsernameCheckInput ...
type UsernameCheckInput struct {
	Username string `json:"username" binding:"required"`
}
