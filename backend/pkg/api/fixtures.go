package api

type EpisodeBySeasonURI struct {
	ShowID  int `binding:"required" uri:"showid"`
	Season  int `binding:"required" uri:"season"`
	Episode int `binding:"required" uri:"episode"`
}

type EpisodeBySerialURI struct {
	ShowID int `binding:"required" uri:"showid"`
	Serial int `binding:"required" uri:"serial"`
}

type SearchShowsQuery struct {
	Message string `form:"message"`
}

type ImportShowsQuery struct {
	Id int `binding:"required" uri:"showid"`
}
