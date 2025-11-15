package repository

import (
	"context"
	"encoding/json"
	"errors"

	"UMKMGo-backend/internal/types/dto"
	"UMKMGo-backend/internal/types/model"

	"gorm.io/gorm"
)

type UsersRepository interface {
	GetAllUsers(ctx context.Context) ([]model.User, error)
	GetUserByID(ctx context.Context, id int) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	UpdateUser(ctx context.Context, user model.User) (model.User, error)
	DeleteUser(ctx context.Context, user model.User) (model.User, error)

	CreateUMKM(ctx context.Context, umkm model.UMKM, user model.User) (dto.UMKMMobile, error)
	GetUMKMByPhone(ctx context.Context, phone string) (model.UMKM, error)

	GetAllRoles(ctx context.Context) ([]model.Role, error)
	GetRoleByID(ctx context.Context, id int) (model.Role, error)
	GetRoleByName(ctx context.Context, name string) (model.Role, error)
	IsRoleExist(ctx context.Context, id int) bool
	IsPermissionExist(ctx context.Context, id []string) ([]int, bool)
	GetListPermissions(ctx context.Context) ([]model.Permission, error)
	GetListPermissionsByRoleID(ctx context.Context, roleID int) ([]string, error)
	GetListRolePermissions(ctx context.Context) ([]dto.RolePermissionsResponse, error)
	DeletePermissionsByRoleID(ctx context.Context, roleID int) error
	AddRolePermissions(ctx context.Context, roleID int, permissions []int) error
}

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &usersRepository{db}
}

func (user_repo *usersRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := user_repo.db.WithContext(ctx).Preload("Roles").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (user_repo *usersRepository) GetUserByID(ctx context.Context, id int) (model.User, error) {
	var user model.User
	err := user_repo.db.WithContext(ctx).First(&user, "id = ?", id).Error
	if err != nil {
		return model.User{}, errors.New("user not found")
	}

	return user, nil
}

func (user_repo *usersRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := user_repo.db.WithContext(ctx).First(&user, "email = ?", email).Error
	if err != nil {
		return model.User{}, errors.New("user not found")
	}

	return user, nil
}

func (user_repo *usersRepository) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	err := user_repo.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return model.User{}, errors.New("failed to create user")
	}

	return user, nil
}

func (user_repo *usersRepository) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	err := user_repo.db.WithContext(ctx).Save(&user).Error
	if err != nil {
		return model.User{}, errors.New("failed to update user")
	}

	return user, nil
}

func (user_repo *usersRepository) DeleteUser(ctx context.Context, user model.User) (model.User, error) {
	err := user_repo.db.WithContext(ctx).Delete(&user).Error
	if err != nil {
		return model.User{}, errors.New("failed to delete user")
	}

	return user, nil
}

func (user_repo *usersRepository) GetAllRoles(ctx context.Context) ([]model.Role, error) {
	var roles []model.Role
	err := user_repo.db.WithContext(ctx).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (user_repo *usersRepository) GetRoleByID(ctx context.Context, id int) (model.Role, error) {
	var role model.Role
	err := user_repo.db.WithContext(ctx).First(&role, "id = ?", id).Error
	if err != nil {
		return model.Role{}, errors.New("role not found")
	}

	return role, nil
}

func (user_repo *usersRepository) GetRoleByName(ctx context.Context, name string) (model.Role, error) {
	var role model.Role
	err := user_repo.db.WithContext(ctx).First(&role, "name = ?", name).Error
	if err != nil {
		return model.Role{}, errors.New("role not found")
	}

	return role, nil
}

func (user_repo *usersRepository) IsRoleExist(ctx context.Context, id int) bool {
	var role model.Role
	err := user_repo.db.WithContext(ctx).First(&role, "id = ?", id).Error
	return err == nil
}

func (user_repo *usersRepository) IsPermissionExist(ctx context.Context, ids []string) ([]int, bool) {
	var permissionIDs []int
	err := user_repo.db.WithContext(ctx).Model(&model.Permission{}).Where("code IN ?", ids).Pluck("id", &permissionIDs).Error
	if err != nil {
		return nil, false
	}
	return permissionIDs, len(permissionIDs) == len(ids)
}

func (user_repo *usersRepository) GetListPermissions(ctx context.Context) ([]model.Permission, error) {
	var permissions []model.Permission
	err := user_repo.db.WithContext(ctx).Where("parent_id IS NOT NULL").Find(&permissions).Error
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func (user_repo *usersRepository) GetListRolePermissions(ctx context.Context) ([]dto.RolePermissionsResponse, error) {
	var rolePermissions []dto.RolePermissionsResponse
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
	err := user_repo.db.WithContext(ctx).Where("role_id = ?", roleID).Unscoped().Delete(&model.RolePermission{}).Error
	if err != nil {
		return errors.New("failed to delete role permissions")
	}
	return nil
}

func (user_repo *usersRepository) AddRolePermissions(ctx context.Context, roleID int, permissions []int) error {
	var rolePermissions []model.RolePermission
	for _, permissionID := range permissions {
		rolePermissions = append(rolePermissions, model.RolePermission{
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

func (user_repo *usersRepository) CreateUMKM(ctx context.Context, umkm model.UMKM, user model.User) (dto.UMKMMobile, error) {
	var umkmResponse dto.UMKMMobile
	err := user_repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Create user first
		if err := tx.Create(&user).Error; err != nil {
			return errors.New("failed to create user")
		}

		// Set UserID in UMKM
		umkm.UserID = user.ID

		// Create UMKM
		if err := tx.Create(&umkm).Error; err != nil {
			return errors.New("failed to create UMKM")
		}

		// Map UMKM to response DTO
		umkmResponse = dto.UMKMMobile{
			ID:           umkm.ID,
			UserID:       umkm.UserID,
			BusinessName: umkm.BusinessName,
			NIK:          umkm.NIK,
			Gender:       umkm.Gender,
			BirthDate:    umkm.BirthDate.Format("2006-01-02"),
			Phone:        umkm.Phone,
			Address:      umkm.Address,
			ProvinceID:   umkm.ProvinceID,
			CityID:       umkm.CityID,
			District:     umkm.District,
			PostalCode:   umkm.PostalCode,
			KartuType:    umkm.KartuType,
			KartuNumber:  umkm.KartuNumber,
			Fullname:     user.Name,
			Email:        user.Email,
		}
		return nil
	})
	if err != nil {
		return dto.UMKMMobile{}, err
	}
	return umkmResponse, nil
}

func (user_repo *usersRepository) GetUMKMByPhone(ctx context.Context, phone string) (model.UMKM, error) {
	var umkm model.UMKM
	err := user_repo.db.WithContext(ctx).Preload("User").First(&umkm, "phone = ?", phone).Error
	if err != nil {
		return model.UMKM{}, errors.New("UMKM not found")
	}

	return umkm, nil
}
