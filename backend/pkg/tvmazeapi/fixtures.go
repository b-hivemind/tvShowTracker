package tvmazeapi

const (
	TVMAZE_BASE_API              = "https://api.tvmaze.com/"
	TVMAZE_SINGLE_SEARCH_API     = TVMAZE_BASE_API + "singlesearch/shows?q="
	TVMAZE_MULTI_SEARCH_API      = TVMAZE_BASE_API + "search/shows?q="
	TVMAZE_ALL_EPISODES_API      = TVMAZE_BASE_API + "shows/%d/episodes"
	TVMAZE_EPISODE_BY_NUMBER_API = TVMAZE_BASE_API + "shows/%d/episodebynumber?season=%d&number=%d"
	TVMAZE_EPISODE_BY_ID_API     = TVMAZE_BASE_API + "episodes/%d"
	TVMZAE_SHOW_API              = TVMAZE_BASE_API + "shows/%d"
)
