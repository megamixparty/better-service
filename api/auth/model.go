package auth

type LoginRequest struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

type RegisterRequest struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Error string `json:"error,omitempty"`
}

type RegisterResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
