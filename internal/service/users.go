package service

import (
	"errors"

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
	GetUserByID(id string) (dto.Users, error)
	GetUserByEmail(email string) (dto.Users, error)
	UpdateUser(id string, userNew dto.Users) (dto.Users, error)
	DeleteUser(id string) (dto.Users, error)
}

type usersService struct {
	userRepository repository.UsersRepository
}

func NewUsersService(usersRepository repository.UsersRepository) UsersService {
	return &usersService{usersRepository}
}

func (user_serv *usersService) Register(user dto.Users) (dto.Users, error) {
	// VALIDASI APAKAH NAME, EMAIL, PASSWORD KOSONG
	if user.Name == "" || user.Email == "" || user.Password == "" {
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
		RoleID:   user.RoleID,
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

	token, err := utils.GenerateToken(userExist.ID, userExist.Name, userExist.Email)
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
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return usersDTO, nil
}

func (user_serv *usersService) GetUserByID(id string) (dto.Users, error) {
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

func (user_serv *usersService) UpdateUser(id string, userNew dto.Users) (dto.Users, error) {
	// MENGAMBIL DATA YANG INGIN DI UPDATE
	user, err := user_serv.userRepository.GetUserByID(id)
	if err != nil {
		return dto.Users{}, err
	}

	// VALIDASI APAKAH FULLNAME & EMAIL KOSONG
	if userNew.Name == "" && userNew.Email == "" {
		return dto.Users{}, errors.New("fullname and email cannot be blank")
	}

	// VALIDASI APAKAH FULLNAME / EMAIL SUDAH DI INPUT
	if userNew.Name != "" {
		user.Name = userNew.Name
	}

	if userNew.Email != "" {
		// VALIDASI UNTUK FORMAT EMAIL SUDAH BENAR
		if isValid := utils.EmailValidator(userNew.Email); !isValid {
			return dto.Users{}, errors.New("please enter a valid email address")
		}
		// MENGECEK APAKAH EMAIL SUDAH DIGUNAKAN
		existingUser, err := user_serv.userRepository.GetUserByEmail(userNew.Email)
		if err == nil && existingUser.ID != user.ID {
			return dto.Users{}, errors.New("email already in use by another user")
		}
		user.Email = userNew.Email
	}

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

func (user_serv *usersService) DeleteUser(id string) (dto.Users, error) {
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
