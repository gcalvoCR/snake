package cmd

import (
	"fmt"

	"github.com/gcalvoCR/snake/util"
	"github.com/spf13/cobra"
)

var snakeCmd = &cobra.Command{
	Use:   "run",
	Short: "Let's call the snake!",
	Long: `Let's run some scripts here`,
	Version: "1.0.0",
	Aliases: []string{"go", "saveme"},
	Run: func(cmd *cobra.Command, args []string) {
		
		// Validate time flag
		time, err := cmd.Flags().GetInt("time")
		if err !=nil {
			fmt.Println(err.Error())
			fmt.Println("Come on, provide a number, you can do it!\nThe format should be similar to the following:\n --time=10")
			return
		}

		// Validate units flag
		u := []string { "ns", "us", "ms", "s", "m", "h" } 
		units, err := cmd.Flags().GetString("units")
		invalidUnitsMessage := "Provide a valid unit!\nThe format should be similar to the following\n --units=m\nValid units: ns, us, ms, s, m, h\n"
		
		if err !=nil {
			fmt.Println(invalidUnitsMessage)
			return
		}

		isUnitOK := util.Contains(u,units)

		if !isUnitOK {
			fmt.Println(invalidUnitsMessage)
			return
		}

		// Process request
		fmt.Println("Great! Let's do this")
		util.MyFunc(fmt.Sprintf("%d%s", time, units))
	},
}

var Source string

func init() {
	rootCmd.AddCommand(snakeCmd)
	snakeCmd.PersistentFlags().Int("time", 30, "Set the running time.")
	snakeCmd.PersistentFlags().String("units", "s", "Set the units (ns, us, ms, s, m, h)")
}
