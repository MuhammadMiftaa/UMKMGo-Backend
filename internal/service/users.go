package service

import (
	"errors"
	"time"

	"sapaUMKM-backend/internal/repository"
	"sapaUMKM-backend/internal/types/dto"
	"sapaUMKM-backend/internal/types/model"
	"sapaUMKM-backend/internal/utils"
)

type UsersService interface {
	Register(user dto.Users) (dto.Users, error)
	Login(user dto.Users) (*string, error)
	VerifyUser(email string) (dto.Users, error)
	GetAllUsers() ([]dto.Users, error)
	GetUserByID(id int) (dto.Users, error)
	GetUserByEmail(email string) (dto.Users, error)
	UpdateUser(id int, userNew dto.Users) (dto.Users, error)
	DeleteUser(id int) (dto.Users, error)

	GetListPermissions() ([]dto.Permissions, error)
	GetListRolePermissions() ([]model.RolePermissionsResponse, error)
	UpdateRolePermissions(rolePermissions dto.RolePermissions) error
}

type usersService struct {
	userRepository repository.UsersRepository
}

func NewUsersService(usersRepository repository.UsersRepository) UsersService {
	return &usersService{usersRepository}
}

func (user_serv *usersService) Register(user dto.Users) (dto.Users, error) {
	// VALIDASI APAKAH NAME, EMAIL, PASSWORD KOSONG
	if user.Name == "" || user.Email == "" || user.Password == "" || user.ConfirmPassword == "" || user.RoleID == nil {
		return dto.Users{}, errors.New("name, email, and password cannot be blank")
	}

	// VALIDASI UNTUK FORMAT EMAIL SUDAH BENAR
	if isValid := utils.EmailValidator(user.Email); !isValid {
		return dto.Users{}, errors.New("please enter a valid email address")
	}

	// MENGECEK APAKAH EMAIL SUDAH DIGUNAKAN
	userExist, err := user_serv.userRepository.GetUserByEmail(user.Email)
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
	if !user_serv.userRepository.IsRoleExist(*user.RoleID) {
		return dto.Users{}, errors.New("role id not found")
	}

	// HASHING PASSWORD MENGGUNAKAN BCRYPT
	hashedPassword, err := utils.PasswordHashing(user.Password)
	if err != nil {
		return dto.Users{}, err
	}
	user.Password = hashedPassword

	newUser, err := user_serv.userRepository.CreateUser(model.Users{
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

func (user_serv *usersService) Login(user dto.Users) (*string, error) {
	// VALIDASI APAKAH EMAIL DAN PASSWORD KOSONG
	if user.Email == "" || user.Password == "" {
		return nil, errors.New("email and password cannot be blank")
	}

	// MENGECEK APAKAH USER SUDAH TERDAFTAR
	userExist, err := user_serv.userRepository.GetUserByEmail(user.Email)
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
	role, err := user_serv.userRepository.GetRoleByID(userExist.RoleID)
	if err != nil {
		return nil, errors.New("role not found")
	}

	// UPDATE LAST LOGIN
	userExist.LastLoginAt = time.Now()
	userExist, err = user_serv.userRepository.UpdateUser(userExist)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(dto.Users{
		ID:       userExist.ID,
		Name:     userExist.Name,
		Email:    userExist.Email,
		RoleID:   &userExist.RoleID,
		RoleName: role.Name,
	})
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (user_serv *usersService) VerifyUser(email string) (dto.Users, error) {
	// MENGAMBIL DATA YANG INGIN DI UPDATE
	user, err := user_serv.userRepository.GetUserByEmail(email)
	if err != nil {
		return dto.Users{}, err
	}

	user.IsActive = true

	userExist, err := user_serv.userRepository.UpdateUser(user)
	if err != nil {
		return dto.Users{}, err
	}

	return dto.Users{
		ID:    userExist.ID,
		Name:  userExist.Name,
		Email: userExist.Email,
	}, nil
}

func (user_serv *usersService) GetAllUsers() ([]dto.Users, error) {
	users, err := user_serv.userRepository.GetAllUsers()
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

func (user_serv *usersService) GetUserByID(id int) (dto.Users, error) {
	user, err := user_serv.userRepository.GetUserByID(id)
	if err != nil {
		return dto.Users{}, err
	}

	return dto.Users{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (user_serv *usersService) GetUserByEmail(email string) (dto.Users, error) {
	user, err := user_serv.userRepository.GetUserByEmail(email)
	if err != nil {
		return dto.Users{}, err
	}

	return dto.Users{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (user_serv *usersService) UpdateUser(id int, userNew dto.Users) (dto.Users, error) {
	// MENGAMBIL DATA YANG INGIN DI UPDATE
	user, err := user_serv.userRepository.GetUserByID(id)
	if err != nil {
		return dto.Users{}, err
	}

	// VALIDASI APAKAH NAME, EMAIL, PASSWORD KOSONG
	if userNew.Name == "" || userNew.Email == "" || userNew.Password == "" || userNew.ConfirmPassword == "" || userNew.RoleID == nil {
		return dto.Users{}, errors.New("name, email, and password cannot be blank")
	}

	// VALIDASI UNTUK FORMAT EMAIL SUDAH BENAR
	if isValid := utils.EmailValidator(userNew.Email); !isValid {
		return dto.Users{}, errors.New("please enter a valid email address")
	}

	// MENGECEK APAKAH EMAIL SUDAH DIGUNAKAN
	userExist, err := user_serv.userRepository.GetUserByEmail(userNew.Email)
	if err == nil && (userExist.Email != "") {
		return dto.Users{}, errors.New("email already exists")
	}

	// VALIDASI PASSWORD SUDAH SESUAI, MIN 8 KARAKTER, MENGANDUNG ALFABET DAN NUMERIK
	hasMinLen, hasLetter, hasDigit := utils.PasswordValidator(userNew.Password)
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
	if userNew.Password != userNew.ConfirmPassword {
		return dto.Users{}, errors.New("password and confirm password do not match")
	}

	// VALIDASI APAKAH ROLE ID ADA
	if !user_serv.userRepository.IsRoleExist(*userNew.RoleID) {
		return dto.Users{}, errors.New("role id not found")
	}

	// HASHING PASSWORD MENGGUNAKAN BCRYPT
	hashedPassword, err := utils.PasswordHashing(userNew.Password)
	if err != nil {
		return dto.Users{}, err
	}
	userNew.Password = hashedPassword

	user.Name = userNew.Name
	user.Email = userNew.Email
	user.Password = userNew.Password
	user.RoleID = *userNew.RoleID

	// UPDATE DATA USER
	userUpdated, err := user_serv.userRepository.UpdateUser(user)
	if err != nil {
		return dto.Users{}, err
	}

	return dto.Users{
		ID:    userUpdated.ID,
		Name:  userUpdated.Name,
		Email: userUpdated.Email,
	}, nil
}

func (user_serv *usersService) DeleteUser(id int) (dto.Users, error) {
	// MENGAMBIL DATA YANG INGIN DI DELETE
	user, err := user_serv.userRepository.GetUserByID(id)
	if err != nil {
		return dto.Users{}, err
	}

	userDeleted, err := user_serv.userRepository.DeleteUser(user)
	if err != nil {
		return dto.Users{}, err
	}

	return dto.Users{
		ID:    userDeleted.ID,
		Name:  userDeleted.Name,
		Email: userDeleted.Email,
	}, nil
}

func (user_serv *usersService) GetListPermissions() ([]dto.Permissions, error) {
	permissions, err := user_serv.userRepository.GetListPermissions()
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

func (user_serv *usersService) GetListRolePermissions() ([]model.RolePermissionsResponse, error) {
	rolePermissions, err := user_serv.userRepository.GetListRolePermissions()
	if err != nil {
		return nil, err
	}

	return rolePermissions, nil
}

func (user_serv *usersService) UpdateRolePermissions(rolePermissions dto.RolePermissions) error {
	// VALIDASI APAKAH ROLE ID ADA
	if !user_serv.userRepository.IsRoleExist(rolePermissions.RoleID) {
		return errors.New("role id not found")
	}

	// VALIDASI APAKAH PERMISSION VALID
	if isAllExists := user_serv.userRepository.IsPermissionExist(rolePermissions.Permissions); !isAllExists {
		return errors.New("one or more permissions are invalid")
	}

	// DELETE PERMISSION YANG ADA
	err := user_serv.userRepository.DeletePermissionsByRoleID(rolePermissions.RoleID)
	if err != nil {
		return err
	}

	// ADD PERMISSION YANG BARU
	err = user_serv.userRepository.AddRolePermissions(rolePermissions.RoleID, rolePermissions.Permissions)
	if err != nil {
		return err
	}

	return nil
}
