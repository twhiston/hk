package cmd

// This file was NOT auto generated as it contains commands that can not easily be composed in a template

import (
	"github.com/spf13/cobra"
	"github.com/twhiston/clitable"
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

		//TODO does not really work at the moment as both structs have different columns
		//This means that we need to actually extract the duration data and present it in a different way.
		//Should be easy right :P
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

//Initialize commands and options
func init() {

	RootCmd.AddCommand(todayCmd)
}
