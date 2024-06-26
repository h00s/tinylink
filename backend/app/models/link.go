package models

import (
	"time"

	"github.com/h00s/tinylink/internal"
)

type Links []Link

var LinkPermittedParams = []string{"URL", "Password", "Valid"}

type Link struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	URL       string     `gorm:"uniqueIndex:idx_url_password;size:2048" json:"url"`
	Password  string     `gorm:"uniqueIndex:idx_url_password;size:60" json:"password"`
	Valid     bool       `json:"-"`
	ExpiresAt *time.Time `json:"expires_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Accesses  Accesses   `gorm:"constraint:OnDelete:CASCADE" json:"-"`
}

type PublicLink struct {
	ID       string `json:"id"`
	Password string `json:"password,omitempty"`
	Link
}

func (l *Link) ToPublicLink() PublicLink {
	return PublicLink{
		ID:       internal.ShortURIfromID(l.ID),
		Password: "",
		Link:     *l,
	}
}
