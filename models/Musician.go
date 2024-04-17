package models

import "gorm.io/gorm"

type Musician struct {
	gorm.Model
	Name         string   `json:"name" gorm:"not null;min:3"`
	MusicianType string   `json:"musician_type"`
	Albums       []*Album `gorm:"many2many:album_musicians;"` // Many-to-many relationship with albums

}
