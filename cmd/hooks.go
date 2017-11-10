package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

func fillStartTimerData(cmd *cobra.Command, args []string, payload *TimerStartPayload) error {
	payload.ID = viper.GetInt("id")
	// Payload can never be 0 so default to 1
	if payload.ID == 0 {
		payload.ID = 1
	}
	return nil
}

func fillRequiredTimeEntryData(cmd *cobra.Command, args []string, payload *TimeEntryPayload) error {
	return fillTimeEntryData(cmd, args, payload, true, false)
}

func fillOptionalTimeEntryData(cmd *cobra.Command, args []string, payload *TimeEntryPayload) error {
	return fillTimeEntryData(cmd, args, payload, true, true)
}

//nolint
func fillTimeEntryData(cmd *cobra.Command, args []string, payload *TimeEntryPayload, validate bool, allowBlank bool) error {
	starts, err := cmd.Flags().GetString("start")
	if err != nil {
		return err
	}

	if (starts == "" && !allowBlank) || (starts != "" && validate) {
		// Format of this time string layout is EXTREMELY important for parsing to work
		// https://stackoverflow.com/questions/14106541/go-parsing-date-time-strings-which-are-not-standard-formats
		_, err = time.Parse("2006-02-01T15:04", starts)
		if err != nil {
			return err
		}
	}

	ends, err := cmd.Flags().GetString("end")
	if err != nil {
		return err
	}
	if (ends == "" && !allowBlank) || (ends != "" && validate) {
		_, err = time.Parse("2006-02-01T15:04", ends)
		if err != nil {
			return err
		}
	}

	timeid, err := cmd.Flags().GetInt("time-id")
	if err != nil {
		return err
	}

	if (timeid == 0 && !allowBlank) || (timeid != 0 && validate) {
		if timeid < 1 {
			return errors.New("time id cannot be less than 1")
		}
	}

	projectid, err := cmd.Flags().GetInt("project-id")
	if err != nil {
		return err
	}

	note, err := cmd.Flags().GetString("note")
	if err != nil {
		return err
	}

	payload.Starts = starts
	payload.Ends = ends
	payload.TimeType = timeid
	payload.ProjectID = projectid
	payload.Note = note

	return nil
}

//Time param handler only deals with special cases, as other dates should either be entered correctly or fail
func timeParamHandler(params *map[string]string) error {
	if (*params)["date"] == "" || (*params)["date"] == "today" || (*params)["date"] == "t" {
		(*params)["date"] = time.Now().Local().Format("2006-01-02")
	} else if (*params)["date"] == "yesterday" || (*params)["date"] == "y" {
		(*params)["date"] = time.Now().Local().AddDate(0, 0, -1).Format("2006-01-02")
	}
	return nil
}

func absenceParamHandler(params *map[string]string) error {
	if (*params)["year"] == "" {
		(*params)["year"] = time.Now().Local().Format("2006")
	}
	return nil
}
