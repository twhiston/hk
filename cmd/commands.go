package cmd

// This file was NOT auto generated as it contains commands that can not easily be composed in a template

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/twhiston/clitable"
	"log"
	"strconv"
	"time"
)

// GET COMMANDS

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "get your stats for today",
	Long: `The today command will print a combination of the time and timer commands,
This returns all times for the current day, including the running timer and a sum of the time values.

The today command does not support user impersonation`,
	Run: func(cmd *cobra.Command, args []string) {

		api := getAPI()
		resp := new(TimeEntryResponseArray)

		querystring := make(map[string]string, 1)
		querystring["date"] = time.Now().Local().Format("2006-01-02")

		_, err := api.Res("time_entries", resp).Get(querystring)
		HandleError(err)

		table := clitable.New()
		table.AddRow("Starts", "Duration", "Duration Seconds", "Ends")

		durationTotal := 0

		for _, v := range *resp {
			table.AddRow(v.Starts, v.Duration, strconv.Itoa(v.DurationSeconds), v.Ends)
			durationTotal += v.DurationSeconds
		}

		tr := new(TimerResponse)
		_, err = api.Res("timer", tr).Get()
		HandleError(err)
		if tr.DurationSeconds > 0 {
			table.AddRow(time.Now().Local().Format("2006-01-02")+"T"+tr.Start, tr.Duration, strconv.Itoa(tr.DurationSeconds))
			durationTotal += tr.DurationSeconds
		}
		table.AddRow("Today", secondsToHoursAndMinutes(durationTotal), strconv.Itoa(durationTotal))
		table.Print()

	},
}

var timerNewCmd = &cobra.Command{
	Use:   "new",
	Short: "creates a new timer even if a timer is currently running",
	Long: `The new command will create a new timer even if one is currently running by executing stop, then start.
Does not support user aliasing`,
	Run: func(cmd *cobra.Command, args []string) {

		skipStatusCheck = true
		childrenCmds := cmd.Parent().Commands()
		for _, v := range childrenCmds {
			if v.Use == "stop" {
				log.Println("Stopping Current Running Timer (if existent)")
				v.Run(v, args)
				break
			}
		}

		for _, v := range childrenCmds {
			if v.Use == "start" {
				log.Println("Starting new timer")
				v.Run(v, args)
				break
			}
		}

	},
}

var timerRunningCmd = &cobra.Command{
	Use:   "running",
	Short: "returns true if a timer is running or false",
	Run: func(cmd *cobra.Command, args []string) {

		api := getAPI()

		resp := new(TimerResponse)

		querystring := make(map[string]string)

		if impersonate != "" {
			querystring["user_id"] = impersonate
		}

		res := api.Res("timer", resp)
		_, err := res.Get(querystring)
		HandleError(err)
		output := "false"
		if resp.DurationSeconds != 0 {
			output = "true"
		}
		fmt.Println(output)

	},
}

//Initialize commands and options
func init() {

	RootCmd.AddCommand(todayCmd)
	timerCmd.AddCommand(timerNewCmd)
	timerCmd.AddCommand(timerRunningCmd)
}
