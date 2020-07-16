package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Tone string
var Message string

func init() {
	rootCmd.AddCommand(versionCmd, statementCmd)
	statementCmd.Flags().StringVarP(&Message, "tone", "t", "", "tone of the message")
	// statementCmd.MarkFlagRequired("tone")
	// viper.BindPFlag("tone", statementCmd.PersistentFlags().Lookup("tone"))
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cobraex",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cobraex v0.1 -- HEAD")
	},
}

var statementCmd = &cobra.Command{
	Use:              "print",
	Short:            "Print the statement from cobraex",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		tStatus, _ := cmd.Flags().GetString("tone")
		b := IsValidTone(tStatus)
		if b {
			if tStatus == "positive" {
				fmt.Println("cobra is so useful and worth understanding")
			}
			fmt.Println("cobra is difficult to understand")
		} else {
			fmt.Println("huh?")
		}
		fmt.Println("cobra is a tool")
	},
}

var toneCmd = &cobra.Command{
	Use:   "tone",
	Short: "tone of the message [positive/negative]",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(args)
		return nil
	},
}

func IsValidTone(arg string) bool {
	if arg == "positive" || arg == "negative" {
		return true
	}
	return false
}
