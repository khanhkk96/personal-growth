package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateAccessToken(ttl time.Duration, payload interface{}, secretKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claim := token.Claims.(jwt.MapClaims)

	claim["sub"] = payload             //subject
	claim["exp"] = now.Add(ttl).Unix() //expired
	claim["iat"] = now.Unix()          //issued at
	claim["nbf"] = now.Unix()          //not before
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", fmt.Errorf("Generate JWT token failed: %w", err)
	}

	return tokenString, nil
}

func VerifyToken(token string, signedJWTKey string) (interface{}, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte(signedJWTKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("Invalid token: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)

	if !ok || !tok.Valid {
		return nil, fmt.Errorf("Invalid token claim")
	}

	return claims["sub"], nil
}

func GenerateTokens(ttl time.Duration, payload interface{}, secretKey string, rfKey string) (string, string, error) {
	// Access token (expires in 15 mins)
	accessClaims := jwt.MapClaims{
		"sub": payload,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(ttl).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessString, err := accessToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", "", err
	}

	// Refresh token (expires in 7 days)
	refreshClaims := jwt.MapClaims{
		"sub": payload,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshString, err := refreshToken.SignedString([]byte(rfKey))
	if err != nil {
		return "", "", err
	}

	return accessString, refreshString, nil
}

// Kiểm tra refresh token hợp lệ
func ValidateRefreshToken(tokenStr string, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}
