package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"

	"github.com/danesparza/tfsutil/tfs"
	"github.com/rs/xid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:   "copy \"<name>\"",
	Short: "Copy a variable group",
	Long: `Copies a variable group and all its variables to a new variable group.
	
NOTE: For variable group names that contain spaces, remember to surround the group name with quotes.  

Example: 
tfsutil vg copy "Test group name"

`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires a variable group name")
		}
		return nil
	},
	Run: vgcopy,
}

func vgcopy(cmd *cobra.Command, args []string) {

	//	First, see what group should be copied
	groupName := args[0]
	log.Printf("[DEBUG] Attempting to copy the group '%s'", groupName)

	//	See if we can find that group by attempting to create a group filter
	groupFilter := fmt.Sprintf("distributedtask/variablegroups?groupName=%s&actionFilter=use&top=50&api-version=4.1-preview.1", url.QueryEscape(groupName))
	listfullurl := fmt.Sprintf("%s/%s", viper.GetString("tfsurl"), groupFilter)

	//	Request a list of variable groups
	resp1, err := tfs.GetAPIResponse(listfullurl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp1.Body.Close()

	//	If the HTTP status code indicates an error, report it and get out
	if resp1.StatusCode >= 400 {
		log.Fatalln("[ERROR] There was an error getting information from TFS: ", resp1.Status)
	}

	//	Decode the return object
	vgroups := VariableGroupsResponse{}
	err = json.NewDecoder(resp1.Body).Decode(&vgroups)
	if err != nil {
		log.Fatal(err)
	}

	//	Did we find it?
	if vgroups.Count < 1 {
		log.Fatalf("Sorry -- I couldn't find the group '%s'", groupName)
	}

	if vgroups.Count > 1 && vgroups.VariableGroups[0].Name != groupName {
		log.Fatalf("Sorry -- Too many groups match '%s' -- please be more specific", groupName)
	}

	//	If we did, see if it has items:
	log.Printf("[DEBUG] Copying '%s' (and %v variables)", vgroups.VariableGroups[0].Name, len(vgroups.VariableGroups[0].Variables))

	//	If we can find it, compose a new request and attempt to add it:
	variableGroupCopy := VariableGroup{}
	variableGroupCopy.Description = vgroups.VariableGroups[0].Description
	variableGroupCopy.Type = vgroups.VariableGroups[0].Type
	variableGroupCopy.Variables = vgroups.VariableGroups[0].Variables

	//	Make the name a bit unique
	guid := xid.New()
	variableGroupCopy.Name = fmt.Sprintf("Copy of %s (%s)", vgroups.VariableGroups[0].Name, guid.String())
	log.Printf("[DEBUG] Creating a group with the name: %s", variableGroupCopy.Name)

	//	Prepare the request body
	requestBytes := new(bytes.Buffer)
	err = json.NewEncoder(requestBytes).Encode(&variableGroupCopy)
	if err != nil {
		log.Fatalf("Sorry -- There was a problem preparing to copy the group '%s'", groupName)
	}

	//	Send the request to the API:
	addfullurl := fmt.Sprintf("%s/distributedtask/variablegroups?api-version=4.1-preview.1", viper.GetString("tfsurl"))
	resp2, err := tfs.PostAPIResponse(addfullurl, requestBytes.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp2.Body.Close()

	//	If the HTTP status code indicates an error, report it and get out
	if resp2.StatusCode >= 400 {
		log.Fatalln("[ERROR] There was an error copying information in TFS: ", resp2.Status)
	}

	fmt.Printf("\nCopied \n %s \nto \n %s \n (including %v variables)", groupName, variableGroupCopy.Name, len(variableGroupCopy.Variables))
}

func init() {
	vgCmd.AddCommand(copyCmd)
}
