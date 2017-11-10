package cmd

import (
	"github.com/spf13/cobra"
	"strconv"
	"testing"
	"time"
)

func TestFillStartTimerData(t *testing.T) {

	cmd := new(cobra.Command)
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

	todayStrings := []string{"", "t", "today"}
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

	yesterdayStrings := []string{"y", "yesterday"}
	yesterday := time.Now().Local().AddDate(0, 0, -1).Format("2006-01-02")

	for _, v := range yesterdayStrings {
		params["date"] = v
		err := timeParamHandler(&params)
		if err != nil {
			t.Fatal(err)
		}
		if params["date"] != yesterday {
			t.Fatal("Passing no value to time parameter handler should return yesterdays date")
		}
	}

}

// Testing for this is a bit more laborious than testing a hand written function call where we would probably pass in the pre verified parameters
// As this is invoked from generated code we have to perform all the logic within the routine. Which is why this test is as it is.
// It kind of sucks that we test the specific message returned by the error but we want to verify at which step of the procedure the error occurred
// nolint
func TestFillTimeEntryData(t *testing.T) {

	startVal := "2017-05-10T23:23"
	endVal := "2017-06-10T12:02"
	timeID := 1
	//fillTimeEntryData(cmd *cobra.Command, args []string, payload *TimeEntryPayload) error
	cmd := new(cobra.Command)
	var args []string
	payload := new(TimeEntryPayload)

	err := fillTimeEntryData(cmd, args, payload)
	if err == nil || err.Error() != "flag accessed but not defined: start" {
		t.Fatal("Must have start flag defined")
	}

	cmd.Flags().String("start", "", "")
	err = fillTimeEntryData(cmd, args, payload)
	_, ok := err.(*time.ParseError)
	if !ok {
		t.Fatal("invalid time format strings should fail validation")
	}

	//This time format should pass
	cmd.Flags().Set("start", startVal)

	err = fillTimeEntryData(cmd, args, payload)
	if err.Error() != "flag accessed but not defined: end" {
		t.Fatal(err)
	}

	cmd.Flags().String("end", "", "")
	err = fillTimeEntryData(cmd, args, payload)
	_, ok = err.(*time.ParseError)
	if !ok {
		t.Fatal("invalid time format strings should fail validation")
	}

	cmd.Flags().Set("end", endVal)

	err = fillTimeEntryData(cmd, args, payload)
	if err.Error() != "flag accessed but not defined: time-id" {
		t.Fatal(err)
	}

	cmd.Flags().Int("time-id", 0, "")
	err = fillTimeEntryData(cmd, args, payload)
	if err.Error() != "time id cannot be less than 1" {
		t.Fatal(err)
	}

	cmd.Flags().Set("time-id", strconv.Itoa(timeID))

	err = fillTimeEntryData(cmd, args, payload)
	if err.Error() != "flag accessed but not defined: project-id" {
		t.Fatal(err)
	}

	cmd.Flags().Int("project-id", 0, "optional project id")

	err = fillTimeEntryData(cmd, args, payload)
	if err.Error() != "flag accessed but not defined: note" {
		t.Fatal(err)
	}
	cmd.Flags().String("note", "", "optional note to add to the entry")

	err = fillTimeEntryData(cmd, args, payload)
	if err != nil {
		t.Fatal(err)
	}

	if payload.Starts != startVal || payload.Ends != endVal || payload.TimeType != timeID {
		t.Fatal("unexpected payload values", payload)
	}

	pid := 23
	note := "Testing a note"
	cmd.Flags().Set("note", note)
	cmd.Flags().Set("project-id", strconv.Itoa(pid))
	err = fillTimeEntryData(cmd, args, payload)
	if err != nil {
		t.Fatal(err)
	}

	if payload.ProjectID != pid || payload.Note != note {
		t.Fatal("unexpected payload values", payload)
	}

}
