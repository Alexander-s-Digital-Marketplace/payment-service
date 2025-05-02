package jwt

import (
	"time"

	jwtconfig "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/config/jwt"
	useraccount "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/models/account_model"
	"github.com/golang-jwt/jwt/v5"
)

// Generate new refresh token
func GenerateRefreshToken(user useraccount.UserAccount) (int, string, string) {

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.Id,
		"exp":  time.Now().Add(time.Hour * 24 * 7).Unix(), //7 days
		"role": user.Role,
	})
	refreshTokenString, err := refreshToken.SignedString(jwtconfig.JWT_KEY)
	if err != nil {
		return 400, "", "Could not generate refresh token"
	}
	return 200, refreshTokenString, ""
}
