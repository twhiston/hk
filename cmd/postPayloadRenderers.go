package cmd

import "github.com/spf13/cobra"

func FillStartTimerData(cmd *cobra.Command, args []string, payload *TimerStartPayload) error{
	//TODO - get from viper or 1 as default
	payload.Id = 1
	return nil
}

func FillStopTimerData(cmd *cobra.Command, args []string, payload *TimerStopPayload) error{
	return nil
}
