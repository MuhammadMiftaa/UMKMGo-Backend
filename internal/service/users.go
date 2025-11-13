package service

import (
	"context"
	"errors"
	"time"

	"sapaUMKM-backend/config/redis"
	"sapaUMKM-backend/internal/repository"
	"sapaUMKM-backend/internal/types/dto"
	"sapaUMKM-backend/internal/types/model"
	"sapaUMKM-backend/internal/utils"
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

	GetListPermissions(ctx context.Context) ([]dto.Permissions, error)
	GetListRolePermissions(ctx context.Context) ([]dto.RolePermissionsResponse, error)
	UpdateRolePermissions(ctx context.Context, rolePermissions dto.RolePermissions) error
}

type usersService struct {
	userRepository  repository.UsersRepository
	redisRepository redis.RedisRepository
}

func NewUsersService(usersRepository repository.UsersRepository, redisRepository redis.RedisRepository) UsersService {
	return &usersService{usersRepository, redisRepository}
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

	token, err := utils.GenerateToken(dto.Users{
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
