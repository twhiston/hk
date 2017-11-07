package cmd

//User represents a user type in Hakuna
type User struct {
	Id    int      `json:"id"`
	Name  string   `json:"name"`
	Teams []string `json:"teams"`
}

//TimerType represents a timer type in Hakuna
type TimerType struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Excluded bool   `json:"excluded_from_calculations"`
}

//Project represents a project item embedded in another type
type Project struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Archived bool   `json:"archived"`
}

//StatsResponse represents your current stats
type StatsResponse struct {
	Overtime        string `json:"overtime"`
	OvertimeSeconds int    `json:"overtime_in_seconds"`
	Vacation        struct {
		Redeemed  float32 `json:"redeemed_days"`
		Remaining float32 `json:"remaining_days"`
	} `json:"vacation"`
}

//PingResponse contains a pong, not actually used in the cli. Included for completeness
type PingResponse struct {
	Pong string `json:"pong"`
}

//TimerType represents a timer type in Hakuna
type TimerResponse struct {
	Date            string    `json:"date"`
	Start           string    `json:"start_time"`
	Duration        string    `json:"duration"`
	DurationSeconds int       `json:"duration_in_seconds"`
	Note            string    `json:"note"`
	User            User      `json:"user"`
	Type            TimerType `json:"type"`
	Project         Project   `json:"project"`
}

//TimeEntryResponseArray represents a slice of time entries
type TimeEntryResponseArray []TimeEntryResponse

//TimeEntryResponse represents a time entry retrieved from hakuna
type TimeEntryResponse struct {
	Id              int       `json:"id"`
	Starts          string    `json:"starts"`
	Ends            string    `json:"ends"`
	Duration        string    `json:"duration"`
	DurationSeconds int       `json:"duration_in_seconds"`
	Note            string    `json:"note"`
	User            User      `json:"user"`
	Type            TimerType `json:"type"`
	Project         Project   `json:"project"`
}

//TimerTypesResponse represents an slice of possible timer types from the Hakuna API
type TimerTypesResponse []struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Archived bool   `json:"archived"`
	Exclude  bool   `json:"exclude_from calculations"`
}

//ProjectResponse represents an slice of  project responses in Hakuna
type ProjectResponse []struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Archived bool     `json:"archived'"`
	Teams    []string `json:"teams,omitempty"`
}

//TimerStartPayload is passed when starting a new timer
type TimerStartPayload struct {
	Id        int    `json:"time_type_id"`
	Start     string `json:"start_time,omitempty"`
	ProjectId string `json:"project_id,omitempty"`
	Note      string `json:"note,omitempty"`
}

//TimerStopPayload sends an optional end time to the api
type TimerStopPayload struct {
	End string `json:"end_time,omitempty"`
}

//TimeEntryPayload represents the data to send when adding a new time entry
type TimeEntryPayload struct {
	Starts    string `json:"starts"`
	Ends      string `json:"ends"`
	TimeType  int    `json:"time_type_id"`
	ProjectId int    `json:"project_id,omitempty"`
	Note      string `json:"note"`
}
