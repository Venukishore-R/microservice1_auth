package endpoints

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Status      int64
	Description string
}

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Status       int64
	AccessToken  string
	RefreshToken string
}
