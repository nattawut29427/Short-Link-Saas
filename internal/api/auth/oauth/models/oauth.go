package models

type GoogleOAuthRequest struct {
	Code string `json:"code"`
}

type OAuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}