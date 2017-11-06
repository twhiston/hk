package cmd

import (
	"github.com/spf13/cobra"
	"time"
	"errors"
)

func FillStartTimerData(cmd *cobra.Command, args []string, payload *TimerStartPayload) error {
	payload.Id = 1
	return nil
}

func FillStopTimerData(cmd *cobra.Command, args []string, payload *TimerStopPayload) error {
	return nil
}

func FillTimeEntryData(cmd *cobra.Command, args []string, payload *TimeEntryPayload) error {
	starts, err := cmd.Flags().GetString("start")
	HandleError(err)
	if starts == "" {
		return errors.New("must set start time")
	}
	ends, err := cmd.Flags().GetString("end")
	HandleError(err)
	if ends == "" {
		return errors.New("must set end time")
	}

	timeid, err := cmd.Flags().GetInt("time-id")
	HandleError(err)
	if timeid < 1 {
		return errors.New("time id cannot be less than 1")
	}

	projectid, err := cmd.Flags().GetInt("project-id")
	HandleError(err)

	note, err := cmd.Flags().GetString("note")
	HandleError(err)

	payload.Starts = starts
	payload.Ends = ends
	payload.TimeType = timeid
	payload.ProjectId = projectid
	payload.Note = note

	return nil
}

func TimeParamHandler(params *map[string]string) error {
	if (*params)["date"] == "" || (*params)["date"] == "today" || (*params)["date"] == "t" {
		(*params)["date"] = time.Now().Local().Format("2006-01-02")
	} else if (*params)["date"] == "yesterday" || (*params)["date"] == "y" {
		(*params)["date"] = time.Now().Local().AddDate(0, 0, -1).Format("2006-01-02")
	}
	return nil
}
