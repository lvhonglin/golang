package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "app [command]",
	Run: func(cmd *cobra.Command, args []string) {
		println("root")
	},
}

func Execute() {
	rootCmd.Execute()
}
