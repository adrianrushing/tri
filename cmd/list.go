/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/adrianrushing/tri/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List todos",
	Long:  `Listing the todos`,
	Run:   listRun,
}

func listRun(cmd *cobra.Command, args []string) {

	items, err := todo.ReadItems(dataFile)

	if err != nil {
		log.Printf("%v", err)
	}
	sort.Sort(todo.ByPri(items))
	w := tabwriter.NewWriter(os.Stdout, 3, -0, 1, ' ', 0)

	for _, i := range items {
		fmt.Fprintln(w, i.PrettyP()+"\t"+i.Text+"\t")
	}

	w.Flush()

}

func init() {
	rootCmd.AddCommand(listCmd)

}
