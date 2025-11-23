package utils

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image/png"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"UMKMGo-backend/config/env"
	"UMKMGo-backend/internal/types/dto"

	"github.com/dgrijalva/jwt-go"
	"github.com/skip2/go-qrcode"
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

// ~ GenerateWebToken creates a JWT token for the user
// ~ It includes user ID, name, email, role, role name, and permissions in the token claims.
func GenerateWebToken(user dto.Users) (string, error) {
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
		"is_admin":    true,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(env.Cfg.Server.JWTSecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ~ GenerateMobileToken creates a JWT token for mobile users
// ~ It includes user ID, name, business name, email, phone, kartu type, and expiration time in the token claims.
func GenerateMobileToken(user dto.UMKMMobile) (string, error) {
	expirationTime := time.Now().Add(3 * 24 * time.Hour)
	claims := jwt.MapClaims{
		"id":            user.ID,
		"name":          user.Fullname,
		"business_name": user.BusinessName,
		"email":         user.Email,
		"phone":         user.Phone,
		"kartu_type":    user.KartuType,
		"iat":           time.Now().Unix(),
		"exp":           expirationTime.Unix(),
		"is_admin":      false,
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

	isAdmin, ok := claims["is_admin"].(bool)
	if !ok {
		return dto.UserData{}, errors.New("token missing or invalid is_admin claim")
	}
	var userData dto.UserData

	userData.ID, _ = claims["id"].(float64)
	userData.Name, _ = claims["name"].(string)
	userData.Email, _ = claims["email"].(string)

	if isAdmin {
		userData.Role, _ = claims["role"].(float64)
		userData.RoleName, _ = claims["role_name"].(string)
	} else {
		userData.BusinessName, _ = claims["business_name"].(string)
		userData.KartuType, _ = claims["kartu_type"].(string)
		userData.Phone, _ = claims["phone"].(string)
	}

	return userData, nil
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

// ~ NIKValidator validates the Indonesian National Identity Number (NIK)
func NIKValidator(nik string) error {
	if len(nik) != 16 {
		return errors.New("NIK harus 16 digit")
	}

	// hanya angka
	matched, err := regexp.MatchString(`^[0-9]{16}$`, nik)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("NIK harus hanya berisi angka")
	}

	// tanggal lahir posisi: 7-8 (0-based index 6-7) untuk “tanggal”
	// bulan 9-10 (index 8-9), tahun 11-12 (index 10-11)
	tStr := nik[6:8]
	mStr := nik[8:10]
	yStr := nik[10:12]

	tInt, err := strconv.Atoi(tStr)
	if err != nil {
		return errors.New("Format tanggal lahir tidak valid")
	}
	mInt, err := strconv.Atoi(mStr)
	if err != nil {
		return errors.New("Format bulan lahir tidak valid")
	}
	yInt, err := strconv.Atoi(yStr)
	if err != nil {
		return errors.New("Format tahun lahir tidak valid")
	}

	// jika tanggal > 40 maka perempuan → kurangi 40
	if tInt > 40 {
		tInt -= 40
	}

	if tInt < 1 || tInt > 31 {
		return errors.New("Tanggal lahir di NIK tidak valid")
	}
	if mInt < 1 || mInt > 12 {
		return errors.New("Bulan lahir di NIK tidak valid")
	}

	// tentukan tahun lengkap (asumsi 1900 atau 2000)
	now := time.Now().Year()
	yy := yInt
	fullYear := 0
	if yy <= (now % 100) {
		fullYear = 2000 + yy
	} else {
		fullYear = 1900 + yy
	}
	// validasi tahun masuk logika (misal: tidak di masa depan)
	if fullYear > now {
		return errors.New("Tahun lahir di NIK lebih besar dari sekarang")
	}

	// Jika semua oke
	return nil
}

// ~ NormalizePhone normalizes an Indonesian phone number to a standard format
func NormalizePhone(phone string) (string, error) {
	if phone == "" {
		return "", errors.New("nomor kosong")
	}

	// Hilangkan spasi, dash, titik, dsb — hanya sisakan + dan digit
	re := regexp.MustCompile(`[^0-9+]+`)
	p := re.ReplaceAllString(phone, "")
	p = strings.TrimSpace(p)

	if p == "" {
		return "", errors.New("nomor tidak valid")
	}

	// Jika sudah diawali '8' → return apa adanya
	if strings.HasPrefix(p, "8") {
		return p, nil
	}

	// +62XXXXXXXX
	if strings.HasPrefix(p, "+62") {
		p = p[3:]
	} else if strings.HasPrefix(p, "62") {
		// 62XXXXXXXX
		p = p[2:]
	} else if strings.HasPrefix(p, "0") {
		// 08XXXXXXXX
		p = p[1:]
	} else {
		return "", errors.New("format nomor tidak dikenali (harus 0, 62, +62, atau 8)")
	}

	// Validasi dasar — panjang nomor minimal 9 digit (contoh: 812xxxxxx)
	if len(p) < 9 {
		return "", errors.New("panjang nomor tidak valid")
	}

	return p, nil
}

// ~ RandomString generates a random hexadecimal string of the specified size
func RandomString(size int) string {
	b := make([]byte, 32)
	rand.Read(b)

	if size <= 0 {
		return fmt.Sprintf("%x", b)
	}

	return fmt.Sprintf("%x", b)[:size]
}

// ~ MaskMiddle masks the middle part of a string with "XXXXXX"
func MaskMiddle(s string) string {
	mask := "XXXXXXXX"
	n := len(s)

	// Jika terlalu pendek atau hampir pendek
	if n <= len(mask) {
		return mask
	}

	// Tentukan posisi potong secara dinamis
	start := n / 3
	end := n - (n / 3)

	// Hindari overlap
	if start >= end {
		start = n / 3
		end = start + 1
	}

	return s[:start] + mask + s[end:]
}

// ~ GenerateQRCode generates a QR code in base64 format from the provided data string
func GenerateQRCode(data string, size int) (string, error) {
	// Generate QR code
	qr, err := qrcode.New(data, qrcode.Medium)
	if err != nil {
		return "", fmt.Errorf("failed to generate QR code: %w", err)
	}

	// Set size (default 256 if not specified)
	if size <= 0 {
		size = 256
	}

	// Encode to PNG
	img := qr.Image(size)

	// Convert to base64
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return "", fmt.Errorf("failed to encode QR code: %w", err)
	}

	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	return base64Str, nil
}
