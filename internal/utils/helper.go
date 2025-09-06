package utils

import (
	"errors"
	"time"

	"sapaUMKM-backend/config/env"
	"sapaUMKM-backend/internal/types/dto"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(ID string, username string, email string) (string, error) {
	expirationTime := time.Now().Add(3 * 24 * time.Hour)
	claims := jwt.MapClaims{
		"id":       ID,
		"username": username,
		"email":    email,
		"exp":      expirationTime.Unix(),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(env.Cfg.Server.JWTSecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(jwtToken string) (dto.UserData, error) {
	token, _ := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("parsing token error occured")
		}
		return []byte(env.Cfg.Server.JWTSecretKey), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return dto.UserData{}, errors.New("token is invalid")
	}

	return dto.UserData{
		ID:       claims["id"].(string),
		Username: claims["username"].(string),
		Email:    claims["email"].(string),
	}, nil
}
