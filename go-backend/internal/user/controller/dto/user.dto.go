package dto

type UserProfileReq struct {
	// AccountID int64 `json:"account_id" binding:"required"`
	AccountID int64 `json:"account_id" validate:"required,min=1,max=50"`
}

type UserProfileRes struct {
	AccountID int64 `json:"account_id" binding:"required"`
	Data      struct {
		Token        string `json:"token" dc:"JWT token"`
		RefreshToken string `json:"refreshToken" dc:"Refresh token"`
		TTL          int64  `json:"ttl" dc:"Token lifetime (in seconds)"`
		AccountInfo  struct {
			Id       int64  `json:"id" dc:"Account ID"`
			Username string `json:"username" dc:"Username"`
			Email    string `json:"email" dc:"Email address"`
			Status   int    `json:"status" dc:"Account status"`
			Lang     string `json:"lang" dc:"Preferred language"`
		} `json:"accountInfo" dc:"Basic account information"`
	} `json:"data"`
}
