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
	Run:   listRun}

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
