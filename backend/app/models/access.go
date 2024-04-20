package models

import "time"

type Accesses []Access

type Access struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	LinkID    uint      `gorm:"index;not null" json:"link_id"`
	CreatedAt time.Time `json:"created_at"`
}
