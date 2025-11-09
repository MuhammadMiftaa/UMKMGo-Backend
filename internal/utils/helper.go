package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"

	"sapaUMKM-backend/config/env"
	"sapaUMKM-backend/internal/types/dto"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// ~ PasswordValidator checks if the password meets the criteria
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

// ~ PasswordHashing hashes the password using bcrypt
func PasswordHashing(str string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hashPassword), nil
}

// ~ ComparePass compares the hashed password with the provided password
func ComparePass(hashPassword, reqPassword string) bool {
	hash, pass := []byte(hashPassword), []byte(reqPassword)

	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}

// ~ GenerateToken creates a JWT token for the user
// ~ It includes user ID, name, email, role, role name, and permissions in the token claims.
func GenerateToken(user dto.Users) (string, error) {
	expirationTime := time.Now().Add(3 * 24 * time.Hour)
	claims := jwt.MapClaims{
		"id":          user.ID,
		"name":        user.Name,
		"email":       user.Email,
		"role":        user.RoleID,
		"role_name":   user.RoleName,
		"permissions": user.Permissions,
		"iat":         time.Now().Unix(),
		"exp":         expirationTime.Unix(),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(env.Cfg.Server.JWTSecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ~ VerifyToken checks the validity of the JWT token
// ~ It parses the token and extracts the claims, returning user data if valid.
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
		ID:       claims["id"].(float64),
		Name:     claims["name"].(string),
		Email:    claims["email"].(string),
		Role:     claims["role"].(float64),
		RoleName: claims["role_name"].(string),
	}, nil
}

// ~ GenerateOTP creates a random 6-digit OTP (One Time Password)
// ~ It uses the math/rand package to generate a random number between 0 and 999999, then formats it as a 6-digit string.
func GenerateOTP() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// ~ EmailValidator checks if the provided string is a valid email format
// ~ It uses a regular expression to validate the email format.
func EmailValidator(str string) bool {
	email_validator := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return email_validator.MatchString(str)
}

// ~ GenerateRequestID creates a random alphanumeric string of length 10
// ~ It uses a predefined character set and the math/rand package to generate the string.
func GenerateRequestID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 10

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// ~ GenerateFileName creates a standardized file name based on the original name and a prefix used for MinIO storage
func GenerateFileName(originalName, prefix string) string {
	return strings.ToLower(fmt.Sprintf("%s/%s_", strings.Join(strings.Split(originalName, " "), "_"), prefix))
}
