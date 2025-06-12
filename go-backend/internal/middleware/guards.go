package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex" // or base64
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/pkg/response"
	"github.com/gin-gonic/gin"
)

const SHARED_SECRET_KEY = "your-very-secret-and-long-key" // !!! in env
const requestValidityDuration = 200 * time.Second

// AuthGuardMiddlewareWithHMAC by HMAC
func AuthGuardMiddlewareWithHMAC() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientSign := ctx.GetHeader("X-Sign")
		requestTimeStr := ctx.GetHeader("X-Request-Time")

		if clientSign == "" || requestTimeStr == "" {
			log.Println("HMAC Auth: Missing X-Sign or X-Request-Time header")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.NewAPIError(http.StatusUnauthorized, "Unauthorized", "Missing required signature headers"))
			return
		}

		// 1. Kiểm tra Timestamp
		requestTime, err := strconv.ParseInt(requestTimeStr, 10, 64)
		if err != nil {
			log.Println("HMAC Auth: Invalid request time format")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response.NewAPIError(http.StatusBadRequest, "Invalid request", "Invalid request time format"))
			return
		}

		now := time.Now().Unix()
		if now-requestTime > int64(requestValidityDuration.Seconds()) || now-requestTime < -5 { // Chấp nhận trễ 5s
			log.Printf("HMAC Auth: Request timestamp out of bounds. Now: %d, ReqTime: %d\n", now, requestTime)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response.NewAPIError(http.StatusBadRequest, "Invalid request", "Request timestamp out of bounds"))
			return
		}

		// 2. Tái tạo String-to-Sign
		// Đọc body (nếu có) và chuẩn bị để đọc lại
		var bodyBytes []byte
		if ctx.Request.Body != nil {
			bodyBytes, err = io.ReadAll(ctx.Request.Body)
			if err != nil {
				log.Println("HMAC Auth: Error reading request body:", err)
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.NewAPIError(http.StatusInternalServerError, "Server Error", "Could not read request body"))
				return
			}
			// Tạo lại body để handler sau có thể đọc
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		stringToSign := buildStringToSign(ctx, requestTimeStr, bodyBytes)
		log.Printf("HMAC Auth: Server StringToSign: [%s]\n", strings.ReplaceAll(stringToSign, "\n", "\\n"))

		// 3. Tính toán HMAC phía server
		serverSign := calculateHMAC(stringToSign, SHARED_SECRET_KEY)
		log.Printf("HMAC Auth: ClientSign: %s, ServerSign: %s\n", clientSign, serverSign)

		// 4. So sánh chữ ký
		// Sử dụng hmac.Equal để so sánh an toàn, chống timing attacks
		if !hmac.Equal([]byte(clientSign), []byte(serverSign)) {
			log.Println("HMAC Auth: Invalid signature")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.NewAPIError(http.StatusUnauthorized, "Unauthorized", "Invalid signature"))
			return
		}

		log.Println("HMAC Auth: Signature verified successfully")
		ctx.Next()
	}
}

func buildStringToSign(ctx *gin.Context, requestTimeStr string, bodyBytes []byte) string {
	method := ctx.Request.Method
	path := ctx.Request.URL.Path // Chỉ path, không có query string

	// Sắp xếp query parameters
	queryParams := ctx.Request.URL.Query()
	var sortedQueryKeys []string
	for k := range queryParams {
		sortedQueryKeys = append(sortedQueryKeys, k)
	}
	sort.Strings(sortedQueryKeys)

	var canonicalQueryParts []string
	for _, k := range sortedQueryKeys {
		// Gin's queryParams[k] is a []string, handle multiple values if necessary
		// For simplicity, we'll take the first one or join them if you expect multiple
		// For this example, let's assume single values for simplicity or take first
		if len(queryParams[k]) > 0 {
			canonicalQueryParts = append(canonicalQueryParts, fmt.Sprintf("%s=%s", k, queryParams[k][0]))
		}
	}
	canonicalQueryString := strings.Join(canonicalQueryParts, "&")

	// Xử lý body:
	//  - Nếu body rỗng, có thể dùng chuỗi rỗng hoặc một hash cố định của chuỗi rỗng.
	//  - Nếu body không rỗng, nên hash body (ví dụ SHA256) rồi đưa hash đó vào stringToSign.
	//    Điều này an toàn hơn là đưa raw body (có thể rất lớn) vào stringToSign.
	//    Ở đây, để đơn giản, tôi sẽ đưa raw bodyBytes (đã được convert sang string) vào,
	//    nhưng HASHING BODY LÀ BEST PRACTICE.
	bodyString := ""
	if len(bodyBytes) > 0 {
		bodyString = string(bodyBytes) // CẢNH BÁO: Nếu body không phải UTF-8, sẽ có vấn đề. Nên hash body!
		// Ví dụ hash body:
		// bodyHasher := sha256.New()
		// bodyHasher.Write(bodyBytes)
		// bodyString = hex.EncodeToString(bodyHasher.Sum(nil))
	}

	// Thứ tự phải nhất quán giữa client và server
	// Ví dụ: METHOD\nPATH\nTIMESTAMP\nSORTED_QUERY_STRING\nBODY_STRING_OR_HASH
	parts := []string{
		method,
		path,
		requestTimeStr,
		canonicalQueryString,
		bodyString, // Hoặc hash của body
	}
	return strings.Join(parts, "\n")
}

func calculateHMAC(data string, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil)) // Hoặc base64.StdEncoding.EncodeToString
}
