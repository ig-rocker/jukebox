package initializers

import "github.com/ig-rocker/jukebox/models"

func SyncDB() {
	DB.AutoMigrate(&models.Album{})
	DB.AutoMigrate(&models.Musician{})
	DB.AutoMigrate(&models.AlbumMusician{})

}
