package notificationmodal

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	GormBase
	EntityID         string    `json:"entity_id" sql:"type:uuid" gorm:"type:uuid; index"`
	EntityType       string    `json:"entity_type" gorm:"type:varchar(10); index; not null"`
	NotificationType string    `json:"type" gorm:"type:varchar(16); index; not null;"`
	RequestData      JSONB     `json:"request_data" sql:"type:jsonb" gorm:"type:jsonb"`
	ResponseData     JSONB     `json:"response_data" sql:"type:jsonb" gorm:"type:jsonb"`
	Status           string    `json:"status" gorm:"type:varchar(32); not null; index; default:draft"`
	Success          bool      `json:"success" gorm:"default:false;"`
	SentTime         time.Time `json:"sent_time"`
}

var EntityTypes = map[string]interface{}{
	"user":   true,
	"lead":   true,
	"public": true,
}

var NotificationTypes = map[string]interface{}{
	"mail": true,
}

var NotificationStatus = map[string]interface{}{
	"draft":           true,
	"sending_request": true,
	"failed":          true,
	"completed":       true,
}

func (notification Notification) BeforeSave(tx *gorm.DB) (err error) {

	if EntityTypes[notification.EntityType] != true {
		return errors.New("invalid entity type")
	}

	if NotificationTypes[notification.NotificationType] != true {
		return errors.New("invalid entity type")
	}
	return
}
