package shiganshina

import (
	"github.com/b-hivemind/preparer/pkg/db"
	"github.com/b-hivemind/preparer/pkg/tvmazeapi"
	"github.com/gin-gonic/gin"
)

func getShow(c *gin.Context) {
	var show tvmazeapi.Show
	if err := c.ShouldBindUri(&show); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	if show, err := db.GetShowFromID(show.ID); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	} else {
		c.JSON(200, show)
	}
}

func getEpisodeFromSeason(c *gin.Context) {
	var validator EpisodeBySeasonURI
	if err := c.ShouldBindUri(&validator); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	if show, err := db.GetShowFromID(validator.ShowID); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	} else if episode, err := show.GetEpisodeBySeason(validator.Season, validator.Episode); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	} else {
		c.JSON(200, episode)
	}
}

func setEpisodeBySeason(c *gin.Context) {
	var validator EpisodeBySeasonURI
	if err := c.ShouldBindUri(&validator); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	if show, err := db.GetShowFromID(validator.ShowID); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	} else if episode, err := show.GetEpisodeBySeason(validator.Season, validator.Episode); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	} else {
		show.ToggleEpisode(episode.ID)
		if err := db.SetEpisodes(show.ID, show.Episodes.All); err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
			return
		}
		// TODO change this to ID, create a GetEpisodeFromSeason method for Show
		if episode, err := show.GetEpisodeBySeason(validator.Season, validator.Episode); err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
			return
		} else {
			c.JSON(200, episode)
		}
	}
}

func getEpisodeFromSerial(c *gin.Context) {
	var validator EpisodeBySerialURI
	if err := c.ShouldBindUri(&validator); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	if show, err := db.GetShowFromID(validator.ShowID); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	} else if episode, err := show.GetEpisodeBySerial(validator.Serial); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	} else {
		c.JSON(200, episode)
	}
}

func setEpisodeBySerial(c *gin.Context) {
	var validator EpisodeBySerialURI
	if err := c.ShouldBindUri(&validator); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	if show, err := db.GetShowFromID(validator.ShowID); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	} else if episode, err := show.GetEpisodeBySerial(validator.Serial); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	} else {
		show.ToggleEpisode(episode.ID)
		if err := db.SetEpisodes(show.ID, show.Episodes.All); err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
			return
		}
		// TODO change this to ID, create a GetEpisodeFromSeason method for Show
		if episode, err := show.GetEpisodeBySerial(validator.Serial); err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
			return
		} else {
			c.JSON(200, episode)
		}
	}
}

func getNext(c *gin.Context) {
	var show tvmazeapi.Show
	if err := c.ShouldBindUri(&show); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	if show, err := db.GetShowFromID(show.ID); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	} else if episode, err := show.GetNextEpisode(); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	} else {
		c.JSON(200, episode)
	}
}

func setNext(c *gin.Context) {
	var show tvmazeapi.Show
	if err := c.ShouldBindUri(&show); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	if showObj, err := db.GetShowFromID(show.ID); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	} else {
		show = showObj
	}
	show.ToggleNextEpisode()
	if err := db.SetEpisodes(show.ID, show.Episodes.All); err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	if episode, err := show.GetNextEpisode(); err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	} else {
		c.JSON(200, episode)
	}
}

func getShows(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://10.0.0.220:8090")
	if shows, err := db.GetAllShows(); err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	} else {
		c.JSON(200, shows)
	}
}
