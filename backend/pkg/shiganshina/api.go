package shiganshina

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Start() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Search
	r.POST("/search", handleSearchShows)
	// Import
	r.POST("/import/:showid", handleImportShow)
	// Shows
	r.GET("/shows", getShows)
	r.GET("/shows/:showid", getShow)
	r.GET("/shows/:showid/season/:season/episode/:episode", getEpisodeFromSeason)
	r.POST("shows/:showid/season/:season/episode/:episode", setEpisodeBySeason)
	r.GET("/shows/:showid/:serial", getEpisodeFromSerial)
	r.POST("/shows/:showid/:serial", setEpisodeBySerial)
	r.GET("/shows/:showid/next", getNext)
	r.POST("/shows/:showid/next", setNext)

	s := &http.Server{
		Addr:    ":10090",
		Handler: r,
	}
	s.ListenAndServe()
}
