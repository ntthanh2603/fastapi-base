package dto

type AccountAppDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Lang     string `json:"lang"`
	Status   int    `json:"status"`
}
