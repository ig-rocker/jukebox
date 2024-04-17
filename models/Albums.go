package models

import (
	"time"

	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	AlbumName     string      `json:"album_name" gorm:"not null; min:5"`
	DateOfRelease time.Time   `josn:"date_of_release" gorm:"not null"`
	Genre         string      `json:"genre"`
	Price         float64     `json:"price" gorm:"not null; check: price>=100 and price <=1000"`
	Description   string      `json:"description"`
	Musicians     []*Musician `gorm:"many2many:album_musicians;"` // Many-to-many relationship with musicians

}
