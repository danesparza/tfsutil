package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy a variable group",
	Long:  `Copies a variable group and all its variables to a new variable group.`,
	Run:   vgcopy,
}

func vgcopy(cmd *cobra.Command, args []string) {
	log.Print("[INFO] copyvg called")
}

func init() {
	vgCmd.AddCommand(copyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// copyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// copyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
