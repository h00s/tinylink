package models

type Accesses []Access

type Access struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	LinkID    uint   `gorm:"not null" json:"link_id"`
	CreatedAt string `json:"created_at"`
}
