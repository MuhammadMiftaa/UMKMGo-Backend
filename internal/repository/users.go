package repository

import (
	"context"
	"encoding/json"
	"errors"

	"sapaUMKM-backend/internal/types/model"

	"gorm.io/gorm"
)

type UsersRepository interface {
	GetAllUsers(ctx context.Context) ([]model.Users, error)
	GetUserByID(ctx context.Context, id int) (model.Users, error)
	GetUserByEmail(ctx context.Context, email string) (model.Users, error)
	CreateUser(ctx context.Context, user model.Users) (model.Users, error)
	UpdateUser(ctx context.Context, user model.Users) (model.Users, error)
	DeleteUser(ctx context.Context, user model.Users) (model.Users, error)

	GetAllRoles(ctx context.Context) ([]model.Roles, error)
	GetRoleByID(ctx context.Context, id int) (model.Roles, error)
	IsRoleExist(ctx context.Context, id int) bool
	IsPermissionExist(ctx context.Context, id []string) ([]int, bool)
	GetListPermissions(ctx context.Context) ([]model.Permissions, error)
	GetListPermissionsByRoleID(ctx context.Context, roleID int) ([]string, error)
	GetListRolePermissions(ctx context.Context) ([]model.RolePermissionsResponse, error)
	DeletePermissionsByRoleID(ctx context.Context, roleID int) error
	AddRolePermissions(ctx context.Context, roleID int, permissions []int) error
}

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &usersRepository{db}
}

func (user_repo *usersRepository) GetAllUsers(ctx context.Context) ([]model.Users, error) {
	var users []model.Users
	err := user_repo.db.WithContext(ctx).Preload("Roles").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (user_repo *usersRepository) GetUserByID(ctx context.Context, id int) (model.Users, error) {
	var user model.Users
	err := user_repo.db.WithContext(ctx).First(&user, "id = ?", id).Error
	if err != nil {
		return model.Users{}, errors.New("user not found")
	}

	return user, nil
}

func (user_repo *usersRepository) GetUserByEmail(ctx context.Context, email string) (model.Users, error) {
	var user model.Users
	err := user_repo.db.WithContext(ctx).First(&user, "email = ?", email).Error
	if err != nil {
		return model.Users{}, errors.New("user not found")
	}

	return user, nil
}

func (user_repo *usersRepository) CreateUser(ctx context.Context, user model.Users) (model.Users, error) {
	err := user_repo.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return model.Users{}, errors.New("failed to create user")
	}

	return user, nil
}

func (user_repo *usersRepository) UpdateUser(ctx context.Context, user model.Users) (model.Users, error) {
	err := user_repo.db.WithContext(ctx).Save(&user).Error
	if err != nil {
		return model.Users{}, errors.New("failed to update user")
	}

	return user, nil
}

func (user_repo *usersRepository) DeleteUser(ctx context.Context, user model.Users) (model.Users, error) {
	err := user_repo.db.WithContext(ctx).Delete(&user).Error
	if err != nil {
		return model.Users{}, errors.New("failed to delete user")
	}

	return user, nil
}

func (user_repo *usersRepository) GetAllRoles(ctx context.Context) ([]model.Roles, error) {
	var roles []model.Roles
	err := user_repo.db.WithContext(ctx).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (user_repo *usersRepository) GetRoleByID(ctx context.Context, id int) (model.Roles, error) {
	var role model.Roles
	err := user_repo.db.WithContext(ctx).First(&role, "id = ?", id).Error
	if err != nil {
		return model.Roles{}, errors.New("role not found")
	}

	return role, nil
}

func (user_repo *usersRepository) IsRoleExist(ctx context.Context, id int) bool {
	var role model.Roles
	err := user_repo.db.WithContext(ctx).First(&role, "id = ?", id).Error
	return err == nil
}

func (user_repo *usersRepository) IsPermissionExist(ctx context.Context, ids []string) ([]int, bool) {
	var permissionIDs []int
	err := user_repo.db.WithContext(ctx).Model(&model.Permissions{}).Where("code IN ?", ids).Pluck("id", &permissionIDs).Error
	if err != nil {
		return nil, false
	}
	return permissionIDs, len(permissionIDs) == len(ids)
}

func (user_repo *usersRepository) GetListPermissions(ctx context.Context) ([]model.Permissions, error) {
	var permissions []model.Permissions
	err := user_repo.db.WithContext(ctx).Where("parent_id IS NOT NULL").Find(&permissions).Error
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func (user_repo *usersRepository) GetListRolePermissions(ctx context.Context) ([]model.RolePermissionsResponse, error) {
	var rolePermissions []model.RolePermissionsResponse
	err := user_repo.db.WithContext(ctx).Raw(`
	SELECT 
		roles.id AS role_id,
		roles.name AS role_name,
		jsonb_agg(permissions.code) AS permissions
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

func (user_repo *usersRepository) GetListPermissionsByRoleID(ctx context.Context, roleID int) ([]string, error) {
	var rolePermissionsRaw string
	var rolePermissions []string
	err := user_repo.db.WithContext(ctx).Raw(`
	SELECT jsonb_agg(permissions.code) AS permissions
	FROM role_permissions
	JOIN roles ON role_permissions.role_id = roles.id
	JOIN permissions ON role_permissions.permission_id = permissions.id
	WHERE roles.id = ?
	`, roleID).Scan(&rolePermissionsRaw).Error

	if err := json.Unmarshal([]byte(rolePermissionsRaw), &rolePermissions); err != nil {
		return nil, errors.New("failed to parse role permissions")
	}
	
	if err != nil {
		return nil, err
	}
	return rolePermissions, nil
}

func (user_repo *usersRepository) DeletePermissionsByRoleID(ctx context.Context, roleID int) error {
	err := user_repo.db.WithContext(ctx).Where("role_id = ?", roleID).Unscoped().Delete(&model.RolePermissions{}).Error
	if err != nil {
		return errors.New("failed to delete role permissions")
	}
	return nil
}

func (user_repo *usersRepository) AddRolePermissions(ctx context.Context, roleID int, permissions []int) error {
	var rolePermissions []model.RolePermissions
	for _, permissionID := range permissions {
		rolePermissions = append(rolePermissions, model.RolePermissions{
			RoleID:       roleID,
			PermissionID: permissionID,
		})
	}

	err := user_repo.db.WithContext(ctx).Omit("CreatedAt", "UpdatedAt", "DeletedAt").Create(&rolePermissions).Error
	if err != nil {
		return errors.New("failed to add role permissions")
	}
	return nil
}
