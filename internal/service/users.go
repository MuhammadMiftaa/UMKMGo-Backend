package service

import (
	"context"
	"errors"
	"time"

	"UMKMGo-backend/config/env"
	"UMKMGo-backend/config/redis"
	"UMKMGo-backend/config/storage"
	"UMKMGo-backend/config/vault"
	"UMKMGo-backend/internal/repository"
	"UMKMGo-backend/internal/types/dto"
	"UMKMGo-backend/internal/types/model"
	"UMKMGo-backend/internal/utils"
	"UMKMGo-backend/internal/utils/constant"

	otp "github.com/MuhammadMiftaa/Internal/golang/otp-whatsapp"
)

type UsersService interface {
	Register(ctx context.Context, user dto.Users) (dto.Users, error)
	Login(ctx context.Context, user dto.Users) (*string, error)
	UpdateProfile(ctx context.Context, id int, userNew dto.Users) (dto.Users, error)
	SetOTP(ctx context.Context, email string, otp string, expiration time.Duration) error
	ValidateOTP(ctx context.Context, email string, otp string) (bool, error)
	VerifyUser(ctx context.Context, email string) (dto.Users, error)
	GetAllUsers(ctx context.Context) ([]dto.Users, error)
	GetUserByID(ctx context.Context, id int) (dto.Users, error)
	GetUserByEmail(ctx context.Context, email string) (dto.Users, error)
	UpdateUser(ctx context.Context, id int, userNew dto.Users) (dto.Users, error)
	DeleteUser(ctx context.Context, id int) (dto.Users, error)

	MetaCityAndProvince(ctx context.Context) ([]dto.MetaCityAndProvince, error)
	RegisterMobile(ctx context.Context, email, phone string) error
	VerifyOTP(ctx context.Context, phone, code string) (*string, error)
	RegisterMobileProfile(ctx context.Context, user dto.UMKMMobile, tempToken string) (*string, error)
	LoginMobile(ctx context.Context, user dto.UMKMMobile) (*string, error)
	ForgotPassword(ctx context.Context, phone string) error
	ResetPassword(ctx context.Context, user dto.ResetPasswordMobile, tempToken string) error

	GetListPermissions(ctx context.Context) ([]dto.Permissions, error)
	GetListRolePermissions(ctx context.Context) ([]dto.RolePermissionsResponse, error)
	UpdateRolePermissions(ctx context.Context, rolePermissions dto.RolePermissions) error
}

type usersService struct {
	userRepository  repository.UsersRepository
	otpRepository   repository.OTPRepository
	redisRepository redis.RedisRepository
	minio           *storage.MinIOManager
}

func NewUsersService(usersRepository repository.UsersRepository, otpRepository repository.OTPRepository, redisRepository redis.RedisRepository, minio *storage.MinIOManager) UsersService {
	return &usersService{usersRepository, otpRepository, redisRepository, minio}
}

func (user_serv *usersService) Register(ctx context.Context, user dto.Users) (dto.Users, error) {
	// VALIDASI APAKAH NAME, EMAIL, PASSWORD KOSONG
	if user.Name == "" || user.Email == "" || user.Password == "" || user.ConfirmPassword == "" || user.RoleID == nil {
		return dto.Users{}, errors.New("name, email, and password cannot be blank")
	}

	// VALIDASI UNTUK FORMAT EMAIL SUDAH BENAR
	if isValid := utils.EmailValidator(user.Email); !isValid {
		return dto.Users{}, errors.New("please enter a valid email address")
	}

	// MENGECEK APAKAH EMAIL SUDAH DIGUNAKAN
	userExist, err := user_serv.userRepository.GetUserByEmail(ctx, user.Email)
	if err == nil && (userExist.Email != "") {
		return dto.Users{}, errors.New("email already exists")
	}

	// VALIDASI PASSWORD SUDAH SESUAI, MIN 8 KARAKTER, MENGANDUNG ALFABET DAN NUMERIK
	hasMinLen, hasLetter, hasDigit := utils.PasswordValidator(user.Password)
	if !hasMinLen {
		return dto.Users{}, errors.New("password must be at least 8 characters long")
	}
	if !hasLetter {
		return dto.Users{}, errors.New("password must contain at least one letter")
	}
	if !hasDigit {
		return dto.Users{}, errors.New("password must contain at least one number")
	}

	// VALIDASI PASSWORD DAN CONFIRM PASSWORD SUDAH SESUAI
	if user.Password != user.ConfirmPassword {
		return dto.Users{}, errors.New("password and confirm password do not match")
	}

	// VALIDASI APAKAH ROLE ID ADA
	if !user_serv.userRepository.IsRoleExist(ctx, *user.RoleID) {
		return dto.Users{}, errors.New("role id not found")
	}

	// HASHING PASSWORD MENGGUNAKAN BCRYPT
	hashedPassword, err := utils.PasswordHashing(user.Password)
	if err != nil {
		return dto.Users{}, err
	}
	user.Password = hashedPassword

	newUser, err := user_serv.userRepository.CreateUser(ctx, model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		RoleID:   *user.RoleID,
		IsActive: true,
	})
	if err != nil {
		return dto.Users{}, err
	}

	return dto.Users{
		ID:    newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email,
	}, nil
}

