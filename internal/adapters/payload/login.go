package payload

// LoginRequest representa la solicitud de inicio de sesión (login).
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
