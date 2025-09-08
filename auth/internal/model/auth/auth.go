package auth

type UserInfo struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}
