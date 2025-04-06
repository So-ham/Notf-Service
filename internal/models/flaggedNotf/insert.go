package flaggedNotf

import (
	"fmt"
	"time"

	"github.com/So-ham/notf-service/internal/entities"
)

// AddComment adds a flaggedNotf to a post in the database
func (m *flaggedNotf) AddFlaggedNotf(userID uint, severity, content string) (err error) {
	flaggedNotf := &entities.FlaggedNotf{
		Severity: severity,
		UserID:   userID,
		Content:  content,
	}
	err = m.DB.Create(flaggedNotf).Error
	return err
}

// UpdateFlagStatus updates the isFlagged status of a flaggedNotf
func (m *flaggedNotf) UpdateFlagStatus(id uint, isFlagged bool) error {
	err := m.DB.Model(&entities.FlaggedNotf{}).Where("id = ?", id).Updates(map[string]interface{}{
		"is_flagged":   isFlagged,
		"moderated_at": time.Now(),
	})
	if err.Error != nil {
		return fmt.Errorf("failed to update flaggedNotf: %v", err.Error)
	}
	return nil
}
