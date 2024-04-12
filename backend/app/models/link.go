package models

import (
	"time"

	"github.com/h00s/tinylink/internal"
)

type Links []Link

var LinkPermittedParams = []string{"URL", "Password"}

type Link struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	URL       string    `gorm:"uniqueIndex:idx_url_password;size:2048" json:"url"`
	Password  string    `gorm:"uniqueIndex:idx_url_password;size:255" json:"password"`
	Valid     bool      `json:"valid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PublicLink struct {
	ID       string `json:"id"`
	Password string `json:"-"`
	Link
}

func (l *Link) ToPublicLink() PublicLink {
	return PublicLink{
		ID:   internal.ShortURIfromID(l.ID),
		Link: *l,
	}
}
