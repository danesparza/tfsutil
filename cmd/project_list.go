package cmd

import (
	"fmt"
	"log"
	"sort"

	"github.com/danesparza/tfsutil/tfs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// projectListCmd represents the projectList command
var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	Long:  `List projects`,
	Run:   projectlist,
}

func projectlist(cmd *cobra.Command, args []string) {

	//	Create a client with our base TFS url
	client := tfs.Client{
		TfsURL: viper.GetString("tfsurl"),
	}

	//	Get the list of Projects.  Report any errors
	retval, err := client.GetListOfProjects(viper.GetString("collection"))
	if err != nil {
		log.Fatalln("[ERROR] Project list \n", err)
	}

	// Closure(s) that orders the VariableGroup structure.
	name := func(p1, p2 *tfs.Project) bool {
		return p1.Name < p2.Name
	}

	//	Sort the projects
	ProjectBy(name).Sort(retval.Projects)

	//	Begin the report:
	fmt.Printf("\nCollection: %v\n", viper.GetString("collection"))
	fmt.Printf("\nProjects found: %v\n====================\n", retval.Count)

	//	List all the projects:
	for _, project := range retval.Projects {
		fmt.Printf("%s\n", project.Name)
	}

}

func init() {
	projectCmd.AddCommand(projectListCmd)

}

// ProjectBy is the type of a "less" function that defines the ordering of its VariableGroup arguments.
type ProjectBy func(p1, p2 *tfs.Project) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by ProjectBy) Sort(items []tfs.Project) {
	ps := &projectSorter{
		items: items,
		by:    by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// projectSorter joins a By function and a slice of items to be sorted.
type projectSorter struct {
	items []tfs.Project
	by    func(p1, p2 *tfs.Project) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *projectSorter) Len() int {
	return len(s.items)
}

// Swap is part of sort.Interface.
func (s *projectSorter) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *projectSorter) Less(i, j int) bool {
	return s.by(&s.items[i], &s.items[j])
}
