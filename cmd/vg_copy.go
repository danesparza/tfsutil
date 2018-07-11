package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/rs/xid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/danesparza/tfsutil/tfs"
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

	//	Create a client with our base TFS url
	client := tfs.Client{
		TfsURL: viper.GetString("tfsurl"),
	}

	//	Get the list of Variable groups.  Report any errors
	vgroups, err := client.GetListOfMatchingVariableGroups(viper.GetString("collection"), viper.GetString("project"), groupName)
	if err != nil {
		log.Fatalln("[ERROR] Finding existing group \n", err)
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
	variableGroupCopy := tfs.VariableGroup{}
	variableGroupCopy.Description = vgroups.VariableGroups[0].Description
	variableGroupCopy.Type = vgroups.VariableGroups[0].Type
	variableGroupCopy.Variables = vgroups.VariableGroups[0].Variables

	//	Make the name a bit unique
	guid := xid.New()
	variableGroupCopy.Name = fmt.Sprintf("Copy of %s (%s)", vgroups.VariableGroups[0].Name, guid.String())
	log.Printf("[DEBUG] Creating a group with the name: %s", variableGroupCopy.Name)

	//	Create a copy of the group.  Report any errors
	err = client.CreateVariableGroup(viper.GetString("collection"), viper.GetString("project"), variableGroupCopy)
	if err != nil {
		log.Fatalf("[ERROR] Copying the group %s - \n %s", groupName, err)
	}

	fmt.Printf("\nCopied \n %s \nto \n %s \n (including %v variables)", groupName, variableGroupCopy.Name, len(variableGroupCopy.Variables))

}

func init() {
	vgCmd.AddCommand(copyCmd)
}
