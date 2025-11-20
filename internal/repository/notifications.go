package repository

import (
	"context"

	"UMKMGo-backend/internal/types/model"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	CreateNotification(ctx context.Context, notification model.Notification) error
	GetNotificationsByUMKMID(ctx context.Context, umkmID int, limit, offset int) ([]model.Notification, error)
	GetUnreadCount(ctx context.Context, umkmID int) (int64, error)
	MarkAsRead(ctx context.Context, notificationIDs []int, umkmID int) error
	MarkAllAsRead(ctx context.Context, umkmID int) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db}
}

func (r *notificationRepository) CreateNotification(ctx context.Context, notification model.Notification) error {
	return r.db.WithContext(ctx).Create(&notification).Error
}

func (r *notificationRepository) GetNotificationsByUMKMID(ctx context.Context, umkmID int, limit, offset int) ([]model.Notification, error) {
	var notifications []model.Notification
	err := r.db.WithContext(ctx).
		Where("umkm_id = ? AND deleted_at IS NULL", umkmID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (r *notificationRepository) GetUnreadCount(ctx context.Context, umkmID int) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&model.Notification{}).
		Where("umkm_id = ? AND is_read = ? AND deleted_at IS NULL", umkmID, false).
		Count(&count).Error
	return count, err
}

func (r *notificationRepository) MarkAsRead(ctx context.Context, notificationIDs []int, umkmID int) error {
	return r.db.WithContext(ctx).
		Model(&model.Notification{}).
		Where("id IN ? AND umkm_id = ?", notificationIDs, umkmID).
		Updates(map[string]interface{}{
			"is_read": true,
			"read_at": "NOW()",
		}).Error
}

func (r *notificationRepository) MarkAllAsRead(ctx context.Context, umkmID int) error {
	return r.db.WithContext(ctx).
		Model(&model.Notification{}).
		Where("umkm_id = ? AND is_read = ?", umkmID, false).
		Updates(map[string]interface{}{
			"is_read": true,
			"read_at": "NOW()",
		}).Error
}