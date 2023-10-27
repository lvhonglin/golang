package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var echoTimes int
var kaiguan bool
var cmdTimes = &cobra.Command{
	Use:   "time [string to time]",
	Short: "Echo anything to the screen more times",
	Long:  (`echo things multiple times back to the user y providing a count and a string`),
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if kaiguan {
			for i := 0; i < echoTimes; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		println("前置执行")
	},
	//postRunE和postRun只能写一个，不能写2个，写2个其中一个会被覆盖
	PostRunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("后置处理")
		if strings.Contains(args[0], "err") {
			return errors.New("字符里面有敏感信息")
		}
		return nil
	},
}

func init() {
	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
	cmdTimes.Flags().BoolVarP(&kaiguan, "kaiguan", "k", false, "zhege shi kaiguan")
	//times是echo的子目录
	cmdEcho.AddCommand(cmdTimes)
}
