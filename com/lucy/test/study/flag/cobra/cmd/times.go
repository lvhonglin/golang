package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var echoTimes int
var cmdTimes = &cobra.Command{
	Use:   "time [string to time]",
	Short: "Echo anything to the screen more times",
	Long:  (`echo things multiple times back to the user y providing a count and a string`),
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < echoTimes; i++ {
			fmt.Println("Echo: " + strings.Join(args, " "))
		}
	},
}

func init() {
	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
	//times是echo的子目录
	cmdEcho.AddCommand(cmdTimes)
}
