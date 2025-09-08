package jwt

import (
	"auth/internal/model/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

func GenerateToken(info auth.UserInfo, secretKey []byte, duration time.Duration) (string, error) {
	claims := auth.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
		Email: info.Email,
		Role:  info.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func VerifyToken(tokenStr string, secretKey []byte) (*auth.UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr,
		&auth.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.Errorf("unexpected token signing method")
			}
			return secretKey, nil
		})
	if err != nil {
		return nil, errors.Errorf("invalid token: %s", err.Error())
	}

	claims, ok := token.Claims.(*auth.UserClaims)
	if !ok {
		return nil, errors.Errorf("invalid token claims")
	}
	return claims, nil

}
