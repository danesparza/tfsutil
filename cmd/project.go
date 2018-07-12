package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Project helpers",
	Long:  `Operations to help with projects.  You can list them`,
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
	rootCmd.AddCommand(projectCmd)
}
