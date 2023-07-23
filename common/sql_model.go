package common

import "time"

type SQLModel struct {
	ID       int        `json:"-" gorm:"primaryKey;column:id"`
	FakeID   *UID       `json:"id" gorm:"-"`
	Status   int        `json:"status" gorm:"column:status;default:1"`
	CreateAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdateAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (m *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.ID), dbType, 0)
	m.FakeID = &uid
}
