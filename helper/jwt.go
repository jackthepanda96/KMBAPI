package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(signKey string, refreshKey string, userID string) map[string]any {
	var result = map[string]any{}
	var accessToken = generateToken(signKey, userID)
	if accessToken == "" {
		return nil
	}
	var refresToken = generateRefreshToken(refreshKey, accessToken)
	if refresToken == "" {
		return nil
	}
	result["access_token"] = accessToken
	result["refresh_toke"] = refresToken
	return result
}

func generateToken(signKey string, id string) string {
	var claims = jwt.MapClaims{}
	claims["id"] = 1
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := sign.SignedString([]byte(signKey))

	if err != nil {
		return ""
	}

	return validToken
}

func generateRefreshToken(signKey string, accessToken string) string {
	var claims = jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := sign.SignedString([]byte(signKey))

	if err != nil {
		return ""
	}

	return refreshToken
}

func ExtractToken(token *jwt.Token) any {
	if token.Valid {
		/*
			TODO: provide more validation here
			ex: add expiration checking
		*/
		return token.Claims
	}
	return nil
}
