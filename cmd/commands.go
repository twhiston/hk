// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"errors"
	"io/ioutil"
	"github.com/twhiston/clitable"
)

// statsCmd represents the ls command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "get an overview of your current overtime and holidays",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		api := GetApi()
		resp := new(StatsResponse)
		api.Res("overview", resp).Get()
		PrintResponse(*resp)

	},
}

var timerTypesCmd = &cobra.Command{
	Use:   "types",
	Short: "Get Timer Types",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		api := GetApi()
		resp := new(TimerTypesResponse)
		_,err := api.Res("time_types", &resp).Get()
		handleError(err)

		//TODO - make generic
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

var timerCmd = &cobra.Command{
	Use:   "timer",
	Short: "Do things with timers",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		api := GetApi()
		resp := new(TimerResponse)
		api.Res("timer", resp).Get()
		PrintResponse(*resp)
	},
}

var timerStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start a timer",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		api := GetApi()
		resp := new(TimerResponse)
		payload := new(TimerStartPayload)
		//TODO - set via flag
		payload.Id = 1
		r, err := api.Res("timer", resp).Post(payload)
		handleError(err)
		if r.Raw.StatusCode != 200 {
			defer r.Raw.Body.Close()
			bodyBytes, err := ioutil.ReadAll(r.Raw.Body)
			handleError(err)
			handleError(errors.New(string(bodyBytes)))
		}
		PrintResponse(*resp)
	},
}

var timerStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop a timer",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		api := GetApi()
		resp := new(TimerResponse)
		payload := new(TimerStopPayload)
		//TODO stop time flag
		api.Res("timer", resp).Put(payload)
		PrintResponse(*resp)
	},
}

var timerCancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "cancel a timer",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		api := GetApi()
		resp := new(TimerResponse)
		//TODO stop time flag
		api.Res("timer", resp).Delete()
		PrintResponse(*resp)
	},
}


func init() {
	RootCmd.AddCommand(statsCmd)

	RootCmd.AddCommand(timerCmd)
	timerCmd.AddCommand(timerTypesCmd)
	timerCmd.AddCommand(timerStartCmd)
	timerCmd.AddCommand(timerStopCmd)
	timerCmd.AddCommand(timerCancelCmd)
}
