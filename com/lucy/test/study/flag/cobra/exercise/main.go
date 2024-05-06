package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "myapp",
		Short: "我的应用程序",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			fmt.Printf("您好，%s！\n", name)
		},
	}

	rootCmd.Flags().String("name", "", "姓名")

	rootCmd.Flags().MarkHidden("age")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

// 忽略未定义的标识参数
func ignoreUndefinedFlags(cmd *cobra.Command, args []string) {
	unknownFlags := make([]string, 0)

	for _, arg := range args {
		isFlagDefined := false
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			if f.Name == arg || f.Shorthand == arg {
				isFlagDefined = true
			}
		})

		if !isFlagDefined {
			unknownFlags = append(unknownFlags, arg)
		}
	}

	for _, flag := range unknownFlags {
		cmd.Flags().MarkHidden(flag)
	}
}
