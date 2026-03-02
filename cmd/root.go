/*
Copyright © 2026 Adrian
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var dataFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tri",
	Short: "Tri is a todo application",
	Long: `Tri will help you get more done in less time. 
	It's designed to be as simple as possible to help
	you accomplish your goals.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	home, err := os.UserHomeDir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set it with --datafile.")
	}
	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", home+string(os.PathSeparator)+".tridos.json", "data file to store todos.")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
