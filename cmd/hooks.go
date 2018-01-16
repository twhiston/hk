package cmd

import (
	"errors"
	"fmt"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

// nolint: gocyclo
func fillStartTimerData(cmd *cobra.Command, args []string, payload *TimerStartPayload) error {

	var err error
	payload.ID, err = cmd.Parent().PersistentFlags().GetInt("id")
	if err != nil {
		return err
	}

	pid, err := cmd.Parent().PersistentFlags().GetInt("project")
	if err != nil {
		return err
	}
	if pid != 0 {
		payload.ProjectID = strconv.Itoa(pid)
	}

	note, err := cmd.Parent().PersistentFlags().GetString("note")
	if err != nil {
		return err
	}
	payload.Note = note

	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	start, err := cmd.Parent().PersistentFlags().GetString("start")
	if err != nil {
		return err
	}
	if start != "" {
		r, err := w.Parse(start, time.Now())
		if err != nil {
			return err
		}
		if r == nil {
			return errors.New("cannot parse starts string")
		}

		start = r.Time.Format("15:04")
		payload.Start = start
	}

	// Payload can never be 0 so default to 1
	if payload.ID == 0 {
		if verbose {
			fmt.Println("Corrected id of 0 to 1")
		}
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
	starts, err := cmd.Parent().PersistentFlags().GetString("start")
	if err != nil {
		return err
	}

	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	if (starts == "" && !allowBlank) || (starts != "" && validate) {

		r, err := w.Parse(starts, time.Now())
		if err != nil {
			return err
		}
		if r == nil {
			return errors.New("cannot parse starts string")
		}

		starts = r.Time.Format("2006-01-02T15:04:05Z07:00")
	}

	ends, err := cmd.Parent().PersistentFlags().GetString("end")
	if err != nil {
		return err
	}
	if (ends == "" && !allowBlank) || (ends != "" && validate) {
		r, err := w.Parse(ends, time.Now())
		if err != nil {
			return err
		}
		if r == nil {
			return errors.New("cannot parse end string")
		}

		ends = r.Time.Format("2006-01-02T15:04:05Z07:00")
	}

	timeid, err := cmd.Parent().PersistentFlags().GetInt("time-id")
	if err != nil {
		return err
	}

	if (timeid == 0 && !allowBlank) || (timeid != 0 && validate) {
		if timeid < 1 {
			return errors.New("time id cannot be less than 1")
		}
	}

	projectid, err := cmd.Parent().PersistentFlags().GetInt("project")
	if err != nil {
		return err
	}

	note, err := cmd.Parent().PersistentFlags().GetString("note")
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

	if (*params)["date"] == "" {
		(*params)["date"] = time.Now().Format("2006-01-02")
		return nil
	}

	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	r, err := w.Parse((*params)["date"], time.Now())
	if err != nil {
		return err
	}
	if r == nil {
		return errors.New("cannot parse date string " + (*params)["date"])
	}

	(*params)["date"] = r.Time.Format("2006-01-02")

	return nil
}

func absenceParamHandler(params *map[string]string) error {
	if (*params)["year"] == "" {
		(*params)["year"] = time.Now().Local().Format("2006")
	}
	return nil
}
