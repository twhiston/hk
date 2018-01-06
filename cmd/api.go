package cmd

//User represents a user type in Hakuna
type User struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Teams []string `json:"teams"`
}

//UserResponse is a slice of users
type UserResponse []User

//TimerType represents a timer type in Hakuna
type TimerType struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Excluded bool   `json:"excluded_from_calculations"`
}

//Project represents a project item embedded in another type
type Project struct {
	ID       int    `json:"id"`
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

//TimerResponse represents a timer response from Hakuna
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
	ID              int       `json:"id"`
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
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Archived bool   `json:"archived"`
	Exclude  bool   `json:"exclude_from calculations"`
}

//ProjectResponse represents an slice of  project responses in Hakuna
type ProjectResponse []struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Archived bool     `json:"archived'"`
	Teams    []string `json:"teams,omitempty"`
}

//TimerStartPayload is passed when starting a new timer
type TimerStartPayload struct {
	ID        int    `json:"time_type_id"`
	Start     string `json:"start_time,omitempty"`
	ProjectID string `json:"project_id,omitempty"`
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
	ProjectID int    `json:"project_id,omitempty"`
	Note      string `json:"note"`
}

//AbsenceResponseArray represents the response to the absences path
type AbsenceResponseArray []AbsenceResponse

//AbsenceResponse is a single absence in hakuna
type AbsenceResponse struct {
	ID                      int       `json:"id"`
	Starts                  string    `json:"start_date"`
	Ends                    string    `json:"end_date"`
	FirstHalfDay            bool      `json:"first_half_day"`
	SecondHalfDay           bool      `json:"second_half_day"`
	Recurring               bool      `json:"is_recurring"`
	WeeklyRepeatingInterval int       `json:"weekly_repeating_interval"`
	User                    User      `json:"user"`
	TimeType                TimerType `json:"time_type"`
}

//OrgUser represents a user in an organization
type OrgUser struct {
	User          User `json:"user"`
	FirstHalfDay  bool `json:"absent_first_half_day"`
	SecondHalfDay bool `json:"absent_second_half_day"`
	TimerRunning  bool `json:"has_timer_running"`
}

//OrgResponse is a slice of org users
type OrgResponse []OrgUser
