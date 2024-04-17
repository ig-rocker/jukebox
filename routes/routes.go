package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ig-rocker/jukebox/controllers"
)

func Run() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "welcomee to Jukebox")
	})
	r.POST("/create-album", controllers.CreateAlbum)
	r.PUT("/update-album", controllers.UpdateAlbum)
	r.GET("/all-albums", controllers.GetAllAlbums)
	r.GET("/album/:musician", controllers.GetAlbumByMusician)

	r.POST("/create-musician", controllers.CreatMusician)
	r.PUT("/update-musician", controllers.UpdateMusician)
	r.GET("/musician/:album", controllers.MusicianByAlbum)

	r.Run(":3030")
}
