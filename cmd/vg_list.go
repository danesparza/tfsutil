package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/danesparza/tfsutil/tfs"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List variable groups",
	Long:  `Lists variable groups`,
	Run:   vglist,
}

func vglist(cmd *cobra.Command, args []string) {

	fullurl := fmt.Sprintf("%s/%s", viper.GetString("tfsurl"), "distributedtask/variablegroups?groupName=*&actionFilter=use&top=50&api-version=4.1-preview.1")

	//	Request a list of variable groups
	resp, err := tfs.GetAPIResponse(fullurl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//	If the HTTP status code indicates an error, report it and get out
	if resp.StatusCode >= 400 {
		log.Fatalln("[ERROR] There was an error getting information from TFS")
	}

	//	Decode the return object
	retval := VariableGroupsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&retval)
	if err != nil {
		log.Fatal(err)
	}

	// Closure(s) that order the VariableGroup structure.
	name := func(p1, p2 *VariableGroup) bool {
		return p1.Name < p2.Name
	}

	//	Sort the variable groups
	By(name).Sort(retval.VariableGroups)

	//	Begin the report:
	fmt.Printf("\nGroups found: %v\n================\n", retval.Count)

	//	List all the groups (and their variable counts):
	for _, group := range retval.VariableGroups {
		fmt.Printf("%s (%v variables)\n", group.Name, len(group.Variables))
	}

}

func init() {
	vgCmd.AddCommand(listCmd)
}
