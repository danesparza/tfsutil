package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var yamlDefault = []byte(`# Config created %s
tfsurl: http://YOURSERVER:8080/tfs
pat: YOUR_PERSONAL_ACCESS_TOKEN
collection: OPTIONAL_DEFAULT_COLLECTION
project: OPTIONAL_DEFAULT_PROJECT
`)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Renders a config file",
	Long:  `Outputs a config file with default values.  Write this to a file called 'tfsutil.yml' and customize the values.`,
	Run: func(cmd *cobra.Command, args []string) {
		t := time.Now()
		fmt.Printf(string(yamlDefault), t.Format(time.RFC3339))
	},
}

func init() {
	configCmd.AddCommand(createCmd)
}