func (user_serv *usersService) Login(ctx context.Context, user dto.Users) (*string, error) {
	// VALIDASI APAKAH EMAIL DAN PASSWORD KOSONG
	if user.Email == "" || user.Password == "" {
		return nil, errors.New("email and password cannot be blank")
	}

	// MENGECEK APAKAH USER SUDAH TERDAFTAR
	userExist, err := user_serv.userRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// VALIDASI APAKAH PASSWORD SUDAH SESUAI
	if !utils.ComparePass(userExist.Password, user.Password) {
		return nil, errors.New("password is incorrect")
	}

	// VALIDASI APAKAH USER SUDAH AKTIF
	if !userExist.IsActive {
		return nil, errors.New("user is not active")
	}

	// AMBIL ROLE NAME
	role, err := user_serv.userRepository.GetRoleByID(ctx, userExist.RoleID)
	if err != nil {
		return nil, errors.New("role not found")
	}

	// UPDATE LAST LOGIN
	userExist.LastLoginAt = time.Now()
	userExist, err = user_serv.userRepository.UpdateUser(ctx, userExist)
	if err != nil {
		return nil, err
	}

	// AMBIL LIST ROLE PERMISSIONS
	rolePermissions, err := user_serv.userRepository.GetListPermissionsByRoleID(ctx, userExist.RoleID)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateWebToken(dto.Users{
		ID:          userExist.ID,
		Name:        userExist.Name,
		Email:       userExist.Email,
		RoleID:      &userExist.RoleID,
		RoleName:    role.Name,
		Permissions: rolePermissions,
	})
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (user_serv *usersService) UpdateProfile(ctx context.Context, id int, userNew dto.Users) (dto.Users, error) {
	// MENGAMBIL DATA YANG INGIN DI UPDATE
	user, err := user_serv.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return dto.Users{}, err
	}

	// VALIDASI APAKAH NAME, EMAIL KOSONG
	if userNew.Name == "" || userNew.Email == "" {
		return dto.Users{}, errors.New("name and email cannot be blank")
	}

	// VALIDASI UNTUK FORMAT EMAIL SUDAH BENAR
	if isValid := utils.EmailValidator(userNew.Email); !isValid {
		return dto.Users{}, errors.New("please enter a valid email address")
	}

	// MENGECEK APAKAH EMAIL SUDAH DIGUNAKAN OLEH USER LAIN
	userExist, err := user_serv.userRepository.GetUserByEmail(ctx, userNew.Email)
	if err == nil && (userExist.Email != "") && userExist.ID != id {
		return dto.Users{}, errors.New("email already exists")
	}

	// UPDATE ONLY NAME AND EMAIL
	user.Name = userNew.Name
	user.Email = userNew.Email

	// UPDATE DATA USER
	userUpdated, err := user_serv.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return dto.Users{}, err
	}

	return dto.Users{
		ID:    userUpdated.ID,
		Name:  userUpdated.Name,
		Email: userUpdated.Email,
	}, nil
}

