package models

import "gorm.io/gorm"

type AuditLog struct {
	gorm.Model
	Entity   string
	EntityID uint
	Action   string
}
