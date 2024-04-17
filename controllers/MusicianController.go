package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ig-rocker/jukebox/initializers"
	"github.com/ig-rocker/jukebox/models"
	"gorm.io/gorm"
)

func CreatMusician(c *gin.Context) {
	var musician models.Musician

	if err := c.ShouldBindJSON(&musician); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Create(&musician).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create musician"})
		return
	}

	c.JSON(201, musician)
}

func UpdateMusician(c *gin.Context) {
	var musician models.Musician
	id := c.Param("id")

	if err := initializers.DB.First(&musician, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Musician not found"})
		return
	}

	if err := c.ShouldBindJSON(&musician); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Save(&musician).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update musician"})
		return
	}

	c.JSON(200, musician)

}

func MusicianByAlbum(c *gin.Context) {
	albumName := c.Param("album")

	var album models.Album
	if err := initializers.DB.Where("album_name = ?", albumName).Preload("Musicians", func(db *gorm.DB) *gorm.DB {
		return db.Order("name asc")
	}).First(&album).Error; err != nil {
		c.JSON(404, gin.H{"error": "Album not found"})
		return
	}
	c.JSON(200, album.Musicians)
}
