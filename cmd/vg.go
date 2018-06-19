package cmd

import (
	"time"

	"github.com/spf13/cobra"
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
}

func init() {
	rootCmd.AddCommand(vgCmd)
}

// VariableGroupsResponse defines the response recieved when querying variable groups
type VariableGroupsResponse struct {
	Count int `json:"count"`
	Value []struct {
		Variables map[string]Variable `json:"variables"`
		ID        int                 `json:"id"`
		Type      string              `json:"type"`
		Name      string              `json:"name"`

		CreatedBy struct {
			DisplayName string `json:"displayName"`
			URL         string `json:"url"`
			Links       struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"_links"`
			ID         string `json:"id"`
			UniqueName string `json:"uniqueName"`
			ImageURL   string `json:"imageUrl"`
		} `json:"createdBy"`

		CreatedOn time.Time `json:"createdOn"`

		ModifiedBy struct {
			DisplayName string `json:"displayName"`
			URL         string `json:"url"`
			Links       struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"_links"`
			ID         string `json:"id"`
			UniqueName string `json:"uniqueName"`
			ImageURL   string `json:"imageUrl"`
		} `json:"modifiedBy"`

		ModifiedOn  time.Time `json:"modifiedOn"`
		Description string    `json:"description"`
	} `json:"value"`
}

// Variable defines a single variable in a variable group
type Variable struct {
	Value string `json:"value"`
}
