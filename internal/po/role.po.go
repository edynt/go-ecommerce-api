package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id; not null; primaryKey; autoIncrement;"`
	RoleName string `gorm:"column:user_name"`
	RoleNote string `gorm:"column:role_note"`
}

func (r *Role) TableName() string {
	return "gb_db_role"
}
