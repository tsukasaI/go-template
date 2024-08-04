package model

import (
	"time"

	ulid "github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Sample struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	gorm.DeletedAt
}

func (p *Sample) BeforeCreate(_ *gorm.DB) (err error) {
	p.ID = ulid.Make().String()
	return
}
