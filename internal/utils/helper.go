package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"time"
	"unicode"

	"sapaUMKM-backend/config/env"
	"sapaUMKM-backend/config/log"
	"sapaUMKM-backend/internal/types/dto"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func PasswordValidator(str string) (bool, bool, bool) {
	var hasLetter, hasDigit, hasMinLen bool
	for _, char := range str {
		switch {
		case unicode.IsLetter(char):
			hasLetter = true
		case unicode.IsDigit(char):
			hasDigit = true
		}
	}

	if len(str) >= 8 {
		hasMinLen = true
	}

	return hasMinLen, hasLetter, hasDigit
}

func PasswordHashing(str string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hashPassword), nil
}

func ComparePass(hashPassword, reqPassword string) bool {
	hash, pass := []byte(hashPassword), []byte(reqPassword)

	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}

func GenerateToken(user dto.Users) (string, error) {
	expirationTime := time.Now().Add(3 * 24 * time.Hour)
	claims := jwt.MapClaims{
		"id":        user.ID,
		"name":      user.Name,
		"email":     user.Email,
		"role":      user.RoleID,
		"role_name": user.RoleName,
		"exp":       expirationTime.Unix(),
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

	log.Debug("ID type: " + fmt.Sprintf("%T", claims["id"]))
	log.Debug("Role type: " + fmt.Sprintf("%T", claims["role"]))
	log.Debug("Email type: " + fmt.Sprintf("%T", claims["email"]))
	log.Debug("Name type: " + fmt.Sprintf("%T", claims["name"]))
	log.Debug("RoleName type: " + fmt.Sprintf("%T", claims["role_name"]))

	return dto.UserData{
		ID:       claims["id"].(float64),
		Name:     claims["name"].(string),
		Email:    claims["email"].(string),
		Role:     claims["role"].(float64),
		RoleName: claims["role_name"].(string),
	}, nil
}

func GenerateOTP() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func EmailValidator(str string) bool {
	email_validator := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return email_validator.MatchString(str)
}
