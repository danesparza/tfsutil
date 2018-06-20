package cmd

import (
	"sort"
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
	Count          int             `json:"count"`
	VariableGroups []VariableGroup `json:"value"`
}

// VariableGroup is a single variable group
type VariableGroup struct {
	Variables map[string]Variable `json:"variables"`
	ID        int                 `json:"id,omitempty"`
	Type      string              `json:"type,omitempty"`
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
	} `json:"createdBy,omitempty"`

	CreatedOn time.Time `json:"createdOn,omitempty"`

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
	} `json:"modifiedBy,omitempty"`

	ModifiedOn  time.Time `json:"modifiedOn,omitempty"`
	Description string    `json:"description"`
}

// Variable defines a single variable in a variable group
type Variable struct {
	Value string `json:"value"`
}

// By is the type of a "less" function that defines the ordering of its VariableGroup arguments.
type By func(p1, p2 *VariableGroup) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(groups []VariableGroup) {
	ps := &vgSorter{
		groups: groups,
		by:     by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// vgSorter joins a By function and a slice of VariableGroups to be sorted.
type vgSorter struct {
	groups []VariableGroup
	by     func(p1, p2 *VariableGroup) bool // Closure used in the Less method.
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
