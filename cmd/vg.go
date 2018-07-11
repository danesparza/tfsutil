package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// vgCmd represents the variable group base command
var vgCmd = &cobra.Command{
	Use:   "vg",
	Short: "Variable group helpers",
	Long:  `Operations to help with variable groups.  You can list them and copy them`,
	Run: func(cmd *cobra.Command, args []string) {
		//	This command on it's own should just show help
		cmd.Help()
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		//	Verify that we have a tfsurl and a pat
		if strings.TrimSpace(viper.GetString("tfsurl")) == "" {
			fmt.Printf("\nThis tool requires a TFS base url to operate.   \n\nPlease specify one on the command line or in the config file 'tfsutil.yml' \nFor help creating a config file, see the command 'tfsutil config create'\n")
			os.Exit(1)
		}

		if strings.TrimSpace(viper.GetString("pat")) == "" {
			fmt.Printf("\nThis tool requires a TFS Personal Access Token (pat) for authentication.  \n\nPlease specify a pat on the command line or in the config file 'tfsutil.yml' \nFor help creating a config file, see the command 'tfsutil config create'\n")
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(vgCmd)
}
