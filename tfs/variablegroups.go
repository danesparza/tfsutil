package tfs

import (
	"time"
)

// VariableGroupsResponse defines the response recieved when querying variable groups
type VariableGroupsResponse struct {
	Count          int             `json:"count"`
	VariableGroups []VariableGroup `json:"value"`
}

// VariableGroup is a single variable group
type VariableGroup struct {
	// Variables is the map of variables associated with this VariableGroup
	Variables map[string]Variable `json:"variables"`
	ID        int                 `json:"id,omitempty"`

	// Type is the variable group type.  By default this should be 'Vsts'
	Type string `json:"type"`

	// Name is the name of the variable group
	Name string `json:"name"`

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

	ModifiedOn time.Time `json:"modifiedOn,omitempty"`

	// Description is the description for this variable group
	Description string `json:"description"`
}

// Variable defines a single variable in a variable group
type Variable struct {
	Value string `json:"value"`
}
