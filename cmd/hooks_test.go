package cmd

import (
	"github.com/spf13/cobra"
	"strconv"
	"testing"
	"time"
)

func TestFillStartTimerData(t *testing.T) {

	parent := new(cobra.Command)
	cmd := new(cobra.Command)
	parent.AddCommand(cmd)

	cmd.Parent().PersistentFlags().Int("id", 1, "")
	cmd.Parent().PersistentFlags().Int("project", 23, "")
	cmd.Parent().PersistentFlags().String("note", "", "")
	cmd.Parent().PersistentFlags().String("start", "23:11", "")

	var args []string

	payload := new(TimerStartPayload)
	err := fillStartTimerData(cmd, args, payload)

	if err != nil {
		t.Fatal(err)
	}

	if payload.ID != 1 {
		t.Fatal("If no ID is supplied then default 1 should be set, received", payload.ID)
	}

}

func TestTimeParamHandler(t *testing.T) {

	todayStrings := []string{"", "today"}
	today := time.Now().Local().Format("2006-01-02")
	params := make(map[string]string, 1)

	for _, v := range todayStrings {
		params["date"] = v
		err := timeParamHandler(&params)
		if err != nil {
			t.Fatal(err)
		}
		if params["date"] != today {
			t.Fatal("Passing no value to time parameter handler should return todays date")
		}
	}

	yesterdayStrings := []string{"yesterday"}
	yesterday := time.Now().Local().AddDate(0, 0, -1).Format("2006-01-02")

	for _, v := range yesterdayStrings {
		params["date"] = v
		err := timeParamHandler(&params)
		if err != nil {
			t.Fatal(err)
		}
		if params["date"] != yesterday {
			t.Fatal("Passing no value to time parameter handler should return yesterdays date", params["date"], yesterday)
		}
	}

}

func TestAbsenceParamHandler(t *testing.T) {

	data := make(map[string]string, 1)
	data["year"] = ""
	err := absenceParamHandler(&data)
	if err != nil {
		t.Fatal(err)
	}

	if data["year"] != time.Now().Local().Format("2006") {
		t.Fatal("year in data does not match this year", data["year"])
	}

}

// Testing for this is a bit more laborious than testing a hand written function call where we would probably pass in the pre verified parameters
// As this is invoked from generated code we have to perform all the logic within the routine. Which is why this test is as it is.
// It kind of sucks that we test the specific message returned by the error but we want to verify at which step of the procedure the error occurred
// nolint
func TestFillTimeEntryData(t *testing.T) {

	startVal := "monday 3pm"
	endVal := "monday 18:15"
	timeID := 1
	//fillTimeEntryData(cmd *cobra.Command, args []string, payload *TimeEntryPayload) error
	parent := new(cobra.Command)
	cmd := new(cobra.Command)
	parent.AddCommand(cmd)
	var args []string
	payload := new(TimeEntryPayload)

	err := fillRequiredTimeEntryData(cmd, args, payload)
	if err == nil || err.Error() != "flag accessed but not defined: start" {
		t.Fatal("Must have start flag defined", err.Error())
	}

	cmd.Parent().PersistentFlags().String("start", "bingbong", "")
	err = fillRequiredTimeEntryData(cmd, args, payload)
	if err == nil {
		t.Fatal("invalid time format strings should fail validation")
	}

	//This time format should pass
	cmd.Parent().PersistentFlags().Set("start", startVal)

	err = fillRequiredTimeEntryData(cmd, args, payload)
	if err.Error() != "flag accessed but not defined: end" {
		t.Fatal(err)
	}

	cmd.Parent().PersistentFlags().String("end", "", "")
	err = fillRequiredTimeEntryData(cmd, args, payload)
	if err == nil {
		t.Fatal("invalid time format strings should fail validation")
	}

	cmd.Parent().PersistentFlags().Set("end", endVal)

	err = fillRequiredTimeEntryData(cmd, args, payload)
	if err.Error() != "flag accessed but not defined: time-id" {
		t.Fatal(err)
	}

	cmd.Parent().PersistentFlags().Int("time-id", 0, "")
	err = fillRequiredTimeEntryData(cmd, args, payload)
	if err.Error() != "time id cannot be less than 1" {
		t.Fatal(err)
	}

	cmd.Parent().PersistentFlags().Set("time-id", strconv.Itoa(timeID))

	err = fillRequiredTimeEntryData(cmd, args, payload)
	if err.Error() != "flag accessed but not defined: project" {
		t.Fatal(err)
	}

	cmd.Parent().PersistentFlags().Int("project", 0, "optional project id")

	err = fillRequiredTimeEntryData(cmd, args, payload)
	if err.Error() != "flag accessed but not defined: note" {
		t.Fatal(err)
	}
	cmd.Parent().PersistentFlags().String("note", "", "optional note to add to the entry")

	err = fillRequiredTimeEntryData(cmd, args, payload)
	if err != nil {
		t.Fatal(err)
	}

	pid := 23
	note := "Testing a note"
	cmd.Parent().PersistentFlags().Set("note", note)
	cmd.Parent().PersistentFlags().Set("project", strconv.Itoa(pid))
	err = fillRequiredTimeEntryData(cmd, args, payload)
	if err != nil {
		t.Fatal(err)
	}

	if payload.ProjectID != pid || payload.Note != note {
		t.Fatal("unexpected payload values", payload)
	}

}
