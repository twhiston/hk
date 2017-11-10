package cmd

// This file was auto generated by the hk code generator
// DO NOT ALTER THIS FILE MANUALLY

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/twhiston/clitable"
	"io/ioutil"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "get your current stats",
	Long:  `returns the /overview endpoint`,
	Run: func(cmd *cobra.Command, args []string) {

		api := GetAPI()

		resp := new(StatsResponse)

		_, err := api.Res("overview", resp).Get()
		HandleError(err)

		PrintResponse(*resp)

	},
}

var timerCmd = &cobra.Command{
	Use:   "timer",
	Short: "Do things with timers",
	Long:  `Get the current timer, or use subcommands to control timers`,
	Run: func(cmd *cobra.Command, args []string) {

		api := GetAPI()

		resp := new(TimerResponse)

		_, err := api.Res("timer", resp).Get()
		HandleError(err)

		PrintResponse(*resp)

	},
}

var typesCmd = &cobra.Command{
	Use:   "types",
	Short: "get the types of timer available",
	Long:  `returns the possible timer options for your hakuna instance, use these id's with timer commands`,
	Run: func(cmd *cobra.Command, args []string) {

		api := GetAPI()

		resp := new(TimerTypesResponse)

		_, err := api.Res("time_types", resp).Get()
		HandleError(err)

		table := clitable.New()
		for k, v := range *resp {
			if k == 0 {
				table.AddRow(getStructTags(v)...)
			}
			table.AddRow(getStructVals(v)...)
		}
		table.Print()

	},
}

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "get a list of all projects",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		api := GetAPI()

		resp := new(ProjectResponse)

		_, err := api.Res("projects", resp).Get()
		HandleError(err)

	},
}

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "get time entries",
	Long:  `get time entries for a specific date`,
	Run: func(cmd *cobra.Command, args []string) {

		api := GetAPI()

		resp := new(TimeEntryResponseArray)

		querystring := make(map[string]string, 1)
		value, e := cmd.Flags().GetString("date")
		HandleError(e)
		querystring["date"] = value

		pe := timeParamHandler(&querystring)
		HandleError(pe)

		_, err := api.Res("time_entries", resp).Get(querystring)
		HandleError(err)

		table := clitable.New()
		for k, v := range *resp {
			if k == 0 {
				table.AddRow(getStructTags(v)...)
			}
			table.AddRow(getStructVals(v)...)
		}
		table.Print()

	},
}

var entryCmd = &cobra.Command{
	Use:   "entry",
	Short: "get a time entry specified by an id",
	Long:  `returns a specific time entry, which is given as an argument to the command`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < (0 + 1) {
			HandleError(errors.New("argument 0 is required for this command"))
		}

		api := GetAPI()

		resp := new(TimeEntryResponse)

		_, err := api.Res("time_entries", resp).Id(args[0]).Get()
		HandleError(err)

		PrintResponse(*resp)

	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a new timer",
	Long:  `Creates a new running timer starting from now`,
	Run: func(cmd *cobra.Command, args []string) {

		api := GetAPI()

		resp := new(TimerResponse)

		payload := new(TimerStartPayload)

		// Payload renderer must have signature (cmd *cobra.Command, args []string, payload *TimerStartPayload) (*TimerStartPayload, error)
		err := fillStartTimerData(cmd, args, payload)
		HandleError(err)
		r, err := api.Res("timer", resp).Post(payload)
		HandleError(err)

		if r.Raw.StatusCode != 201 {
			defer deferredBodyClose(r)
			bodyBytes, err := ioutil.ReadAll(r.Raw.Body)
			HandleError(err)
			HandleError(errors.New(string(bodyBytes)))
		}

	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a time entry in the calendar",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		api := GetAPI()

		resp := new(TimeEntryResponse)

		payload := new(TimeEntryPayload)

		// Payload renderer must have signature (cmd *cobra.Command, args []string, payload *TimeEntryPayload) (*TimeEntryPayload, error)
		err := fillTimeEntryData(cmd, args, payload)
		HandleError(err)
		r, err := api.Res("time_entries", resp).Post(payload)
		HandleError(err)

		if r.Raw.StatusCode != 201 {
			defer deferredBodyClose(r)
			bodyBytes, err := ioutil.ReadAll(r.Raw.Body)
			HandleError(err)
			HandleError(errors.New(string(bodyBytes)))
		}

		PrintResponse(*resp)

	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop a timer",
	Long:  `stops a timer and optionally sets a stop time`,
	Run: func(cmd *cobra.Command, args []string) {

		api := GetAPI()

		resp := new(TimeEntryResponse)

		payload := new(TimerStopPayload)

		r, err := api.Res("timer", resp).Put(payload)
		HandleError(err)

		if r.Raw.StatusCode != 200 {
			defer deferredBodyClose(r)
			bodyBytes, err := ioutil.ReadAll(r.Raw.Body)
			HandleError(err)
			HandleError(errors.New(string(bodyBytes)))
		}

		PrintResponse(*resp)

	},
}

var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "cancel a timer",
	Long:  `deletes the currently running timer`,
	Run: func(cmd *cobra.Command, args []string) {

		api := GetAPI()

		r, err := api.Res("timer").Delete()
		HandleError(err)

		if r.Raw.StatusCode != 205 {
			defer deferredBodyClose(r)
			bodyBytes, err := ioutil.ReadAll(r.Raw.Body)
			HandleError(err)
			HandleError(errors.New(string(bodyBytes)))
		}

	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a time",
	Long:  `deletes a time via it's id.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < (0 + 1) {
			HandleError(errors.New("argument 0 is required for this command"))
		}

		api := GetAPI()

		r, err := api.Res("time_entries").Id(args[0]).Delete()
		HandleError(err)

		if r.Raw.StatusCode != 204 {
			defer deferredBodyClose(r)
			bodyBytes, err := ioutil.ReadAll(r.Raw.Body)
			HandleError(err)
			HandleError(errors.New(string(bodyBytes)))
		}

	},
}

//Initialize commands and options
func init() {

	RootCmd.AddCommand(statsCmd)

	RootCmd.AddCommand(timerCmd)

	timerCmd.AddCommand(typesCmd)

	RootCmd.AddCommand(projectsCmd)

	RootCmd.AddCommand(timeCmd)
	timeCmd.Flags().StringP("date", "d", "", "enter the date that you want to look for details of. If left blank will use todays date")

	timeCmd.AddCommand(entryCmd)

	timerCmd.AddCommand(startCmd)
	startCmd.Flags().Int("id", 1, "enter the id of the timer payload, should correspond to a timer type ID, default 1")

	timeCmd.AddCommand(createCmd)
	createCmd.Flags().String("start", "", "enter the start date in the format yyyy-dd-mmThh:mm")
	createCmd.Flags().String("end", "", "enter the start date in the format yyyy-dd-mmThh:mm")
	createCmd.Flags().Int("time-id", 1, "enter the time type id. You can get these with the types command, defaults to 1 which is usually Arbeit")
	createCmd.Flags().Int("project-id", 0, "optional project id")
	createCmd.Flags().String("note", "", "optional note to add to the entry")

	timerCmd.AddCommand(stopCmd)

	timerCmd.AddCommand(cancelCmd)

	timeCmd.AddCommand(deleteCmd)
}
