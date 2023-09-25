package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
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
	result["refresh_token"] = refresToken
	return result
}

func RefereshJWT(accessToken *jwt.Token, refreshToken *jwt.Token) map[string]any {
	var result = map[string]any{}
	expTime, err := refreshToken.Claims.GetExpirationTime()
	logrus.Info(expTime)
	if err != nil {
		logrus.Error("get token expiration error", err.Error())
		return nil
	}
	if refreshToken.Valid && expTime.Time.Compare(time.Now()) > 0 {
		var newClaim = accessToken.Claims.(jwt.MapClaims)

		newClaim["iat"] = time.Now().Unix()
		newClaim["exp"] = time.Now().Add(time.Hour * 1).Unix()

		var newToken = jwt.NewWithClaims(accessToken.Method, newClaim)
		newSignedToken, _ := newToken.SignedString(accessToken.Signature)

		var newRefreshClaim = refreshToken.Claims.(jwt.MapClaims)
		newRefreshClaim["exp"] = time.Now().Add(time.Hour * 24).Unix()

		var newRefreshToken = jwt.NewWithClaims(refreshToken.Method, newRefreshClaim)
		newSignedRefreshToken, _ := newRefreshToken.SignedString(refreshToken.Signature)

		result["access_token"] = newSignedToken
		result["refresh_token"] = newSignedRefreshToken
		return result
	}

	return nil
}

func generateToken(signKey string, id string) string {
	var claims = jwt.MapClaims{}
	claims["id"] = id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Second * 10).Unix()

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
		var claims = token.Claims
		expTime, _ := claims.GetExpirationTime()
		fmt.Println(expTime.Time.Compare(time.Now()))
		if expTime.Time.Compare(time.Now()) > 0 {

			return token.Claims
		}

		logrus.Error("Token expired")
		return nil

	}
	return nil
}
