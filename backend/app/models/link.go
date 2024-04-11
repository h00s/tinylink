package models

import (
	"time"

	"github.com/h00s/tinylink/internal"
)

type Links []Link

type Link struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	URL       string    `gorm:"size:2048;unique" json:"url"`
	Password  string    `gorm:"size:255" json:"-"`
	Valid     bool      `json:"valid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PublicLink struct {
	ID string `json:"id"`
	Link
}

func (l *Link) ToPublicLink() PublicLink {
	return PublicLink{
		ID:   internal.ShortURIfromID(l.ID),
		Link: *l,
	}
}
