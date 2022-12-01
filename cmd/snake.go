package cmd

import (
	"github.com/gcalvoCR/snake/util"
	"github.com/spf13/cobra"
)

var snakeCmd = &cobra.Command{
	Use:   "run",
	Short: "Let's call the snake!",
	Long: `Let's run some scripts here`,
	Version: "1.0.0",
	Aliases: []string{"go", "doit"},
	Run: func(cmd *cobra.Command, args []string) {
		util.MyFunc()
	},
}

func init() {
	rootCmd.AddCommand(snakeCmd)
}
