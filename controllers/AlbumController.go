package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ig-rocker/jukebox/initializers"
	"github.com/ig-rocker/jukebox/models"
)

func CreateAlbum(c *gin.Context) {
	var album models.Album

	if err := c.BindJSON(&album); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Create(&album).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create album"})
		return
	}

	c.JSON(200, album)
}

func UpdateAlbum(c *gin.Context) {
	var album models.Album
	id := c.Param("id")

	if err := initializers.DB.First(&album, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Album not found"})
		return
	}

	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Save(&album).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update album"})
		return
	}

	c.JSON(200, album)
}

func GetAllAlbums(c *gin.Context) {
	var albums []models.Album

	if err := initializers.DB.Order("date_of_release asc").Find(&albums).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve albums"})
		return
	}

	c.JSON(200, albums)
}

func GetAlbumByMusician(c *gin.Context) {
	musicianName := c.Param("musician")

	var musician models.Musician
	if err := initializers.DB.Where("name = ?", musicianName).Preload("Albums").Order("price asc").First(&musician).Error; err != nil {
		c.JSON(404, gin.H{"error": "Musician not found"})
		return
	}

	c.JSON(200, musician.Albums)
}
