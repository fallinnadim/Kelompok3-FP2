package helper

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

func GenerateToken(ttl time.Duration, payload interface{}, secretJWTKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	now := time.Now().UTC()
	claim := token.Claims.(jwt.MapClaims)

	claim["sub"] = payload             // subject
	claim["exp"] = now.Add(ttl).Unix() // expires at
	claim["iat"] = now.Unix()          // issued at
	claim["nbf"] = now.Unix()          // not before

	tokenString, err := token.SignedString([]byte(secretJWTKey))
	if err != nil {
		return "", fmt.Errorf("Failed generating JWT Token %w", err)
	}
	return tokenString, nil
}

func ValidateToken(token string, signedJWTKey string) (interface{}, error) {
	tok, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected method : %s", t.Header["alg"])
		}
		return []byte(signedJWTKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("Invalid token %w", err)
	}
	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("Invalid token claims")
	}
	return claims["sub"], nil
}
