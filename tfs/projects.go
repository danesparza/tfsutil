package tfs

// ProjectResponse defines the response recieved when querying projects
type ProjectResponse struct {
	Count    int       `json:"count"`
	Projects []Project `json:"value"`
}

// Project is a single TFS project
type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	State       string `json:"state"`
	Revision    int    `json:"revision"`
	Visibility  string `json:"visibility"`
}
