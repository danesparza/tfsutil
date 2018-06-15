package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// copyvgCmd represents the copyvg command
var copyvgCmd = &cobra.Command{
	Use:   "copyvg",
	Short: "Copy a variable group",
	Long:  `Copies a variable group and all its variables to a new variable group.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("copyvg called")
	},
}

func init() {
	rootCmd.AddCommand(copyvgCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// copyvgCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// copyvgCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
