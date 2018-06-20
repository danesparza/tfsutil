package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
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
	log.Printf("We will copy the group '%s'", args[0])

	//	See if we can find that group

	//	If we can, compose a new request and attempt to add it:

}

func init() {
	vgCmd.AddCommand(copyCmd)
}
