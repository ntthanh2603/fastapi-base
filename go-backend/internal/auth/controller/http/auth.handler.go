package http

import (
	"fmt"
	"net/http"

	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/application/service"
	appDto "github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/application/service/dto"
	ctlDto "github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/controller/dto"
	_ "github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/auth/domain/model/entity"
	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/pkg/response"
	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// GetUserProfile godoc
// @Summary get user profile
// @Description return user profile information by AccountID
// @Tags Auth
// @Accept json
// @Produce json
// @Param X-Sign header string true "HMAC Signature for request authentication"
// @Param X-Request-Time header string true "Request timestamp for validity check"
// @Param request body ctlDto.UserProfileReq true "User Profile Request"
// @Success 200 {object} entity.Account
// @Failure 400 {object} response.APIError
// @Router /auth/profile [post]
func (ah *AuthHandler) GetUserProfile(ctx *gin.Context) (res interface{}, err error) {
	fmt.Println("---> GetUserProfile")
	// // 1: get sign from request header
	// clientSign := ctx.GetHeader("X-Sign")
	// if clientSign != expectedSign {
	// 	return nil, response.NewAPIError(http.StatusUnauthorized, "Unauthorized", "Invalid sign")
	// }

	// // 2: validate request validity duration
	// requestTimeStr := ctx.GetHeader("X-Request-Time")
	// requestTime, err := strconv.ParseInt(requestTimeStr, 10, 64) // int64
	// if err != nil {
	// 	return nil, response.NewAPIError(http.StatusBadRequest, "Invalid request", "Invalid request time format")
	// }

	// now := time.Now().Unix()
	// log.Println("---> GetUserProfile | Current Time: ", now)
	// log.Println("---> GetUserProfile | Request Time: ", requestTime)
	// log.Println("---> GetUserProfile | Request Validity Duration: ", int64(requestValidityDuration.Seconds()), "seconds")
	// // Check if the request is within the validity duration
	// if now-requestTime > int64(requestValidityDuration.Seconds()) {
	// 	return nil, response.NewAPIError(http.StatusBadRequest, "Invalid request", "Request is too old")
	// }

	var req ctlDto.UserProfileReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, response.NewAPIError(http.StatusOK, "Invalid request", err.Error())
	}
	validation, exists := ctx.Get("validation")
	if !exists {
		return nil, response.NewAPIError(http.StatusOK, "Invalid request", "Validation not found in context")
	}
	fmt.Println("---> GetUserProfile | AccountID ", req.AccountID)

	if apiErr := utils.ValidateStruct(req, validation.(*validator.Validate)); apiErr != nil {
		return nil, apiErr
	}

	account, err := ah.service.GetAccountById(ctx, req.AccountID)
	if err != nil {
		return nil, response.NewAPIError(http.StatusOK, "User not found", err.Error())
	}
	fmt.Println("---> GetUserProfile | AccountID | account ", account)
	return account, nil
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Creates a new user account with the provided details.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body ctlDto.UserRegisterReq true "User Registration Request Details"
// @Success 201 {object} object{account_id=string} "User registered successfully, returns the new account ID" // SỬA Ở ĐÂY
// @Failure 400 {object} response.APIError "Invalid request (e.g., validation error, bad payload)"
// @Failure 409 {object} response.APIError "Conflict (e.g., email or username already exists)"
// @Failure 500 {object} response.APIError "Internal server error (e.g., registration process failed)"
// @Router /auth/register [post]
func (ah *AuthHandler) RegisterUser(ctx *gin.Context) (res interface{}, err error) {
	fmt.Println("---> RegisterUser")

	var req ctlDto.UserRegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, response.NewAPIError(http.StatusOK, "Invalid request", err.Error())
	}
	validation, exists := ctx.Get("validation")
	if !exists {
		return nil, response.NewAPIError(http.StatusOK, "Invalid request", "Validation not found in context")
	}

	if apiErr := utils.ValidateStruct(req, validation.(*validator.Validate)); apiErr != nil {
		return nil, apiErr
	}

	// Create account -> dto application
	account := appDto.AccountAppDTO{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Lang:     req.Lang,
		Status:   1,
	}

	accountId, err := ah.service.Create(ctx, account)
	if err != nil {
		return nil, response.NewAPIError(http.StatusOK, "Registration failed", err.Error())
	}

	return accountId, nil
}
