package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var yamlDefault = []byte(`# Default config
tfsurl: http://YOURSERVER:8080/tfs/YOURCOLLECTION/YOURPROJECT/_apis
pat: YOUR_PERSONAL_ACCESS_TOKEN
`)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Renders a config file",
	Long:  `Outputs a config file with default values.  Write this to a file called 'tfsutil.yml' and customize the values.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s", yamlDefault)
	},
}

func init() {
	configCmd.AddCommand(createCmd)
}