func (user_serv *usersService) SetOTP(ctx context.Context, email string, otp string, expiration time.Duration) error {
	// VALIDASI APAKAH EMAIL KOSONG
	if email == "" {
		return errors.New("email cannot be blank")
	}

	// VALIDASI UNTUK FORMAT EMAIL SUDAH BENAR
	if isValid := utils.EmailValidator(email); !isValid {
		return errors.New("please enter a valid email address")
	}

	// MENGECEK APAKAH USER SUDAH TERDAFTAR
	userExist, err := user_serv.userRepository.GetUserByEmail(ctx, email)
	if err != nil || userExist.Email == "" {
		return errors.New("user not found")
	}

	// SET OTP KE REDIS
	err = user_serv.redisRepository.Set(ctx, email, otp, expiration)
	if err != nil {
		return err
	}

	return nil
}

func (user_serv *usersService) ValidateOTP(ctx context.Context, email string, otp string) (bool, error) {
	// VALIDASI APAKAH EMAIL KOSONG
	if email == "" {
		return false, errors.New("email cannot be blank")
	}

	// VALIDASI UNTUK FORMAT EMAIL SUDAH BENAR
	if isValid := utils.EmailValidator(email); !isValid {
		return false, errors.New("please enter a valid email address")
	}

	// MENGECEK APAKAH USER SUDAH TERDAFTAR
	userExist, err := user_serv.userRepository.GetUserByEmail(ctx, email)
	if err != nil || userExist.Email == "" {
		return false, errors.New("user not found")
	}

	// VALIDASI OTP DARI REDIS
	validOTP, err := user_serv.redisRepository.Get(ctx, email)
	if err != nil || validOTP != otp {
		return false, errors.New("invalid or expired OTP")
	}

	return true, nil
}

func (user_serv *usersService) VerifyUser(ctx context.Context, email string) (dto.Users, error) {
	// MENGAMBIL DATA YANG INGIN DI UPDATE
	user, err := user_serv.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return dto.Users{}, err
	}

	user.IsActive = true

	userExist, err := user_serv.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return dto.Users{}, err
	}

	return dto.Users{
		ID:    userExist.ID,
		Name:  userExist.Name,
		Email: userExist.Email,
	}, nil
}

