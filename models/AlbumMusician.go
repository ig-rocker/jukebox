package models

import "gorm.io/gorm"

type AlbumMusician struct {
	gorm.Model
	AlbumID        uint
	MusicianID     uint
	AlbumMusicians []*AlbumMusician
}
