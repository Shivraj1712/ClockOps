package domain

import (
	"time"

	"github.com/google/uuid"
)

type Attendence struct {
	ID           uuid.UUID  `json:"attendence_id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	EmployeeID   uuid.UUID  `json:"employee_id" gorm:"type:uuid;not null;index"`
	Employee     User       `json:"user" gorm:"foreignkey:EmployeeID"`
	AdminId      uuid.UUID  `json:"admin_id" gorm:"type:uuid;not null;index"`
	Admin        User       `json:"admin" gorm:"foreignKey:AdminID"`
	CheckInTime  time.Time  `json:"checkin_time" gorm:"type:timestamp;not null"`
	CheckOutTime *time.Time `json:"checkout_time" gorm:"type:time"`
	CreatedAt    time.Time  `json:"autoCreateTime"`
}
