package jwt

import (
	"time"

	jwtconfig "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/config/jwt"
	useraccount "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/models/account_model"
	"github.com/golang-jwt/jwt/v5"
)

// Generate new access token
func GenerateAccessToken(user useraccount.UserAccount) (int, string, string) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.Id,
		"exp":  time.Now().Add(3600 * time.Second).Unix(), //30 sec
		"role": user.Role,
	})
	accessTokenString, err := accessToken.SignedString(jwtconfig.JWT_KEY)
	if err != nil {
		return 400, "", "Could not generate access token"
	}

	return 200, accessTokenString, ""
}
