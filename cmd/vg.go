package cmd

import (
	"github.com/spf13/cobra"
)

// vgCmd represents the vg command
var vgCmd = &cobra.Command{
	Use:   "vg",
	Short: "Variable group helpers",
	Long:  `Operations to help with variable groups.  You can list them and copy them`,
	Run: func(cmd *cobra.Command, args []string) {
		//	This command on it's own should just show
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(vgCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vgCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vgCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
