package notificationmodal

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	// "reflect"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type GormBase struct {
	ID        string     `json:"id" sql:"type:uuid;primary_key;default:uuid_generate_v4()" gorm:"type:uuid"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *JSONB) Scan(value interface{}) error {
	switch t := value.(type) {
	case []uint8:
		fmt.Println("yaha hai baba")
		if err := json.Unmarshal(t, &j); err != nil {
			return err
		}
	default:
		if value.(string) != "null" {
			if err := json.Unmarshal([]byte(value.(string)), &j); err != nil {
				return err
			}
		}
	}

	return nil
}

func (gormBase *GormBase) BeforeCreate(tx *gorm.DB) (err error) {
	u := uuid.NewV4()
	uuid := u.String()
	gormBase.ID = uuid
	return
}
