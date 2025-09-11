package repository

import (
	"errors"

	"sapaUMKM-backend/internal/types/model"

	"gorm.io/gorm"
)

type UsersRepository interface {
	GetAllUsers() ([]model.Users, error)
	GetUserByID(id int) (model.Users, error)
	GetUserByEmail(email string) (model.Users, error)
	CreateUser(user model.Users) (model.Users, error)
	UpdateUser(user model.Users) (model.Users, error)
	DeleteUser(user model.Users) (model.Users, error)

	GetAllRoles() ([]model.Roles, error)
	GetRoleByID(id int) (model.Roles, error)
	IsRoleExist(id int) bool
	IsPermissionExist(id []int) bool
	GetListPermissions() ([]model.Permissions, error)
	GetListRolePermissions() ([]model.RolePermissionsResponse, error)
	DeletePermissionsByRoleID(roleID int) error
	AddRolePermissions(roleID int, permissions []int) error
}

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &usersRepository{db}
}

func (user_repo *usersRepository) GetAllUsers() ([]model.Users, error) {
	var users []model.Users
	err := user_repo.db.Preload("Roles").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (user_repo *usersRepository) GetUserByID(id int) (model.Users, error) {
	var user model.Users
	err := user_repo.db.First(&user, "id = ?", id).Error
	if err != nil {
		return model.Users{}, errors.New("user not found")
	}

	return user, nil
}

func (user_repo *usersRepository) GetUserByEmail(email string) (model.Users, error) {
	var user model.Users
	err := user_repo.db.First(&user, "email = ?", email).Error
	if err != nil {
		return model.Users{}, errors.New("user not found")
	}

	return user, nil
}

func (user_repo *usersRepository) CreateUser(user model.Users) (model.Users, error) {
	err := user_repo.db.Create(&user).Error
	if err != nil {
		return model.Users{}, errors.New("failed to create user")
	}

	return user, nil
}

func (user_repo *usersRepository) UpdateUser(user model.Users) (model.Users, error) {
	err := user_repo.db.Save(&user).Error
	if err != nil {
		return model.Users{}, errors.New("failed to update user")
	}

	return user, nil
}

func (user_repo *usersRepository) DeleteUser(user model.Users) (model.Users, error) {
	err := user_repo.db.Delete(&user).Error
	if err != nil {
		return model.Users{}, errors.New("failed to delete user")
	}

	return user, nil
}

func (user_repo *usersRepository) GetAllRoles() ([]model.Roles, error) {
	var roles []model.Roles
	err := user_repo.db.Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (user_repo *usersRepository) GetRoleByID(id int) (model.Roles, error) {
	var role model.Roles
	err := user_repo.db.First(&role, "id = ?", id).Error
	if err != nil {
		return model.Roles{}, errors.New("role not found")
	}

	return role, nil
}

func (user_repo *usersRepository) IsRoleExist(id int) bool {
	var role model.Roles
	err := user_repo.db.First(&role, "id = ?", id).Error
	return err == nil
}

func (user_repo *usersRepository) IsPermissionExist(ids []int) bool {
	var count int64
	err := user_repo.db.Model(&model.Permissions{}).Where("id IN ?", ids).Count(&count).Error
	if err != nil {
		return false
	}
	return count == int64(len(ids))
}

func (user_repo *usersRepository) GetListPermissions() ([]model.Permissions, error) {
	var permissions []model.Permissions
	err := user_repo.db.Where("parent_id IS NOT NULL").Find(&permissions).Error
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func (user_repo *usersRepository) GetListRolePermissions() ([]model.RolePermissionsResponse, error) {
	var rolePermissions []model.RolePermissionsResponse
	err := user_repo.db.Raw(`
	SELECT 
		roles.id AS role_id,
		roles.name AS role_name,
		jsonb_agg(permissions.name) AS permissions
	FROM role_permissions
	JOIN roles ON role_permissions.role_id = roles.id
	JOIN permissions ON role_permissions.permission_id = permissions.id
	GROUP BY roles.id, roles.name;
	`).Scan(&rolePermissions).Error
	if err != nil {
		return nil, err
	}
	return rolePermissions, nil
}

func (user_repo *usersRepository) DeletePermissionsByRoleID(roleID int) error {
	err := user_repo.db.Where("role_id = ?", roleID).Delete(&model.RolePermissions{}).Error
	if err != nil {
		return errors.New("failed to delete role permissions")
	}
	return nil
}

func (user_repo *usersRepository) AddRolePermissions(roleID int, permissions []int) error {
	var rolePermissions []model.RolePermissions
	for _, permissionID := range permissions {
		rolePermissions = append(rolePermissions, model.RolePermissions{
			RoleID:       roleID,
			PermissionID: permissionID,
		})
	}

	err := user_repo.db.Create(&rolePermissions).Error
	if err != nil {
		return errors.New("failed to add role permissions")
	}
	return nil
}