func (user_serv *usersService) GetAllUsers(ctx context.Context) ([]dto.Users, error) {
	users, err := user_serv.userRepository.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	var usersDTO []dto.Users
	for _, user := range users {
		usersDTO = append(usersDTO, dto.Users{
			ID:          user.ID,
			Name:        user.Name,
			Email:       user.Email,
			RoleName:    user.Roles.Name,
			IsActive:    user.IsActive,
			LastLoginAt: user.LastLoginAt.Format("2006-01-02 15:04:05"),
			CreatedAt:   user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return usersDTO, nil
}

func (user_serv *usersService) GetUserByID(ctx context.Context, id int) (dto.Users, error) {
	user, err := user_serv.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return dto.Users{}, err
	}

	return dto.Users{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (user_serv *usersService) GetUserByEmail(ctx context.Context, email string) (dto.Users, error) {
	user, err := user_serv.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return dto.Users{}, err
	}

	return dto.Users{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (user_serv *usersService) UpdateUser(ctx context.Context, id int, userNew dto.Users) (dto.Users, error) {
	// MENGAMBIL DATA YANG INGIN DI UPDATE
	user, err := user_serv.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return dto.Users{}, err
	}

	// VALIDASI APAKAH NAME, EMAIL, ROLE KOSONG
	if userNew.Name == "" || userNew.Email == "" || userNew.RoleID == nil {
		return dto.Users{}, errors.New("name, email, and role cannot be blank")
	}

	// VALIDASI UNTUK FORMAT EMAIL SUDAH BENAR
	if isValid := utils.EmailValidator(userNew.Email); !isValid {
		return dto.Users{}, errors.New("please enter a valid email address")
	}

	// MENGECEK APAKAH EMAIL SUDAH DIGUNAKAN
	userExist, err := user_serv.userRepository.GetUserByEmail(ctx, userNew.Email)
	if err == nil && (userExist.Email != "") {
		return dto.Users{}, errors.New("email already exists")
	}

	// VALIDASI APAKAH ROLE ID ADA
	if !user_serv.userRepository.IsRoleExist(ctx, *userNew.RoleID) {
		return dto.Users{}, errors.New("role id not found")
	}

	user.Name = userNew.Name
	user.Email = userNew.Email
	user.RoleID = *userNew.RoleID

	// UPDATE DATA USER
	userUpdated, err := user_serv.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return dto.Users{}, err
	}

	return dto.Users{
		ID:    userUpdated.ID,
		Name:  userUpdated.Name,
		Email: userUpdated.Email,
	}, nil
}

func (user_serv *usersService) DeleteUser(ctx context.Context, id int) (dto.Users, error) {
	// MENGAMBIL DATA YANG INGIN DI DELETE
	user, err := user_serv.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return dto.Users{}, err
	}

	userDeleted, err := user_serv.userRepository.DeleteUser(ctx, user)
	if err != nil {
		return dto.Users{}, err
	}

	return dto.Users{
		ID:    userDeleted.ID,
		Name:  userDeleted.Name,
		Email: userDeleted.Email,
	}, nil
}

func (user_serv *usersService) GetListPermissions(ctx context.Context) ([]dto.Permissions, error) {
	permissions, err := user_serv.userRepository.GetListPermissions(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]dto.Permissions, len(permissions))
	for i, p := range permissions {
		result[i] = dto.Permissions{
			ID:          p.ID,
			Name:        p.Name,
			Code:        p.Code,
			Description: p.Description,
		}
	}

	return result, nil
}

func (user_serv *usersService) GetListRolePermissions(ctx context.Context) ([]dto.RolePermissionsResponse, error) {
	rolePermissions, err := user_serv.userRepository.GetListRolePermissions(ctx)
	if err != nil {
		return nil, err
	}

	return rolePermissions, nil
}

func (user_serv *usersService) UpdateRolePermissions(ctx context.Context, rolePermissions dto.RolePermissions) error {
	// VALIDASI APAKAH ROLE ID ADA
	if !user_serv.userRepository.IsRoleExist(ctx, rolePermissions.RoleID) {
		return errors.New("role id not found")
	}

	// VALIDASI APAKAH PERMISSION VALID
	permissionIDs, isAllExists := user_serv.userRepository.IsPermissionExist(ctx, rolePermissions.Permissions)
	if !isAllExists {
		return errors.New("one or more permissions are invalid")
	}

	// DELETE PERMISSION YANG ADA
	err := user_serv.userRepository.DeletePermissionsByRoleID(ctx, rolePermissions.RoleID)
	if err != nil {
		return err
	}

	// ADD PERMISSION YANG BARU
	err = user_serv.userRepository.AddRolePermissions(ctx, rolePermissions.RoleID, permissionIDs)
	if err != nil {
		return err
	}

	return nil
}

// ====================== Mobile Auth =================================

func (user_serv *usersService) MetaCityAndProvince(ctx context.Context) ([]dto.MetaCityAndProvince, error) {
	var result []dto.MetaCityAndProvince
	provinces, err := user_serv.userRepository.GetProvinces(ctx)
	if err != nil {
		return nil, err
	}

	cities, err := user_serv.userRepository.GetCities(ctx)
	if err != nil {
		return nil, err
	}

	result = append(result, dto.MetaCityAndProvince{
		Provinces: provinces,
		Cities:    cities,
	})

	return result, nil
}

func (user_serv *usersService) RegisterMobile(ctx context.Context, email, phone string) error {
	// VALIDASI APAKAH EMAIL DAN PHONE KOSONG
	if email == "" || phone == "" {
		return errors.New("email and phone cannot be blank")
	}

	// VALIDASI UNTUK FORMAT EMAIL SUDAH BENAR
	if isValid := utils.EmailValidator(email); !isValid {
		return errors.New("please enter a valid email address")
	}

	// VALIDASI NOMOR TELEPON
	validPhone, err := utils.NormalizePhone(phone)
	if err != nil {
		return errors.New("please enter a valid phone number")
	}

	// MENGECEK APAKAH USER SUDAH TERDAFTAR DAN NOMOr TELEPON SUDAH DIGUNAKAN
	userExist, err := user_serv.userRepository.GetUMKMByPhone(ctx, validPhone)
	if err == nil {
		return errors.New("phone already exists")
	}

	// MENGECEK APAKAH EMAIL SUDAH DIGUNAKAN
	if userExist.User.Email == email {
		return errors.New("email already exists")
	}

	otpCode := utils.GenerateOTP()
	OTP := model.OTP{
		PhoneNumber: validPhone,
		Email:       email,
		OTPCode:     otpCode,
		ExpiresAt:   time.Now().Add(5 * time.Minute),
		Status:      constant.OTPStatusActive,
	}

	if err := user_serv.otpRepository.CreateOTP(ctx, OTP); err != nil {
		return errors.New("failed to create OTP")
	}

	vendor, err := otp.InitVendor(otp.VENDOR_FONNTE, env.Cfg.Fonnte.Token, "", "")
	if err != nil {
		return errors.New("failed to initialize OTP vendor")
	}

	if _, err := otp.SendOTP(vendor, validPhone, otpCode); err != nil {
		return errors.New("failed to send OTP")
	}

	return nil
}

func (user_serv *usersService) VerifyOTP(ctx context.Context, phone, code string) (*string, error) {
	phone, err := utils.NormalizePhone(phone)
	if err != nil {
		return nil, errors.New("please enter a valid phone number")
	}

	OTP, err := user_serv.otpRepository.GetOTPByPhone(ctx, phone)
	if err != nil {
		return nil, errors.New("failed to get OTP")
	}

	if OTP == nil || OTP.ExpiresAt.Before(time.Now()) || OTP.Status != constant.OTPStatusActive {
		return nil, errors.New("OTP expired or not found")
	}

	if OTP.OTPCode != code {
		return nil, errors.New("invalid OTP")
	}

	tempToken := utils.RandomString(10)
	OTP.TempToken = &tempToken

	if err := user_serv.otpRepository.UpdateOTP(ctx, *OTP); err != nil {
		return nil, errors.New("failed to update OTP")
	}

	return &tempToken, nil
}

func (user_serv *usersService) RegisterMobileProfile(ctx context.Context, user dto.UMKMMobile, tempToken string) (*string, error) {
	OTP, err := user_serv.otpRepository.GetOTPByTempToken(ctx, tempToken)
	if err != nil {
		return nil, errors.New("failed to get OTP")
	}

	if OTP == nil || OTP.Status != constant.OTPStatusActive {
		return nil, errors.New("OTP expired or not found")
	}

	// INPUT VALIDATION
	if user.Fullname == "" {
		return nil, errors.New("fullname cannot be blank")
	}
	if user.BusinessName == "" {
		return nil, errors.New("business name cannot be blank")
	}
	if user.NIK == "" {
		return nil, errors.New("NIK cannot be blank")
	}
	if user.BirthDate == "" {
		return nil, errors.New("birth date cannot be blank")
	}
	if user.Gender == "" {
		return nil, errors.New("gender cannot be blank")
	}
	if user.Address == "" {
		return nil, errors.New("address cannot be blank")
	}
	if user.ProvinceID == 0 {
		return nil, errors.New("province cannot be blank")
	}
	if user.CityID == 0 {
		return nil, errors.New("city cannot be blank")
	}
	if user.District == "" {
		return nil, errors.New("district cannot be blank")
	}
	if user.PostalCode == "" {
		return nil, errors.New("postal code cannot be blank")
	}
	if user.KartuType == "" {
		return nil, errors.New("kartu type cannot be blank")
	}
	if user.KartuNumber == "" {
		return nil, errors.New("kartu number cannot be blank")
	}
	if hasLetter, hasDigit, hasMinLen := utils.PasswordValidator(user.Password); !hasLetter || !hasDigit || !hasMinLen {
		return nil, errors.New("password must contain at least 8 characters, 1 letter and 1 number")
	}

	// VALIDASI NOMOR TELEPON
	validPhone, err := utils.NormalizePhone(OTP.PhoneNumber)
	if err != nil {
		return nil, errors.New("please enter a valid phone number")
	}

	// ! VALIDASI UNTUK FORMAT NIK
	// if err := utils.NIKValidator(user.NIK); err != nil {
	// 	return nil, errors.New("please enter a valid NIK")
	// }

	// VALIDASI TANGGAL LAHIR
	birthDate, err := time.Parse("2006-01-02", user.BirthDate)
	if err != nil {
		return nil, errors.New("please enter a valid birth date in format YYYY-MM-DD")
	}

	// HASHING PASSWORD MENGGUNAKAN BCRYPT
	hashedPassword, err := utils.PasswordHashing(user.Password)
	if err != nil {
		return nil, err
	}

	// AMBIL ROLE UMKM DARI DATABASE
	role, err := user_serv.userRepository.GetRoleByName(ctx, constant.RoleUMKM)
	if err != nil {
		return nil, errors.New("role UMKM not found")
	}

	// PROSES ENKRIPSI NIK MENGGUNAKAN VAULT DENGAN ALGORITMA AES256-GCM96
	ciphertextNIK, err := vault.EncryptTransit(ctx, env.Cfg.Vault.TransitPath, env.Cfg.Vault.NIKEncryptionKey, []byte(user.NIK))
	if err != nil {
		return nil, errors.New("failed to encrypt NIK - " + err.Error())
	}

	// PROSES ENKRIPSI KARTU NUMBER MENGGUNAKAN VAULT DENGAN ALGORITMA AES256-GCM96
	ciphertextKartuNumber, err := vault.EncryptTransit(ctx, env.Cfg.Vault.TransitPath, env.Cfg.Vault.KartuEncryptionKey, []byte(user.KartuNumber))
	if err != nil {
		return nil, errors.New("failed to encrypt Kartu Number - " + err.Error())
	}

	// Generate QR Code from Kartu Number
	qrCodeBase64, err := utils.GenerateQRCode(string(user.KartuNumber), 256)
	if err != nil {
		return nil, errors.New("failed to generate QR code: " + err.Error())
	}

	// Upload QR Code to MinIO
	qrCode, err := user_serv.minio.UploadFile(ctx, storage.UploadRequest{
		Base64Data: qrCodeBase64,
		BucketName: storage.UMKMBucket,
		Prefix:     utils.GenerateFileName(user.Fullname, "qrcode_"),
	})
	if err != nil {
		return nil, errors.New("failed to upload QR code: " + err.Error())
	}

	res, err := user_serv.userRepository.CreateUMKM(ctx,
		model.UMKM{
			BusinessName: user.BusinessName,
			NIK:          ciphertextNIK,
			Gender:       user.Gender,
			BirthDate:    birthDate,
			Phone:        validPhone,
			Address:      user.Address,
			ProvinceID:   user.ProvinceID,
			CityID:       user.CityID,
			District:     user.District,
			PostalCode:   user.PostalCode,
			KartuType:    user.KartuType,
			KartuNumber:  ciphertextKartuNumber,
			QRCode:       qrCode.URL,
		},
		model.User{
			Name:        user.Fullname,
			Email:       OTP.Email,
			Password:    hashedPassword,
			RoleID:      role.ID,
			IsActive:    true,
			LastLoginAt: time.Now(),
		})
	if err != nil {
		return nil, err
	}

	OTP.Status = constant.OTPStatusUsed
	if err := user_serv.otpRepository.UpdateOTP(ctx, *OTP); err != nil {
		return nil, errors.New("failed to update OTP status")
	}

	token, err := utils.GenerateMobileToken(res)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (user_serv *usersService) LoginMobile(ctx context.Context, user dto.UMKMMobile) (*string, error) {
	// VALIDASI APAKAH PHONE DAN PASSWORD KOSONG
	if user.Phone == "" || user.Password == "" {
		return nil, errors.New("phone and password cannot be blank")
	}

	// MENGGUNAKAN NORMALIZE PHONE UNTUK MEMASTIKAN FORMAT YANG BENAR
	validPhone, err := utils.NormalizePhone(user.Phone)
	if err != nil {
		return nil, errors.New("please enter a valid phone number")
	}

	// MENGECEK APAKAH USER SUDAH TERDAFTAR
	userExist, err := user_serv.userRepository.GetUMKMByPhone(ctx, validPhone)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// VALIDASI APAKAH PASSWORD SUDAH SESUAI
	if !utils.ComparePass(userExist.User.Password, user.Password) {
		return nil, errors.New("password is incorrect")
	}

	token, err := utils.GenerateMobileToken(dto.UMKMMobile{
		ID:           userExist.ID,
		Fullname:     userExist.User.Name,
		BusinessName: userExist.BusinessName,
		Email:        userExist.User.Email,
		Phone:        userExist.Phone,
		KartuType:    userExist.KartuType,
	})
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (user_serv *usersService) ForgotPassword(ctx context.Context, phone string) error {
	// VALIDASI APAKAH PHONE KOSONG
	if phone == "" {
		return errors.New("phone cannot be blank")
	}

	// VALIDASI NOMOR TELEPON
	validPhone, err := utils.NormalizePhone(phone)
	if err != nil {
		return errors.New("please enter a valid phone number")
	}

	// MENGECEK APAKAH USER SUDAH TERDAFTAR
	userExist, err := user_serv.userRepository.GetUMKMByPhone(ctx, validPhone)
	if err != nil {
		return errors.New("user not found")
	}

	otpCode := utils.GenerateOTP()
	OTP := model.OTP{
		PhoneNumber: validPhone,
		Email:       userExist.User.Email,
		OTPCode:     otpCode,
		ExpiresAt:   time.Now().Add(5 * time.Minute),
		Status:      constant.OTPStatusActive,
	}

	if err := user_serv.otpRepository.CreateOTP(ctx, OTP); err != nil {
		return errors.New("failed to create OTP")
	}

	vendor, err := otp.InitVendor(otp.VENDOR_FONNTE, env.Cfg.Fonnte.Token, "", "")
	if err != nil {
		return errors.New("failed to initialize OTP vendor")
	}

	if _, err := otp.SendOTP(vendor, validPhone, otpCode); err != nil {
		return errors.New("failed to send OTP")
	}

	return nil
}

func (user_serv *usersService) ResetPassword(ctx context.Context, user dto.ResetPasswordMobile, tempToken string) error {
	// VALIDASI APAKAH PHONE, PASSWORD, DAN CONFIRM PASSWORD KOSONG
	if user.Password == "" || user.ConfirmPassword == "" {
		return errors.New("password cannot be blank")
	}

	// MENGAMBIL OTP DARI TEMP TOKEN
	OTP, err := user_serv.otpRepository.GetOTPByTempToken(ctx, tempToken)
	if err != nil {
		return errors.New("failed to get OTP")
	}

	if OTP == nil || OTP.Status != constant.OTPStatusActive {
		return errors.New("OTP expired or not found")
	}

	// MENGECEK APAKAH USER SUDAH TERDAFTAR
	userExist, err := user_serv.userRepository.GetUserByEmail(ctx, OTP.Email)
	if err != nil {
		return errors.New("user not found")
	}

	// VALIDASI PASSWORD SUDAH SESUAI, MIN 8 KARAKTER, MENGANDUNG ALFABET DAN NUMERIK
	if hasLetter, hasDigit, hasMinLen := utils.PasswordValidator(user.Password); !hasLetter || !hasDigit || !hasMinLen {
		return errors.New("password must contain at least 8 characters, 1 letter and 1 number")
	}

	// VALIDASI PASSWORD DAN CONFIRM PASSWORD SUDAH SESUAI
	if user.Password != user.ConfirmPassword {
		return errors.New("password and confirm password do not match")
	}

	// HASHING PASSWORD MENGGUNAKAN BCRYPT
	hashedPassword, err := utils.PasswordHashing(user.Password)
	if err != nil {
		return err
	}

	userExist.Password = hashedPassword

	// UPDATE PASSWORD USER
	if _, err := user_serv.userRepository.UpdateUser(ctx, userExist); err != nil {
		return errors.New("failed to update password")
	}

	// UPDATE STATUS OTP MENJADI USED
	OTP.Status = constant.OTPStatusUsed
	if err := user_serv.otpRepository.UpdateOTP(ctx, *OTP); err != nil {
		return errors.New("failed to update OTP status")
	}

	return nil
}
