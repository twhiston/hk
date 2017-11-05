package cmd

type StatsResponse struct {
	Overtime        string `json:"overtime"`
	OvertimeSeconds int    `json:"overtime_in_seconds"`
	Vacation struct {
		Redeemed  float32 `json:"redeemed_days"`
		Remaining float32 `json:"remaining_days"`
	} `json:"vacation"`
}

type TimerResponse struct {
	Date            string `json:"date"`
	Start           string `json:"start_time"`
	Duration        string `json:"duration"`
	DurationSeconds int    `json:"duration_in_seconds"`
	Note            string `json:"note"`
	User struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Teams []string `json:"teams"`
	} `json:"user"`
	Type struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Type     string `json:"type"`
		Excluded bool `json:"excluded_from_calculations"`
	}
	Project struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Archived bool `json:"archived"`
	}
}

type TimerTypesResponse []struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Archived bool `json:"archived"`
	Exclude bool `json:"exclude_from calculations"`
}

type TimerStartPayload struct {
	Id int `json:"time_type_id"`
	Start string `json:"start_time,omitempty"`
	ProjectId string `json:"project_id,omitempty"`
	Note string `json:"note,omitempty"`
}

type TimerStopPayload struct {
	End string `json:"end_time,omitempty"`
}

type ProjectResponse []struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Archived bool `json:"archived'"`
	Teams []string `json:"teams,omitempty"`
}