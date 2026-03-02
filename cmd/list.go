/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/adrianrushing/tri/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List todos",
	Long: `A longer that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: listRun}

func listRun(cmd *cobra.Command, args []string) {

	items, err := todo.ReadItems("/home/adrian/.tridos.json")

	if err != nil {
		log.Printf("%v", err)
	}

	fmt.Println(items)

}

func init() {
	rootCmd.AddCommand(listCmd)

}
