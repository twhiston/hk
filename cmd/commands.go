package cmd

// This file was NOT auto generated as it contains commands that can not easily be composed in a template

import (
	"github.com/spf13/cobra"
	"github.com/twhiston/clitable"
	"strconv"
	"time"
	"fmt"
)

// GET COMMANDS

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "get your stats for today",
	Long: `The today command will print a combination of the time and timer commands,
This returns all times for the current day, including the running timer and a sum of the time values`,
	Run: func(cmd *cobra.Command, args []string) {

		//TODO does not really work at the moment as both structs have different columns
		//This means that we need to actually extract the duration data and present it in a different way.
		//Should be easy right :P
		api := GetApi()
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
		table.AddRow("Today", secondsToHoursAndMinutes(durationTotal), strconv.Itoa(durationTotal) )
		table.Print()

	},
}


func secondsToHoursAndMinutes(inSeconds int) string {
	minutes := inSeconds / 60
	hours := minutes / 60
	seconds := minutes % 60
	return fmt.Sprintf("%d:%02d", hours, seconds)
}


//Initialize commands and options
func init() {

	RootCmd.AddCommand(todayCmd)
}
