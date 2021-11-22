package middleware

import (
	"fmt"
	"go-jwt-access-refresh-token/databaseConnection"
	"go-jwt-access-refresh-token/errorsHandler"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Authorization(c *gin.Context) {

	// แกะ Bearer ออก เอาแค่เฉพาะ token
	token, err := getToken(c)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": err})
		c.Abort()
		return
	}

	// ส่ง Token ไปตรวจสอบว่าได้รับสิทธิ์เข้าใช้งานหรือไม่
	isToken, err := verifyAccessToken(token)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": err})
		c.Abort()
		return
	}

	if isToken {
		c.Next()
	}

}

func getToken(c *gin.Context) (string, error) {

	const BEARER_SCHEMA = "Bearer "

	AUTH_HEADER := c.GetHeader("Authorization")
	if len(AUTH_HEADER) == 0 {
		return "", errorsHandler.NewMessageAndStatusCode(http.StatusUnauthorized, "authorization key in header not found")
	}

	if strings.HasPrefix(AUTH_HEADER, BEARER_SCHEMA) {
		tokenString := AUTH_HEADER[len(BEARER_SCHEMA):]
		return tokenString, nil
	} else {
		return "", errorsHandler.NewMessageAndStatusCode(http.StatusUnauthorized, "Bearer signature key was not found")
	}

}

func verifyAccessToken(token string) (bool, error) {

	rdb := databaseConnection.NewDatabaseConnection().RedisConnection()
	defer rdb.Close()

	claims, err := getAccessTokenClaims(token)
	if err != nil {
		return false, err
	}

	_, err = rdb.Get(ctx, claims.AccessTokenUUID).Result()
	if err != nil {
		return false, errorsHandler.NewUnauthorizedError()
	}

	return true, nil
}

func getAccessTokenClaims(encodedToken string) (*ClaimsToken, error) {

	parseToken, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errorsHandler.NewMessageAndStatusCode(http.StatusUnauthorized, fmt.Sprint(token.Header["alg"]))
		}
		return []byte(viper.GetString("token.secretKey")), nil
	})
	if err != nil {
		return nil, err
	}

	claimsToken := &ClaimsToken{}
	parseClaims := parseToken.Claims.(jwt.MapClaims)

	if len(parseClaims["issuer"].(string)) != 0 {
		claimsToken.Issuer = parseClaims["issuer"].(string)
	}

	if len(parseClaims["subject"].(string)) != 0 {
		claimsToken.Subject = parseClaims["subject"].(string)
	}

	if len(parseClaims["role"].(string)) != 0 {
		claimsToken.Role = parseClaims["role"].(string)
	}

	if len(parseClaims["access_token_uuid"].(string)) != 0 {
		claimsToken.AccessTokenUUID = parseClaims["access_token_uuid"].(string)
	}

	if len(parseClaims["refresh_token_uuid"].(string)) != 0 {
		claimsToken.RefreshTokenUUID = parseClaims["refresh_token_uuid"].(string)
	}

	if len(parseClaims["expiration_time"].(string)) != 0 {
		claimsToken.ExpiresAccessToken = parseClaims["expiration_time"].(string)
	}

	fmt.Println(parseToken)

	fmt.Println(parseToken)

	fmt.Println(parseToken)


	return claimsToken, nil
}
