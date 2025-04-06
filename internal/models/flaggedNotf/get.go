package flaggedNotf

import (
	"context"
	"time"

	"github.com/So-ham/notf-service/internal/entities"
)

func (m *flaggedNotf) GetNotifications(ctx context.Context, userID uint) (notifications []entities.FlaggedNotf, err error) {
	// Fetch notifications
	err = m.DB.Where("user_id = ? AND notified = ?", userID, false).Find(&notifications).Error
	if err != nil {
		return nil, err
	}

	// If no notifications, skip update
	if len(notifications) == 0 {
		return
	}

	// Collect IDs of the notifications
	var ids []uint
	for _, n := range notifications {
		ids = append(ids, n.ID)
	}

	// Mark them as notified
	updateResult := m.DB.Model(&entities.FlaggedNotf{}).
		Where("id IN ?", ids).
		Updates(map[string]interface{}{
			"notified":    true,
			"notified_at": time.Now(),
		})

	if updateResult.Error != nil {
		return nil, updateResult.Error
	}

	return
}
