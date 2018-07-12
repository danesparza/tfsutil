package cmd

import (
	"fmt"
	"log"
	"sort"

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

	//	Create a client with our base TFS url
	client := tfs.Client{
		TfsURL: viper.GetString("tfsurl"),
	}

	//	Get the list of Variable groups.  Report any errors
	retval, err := client.GetListOfVariableGroups(viper.GetString("collection"), viper.GetString("project"))
	if err != nil {
		log.Fatalln("[ERROR] Variable group list \n", err)
	}

	// Closure(s) that orders the VariableGroup structure.
	name := func(p1, p2 *tfs.VariableGroup) bool {
		return p1.Name < p2.Name
	}

	//	Sort the variable groups
	VGBy(name).Sort(retval.VariableGroups)

	//	Begin the report:
	fmt.Printf("\nCollection: %v", viper.GetString("collection"))
	fmt.Printf("\nProject: %v\n", viper.GetString("project"))
	fmt.Printf("\nVariable groups found: %v\n==========================\n", retval.Count)

	//	List all the groups (and their variable counts):
	for _, group := range retval.VariableGroups {
		fmt.Printf("%s (%v variables)\n", group.Name, len(group.Variables))
	}

}

func init() {
	vgCmd.AddCommand(listCmd)
}

// VGBy is the type of a "less" function that defines the ordering of its VariableGroup arguments.
type VGBy func(p1, p2 *tfs.VariableGroup) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by VGBy) Sort(groups []tfs.VariableGroup) {
	ps := &vgSorter{
		groups: groups,
		by:     by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// vgSorter joins a By function and a slice of VariableGroups to be sorted.
type vgSorter struct {
	groups []tfs.VariableGroup
	by     func(p1, p2 *tfs.VariableGroup) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *vgSorter) Len() int {
	return len(s.groups)
}

// Swap is part of sort.Interface.
func (s *vgSorter) Swap(i, j int) {
	s.groups[i], s.groups[j] = s.groups[j], s.groups[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *vgSorter) Less(i, j int) bool {
	return s.by(&s.groups[i], &s.groups[j])
}
