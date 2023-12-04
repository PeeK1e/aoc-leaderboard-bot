package scrape

type StoredData struct {
	TimeStamp   int64       `json:"time_stamp"`
	LeaderBoard LeaderBoard `json:"leader_board"`
}

type LeaderBoard struct {
	OwnerID int               `json:"owner_id"`
	Event   string            `json:"event"`
	Members map[string]Member `json:"members"`
}

type Member struct {
	LastStarTs         int64                      `json:"last_star_ts"`
	LocalScore         int                        `json:"local_score"`
	Stars              int                        `json:"stars"`
	GlobalScore        int                        `json:"global_score"`
	ID                 int                        `json:"id"`
	Name               string                     `json:"name"`
	CompletionDayLevel map[string]map[string]Star `json:"completion_day_level"`
}

type Star struct {
	GetStarTs int64 `json:"get_star_ts"`
	StarIndex int   `json:"star_index"`
}
