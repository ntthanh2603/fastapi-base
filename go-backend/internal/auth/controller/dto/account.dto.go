package dto

type UserProfileReq struct {
	// AccountID int64 `json:"account_id" binding:"required"`
	AccountID int64 `json:"account_id" validate:"required,min=1,max=50"`
}

type UserProfileRes struct {
	AccountID int64 `json:"account_id" binding:"required"`
	Data      struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refreshToken"`
		TTL          int64  `json:"ttl"`
		AccountInfo  struct {
			Id       int64  `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
			Status   int    `json:"status"`
			Lang     string `json:"lang"`
		} `json:"accountInfo"`
	} `json:"data"`
}

// REGISTER DTOs
type UserRegisterReq struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Lang     string `json:"lang" validate:"required,min=2,max=5"`
}

type UserRegisterRes struct {
	AccountID int64  `json:"account_id"`
	Message   string `json:"message"`
}

// LOGIN DTOs

type UserLoginReq struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

type UserLoginRes struct {
	AccountID int64 `json:"account_id" binding:"required"`
	Data      struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refreshToken"`
		TTL          int64  `json:"ttl"`
		AccountInfo  struct {
			Id       int64  `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
			Status   int    `json:"status"`
			Lang     string `json:"lang"`
		} `json:"accountInfo"`
	} `json:"data"`
}
