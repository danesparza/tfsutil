package tfs

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/spf13/viper"
)

// Client is a TFS client
type Client struct {
	TfsURL            string
	DefaultCollection string
}

// GetFormattedURL gets the formatted TFS url to use
func (client Client) GetFormattedURL(collection, project, area, resource, query string) (string, error) {

	//	If the collection is not blank, use it.  Otherwise, use defaults
	urlcol := client.DefaultCollection
	if collection != "" {
		urlcol = collection
	}

	//	If urlcol is blank, we have a problem.  Return an error:
	if urlcol == "" {
		return "", errors.New("TFS collection isn't specified, but is required")
	}

	//	Parse the base url
	u, err := url.Parse(client.TfsURL)
	if err != nil {
		return "", err
	}

	//	Assemble the component parts
	u.Path = path.Join(u.Path, urlcol, project, "_apis", area, resource)

	//	Add the querystring
	u.RawQuery = query

	//	Return the full url
	return u.String(), nil
}

// GetListOfProjects gets a list of projects for the given collection
func (client Client) GetListOfProjects(collection string) (ProjectResponse, error) {

	//	Our return value:
	retval := ProjectResponse{}

	//	Format the url
	fullurl, err := client.GetFormattedURL(collection, "", "", "projects", "")
	if err != nil {
		apperr := fmt.Errorf("Unable to format url: %s", err)
		return retval, apperr
	}

	//	Make a GET reqeust to TFS
	resp, err := getAPIResponse(fullurl)
	if err != nil {
		apperr := fmt.Errorf("There was a problem calling TFS: %s", err)
		return retval, apperr
	}
	defer resp.Body.Close()

	//	If the HTTP status code indicates an error, report it and get out
	if resp.StatusCode >= 400 {
		apperr := fmt.Errorf("There was an error getting information from TFS: %s", resp.Status)
		return retval, apperr
	}

	//	Decode the return object
	err = json.NewDecoder(resp.Body).Decode(&retval)
	if err != nil {
		apperr := fmt.Errorf("There was a problem decoding the response from TFS: %s", err)
		return retval, apperr
	}

	return retval, nil
}

// GetListOfVariableGroups gets a list of variable groups for the given collection and project
func (client Client) GetListOfVariableGroups(collection, project string) (VariableGroupsResponse, error) {

	//	Our return value:
	retval := VariableGroupsResponse{}

	//	Format the url
	fullurl, err := client.GetFormattedURL(collection, project, "distributedtask", "variablegroups", "groupName=*&actionFilter=use&top=50&api-version=4.1-preview.1")
	if err != nil {
		apperr := fmt.Errorf("Unable to format url: %s", err)
		return retval, apperr
	}

	//	Request a list of variable groups
	resp, err := getAPIResponse(fullurl)
	if err != nil {
		apperr := fmt.Errorf("There was a problem calling TFS: %s", err)
		return retval, apperr
	}
	defer resp.Body.Close()

	//	If the HTTP status code indicates an error, report it and get out
	if resp.StatusCode >= 400 {
		apperr := fmt.Errorf("There was an error getting information from TFS: %s", resp.Status)
		return retval, apperr
	}

	//	Decode the return object
	err = json.NewDecoder(resp.Body).Decode(&retval)
	if err != nil {
		apperr := fmt.Errorf("There was a problem decoding the response from TFS: %s", err)
		return retval, apperr
	}

	return retval, nil
}

// GetListOfMatchingVariableGroups gets a list of variable groups for the given collection, project, and group name
func (client Client) GetListOfMatchingVariableGroups(collection, project, groupName string) (VariableGroupsResponse, error) {

	//	Our return value:
	retval := VariableGroupsResponse{}

	//	Format the url
	escapedGroup := url.QueryEscape(groupName)
	formattedQuery := fmt.Sprintf("groupName=%s&actionFilter=use&top=50&api-version=4.1-preview.1", escapedGroup)
	fullurl, err := client.GetFormattedURL(collection, project, "distributedtask", "variablegroups", formattedQuery)
	if err != nil {
		apperr := fmt.Errorf("Unable to format url: %s", err)
		return retval, apperr
	}

	//	Request a list of variable groups
	resp, err := getAPIResponse(fullurl)
	if err != nil {
		apperr := fmt.Errorf("There was a problem calling TFS: %s", err)
		return retval, apperr
	}
	defer resp.Body.Close()

	//	If the HTTP status code indicates an error, report it and get out
	if resp.StatusCode >= 400 {
		apperr := fmt.Errorf("There was an error getting information from TFS: %s", resp.Status)
		return retval, apperr
	}

	//	Decode the return object
	err = json.NewDecoder(resp.Body).Decode(&retval)
	if err != nil {
		apperr := fmt.Errorf("There was a problem decoding the response from TFS: %s", err)
		return retval, apperr
	}

	return retval, nil
}

// CreateVariableGroup creates a variable group in the given collection and project
func (client Client) CreateVariableGroup(collection, project string, newGroup VariableGroup) error {

	//	Prepare the request body
	requestBytes := new(bytes.Buffer)
	err := json.NewEncoder(requestBytes).Encode(&newGroup)
	if err != nil {
		apperr := fmt.Errorf("There was a problem preparing to create the group: %s", err)
		return apperr
	}

	//	Format the url
	fullurl, err := client.GetFormattedURL(collection, project, "distributedtask", "variablegroups", "api-version=4.1-preview.1")
	if err != nil {
		apperr := fmt.Errorf("Unable to format url: %s", err)
		return apperr
	}

	//	Send the request to the API:
	resp, err := postAPIResponse(fullurl, requestBytes.String())
	if err != nil {
		apperr := fmt.Errorf("There was a problem calling TFS: %s", err)
		return apperr
	}
	defer resp.Body.Close()

	//	If the HTTP status code indicates an error, report it and get out
	if resp.StatusCode >= 400 {
		apperr := fmt.Errorf("There was an error getting information from TFS: %s", resp.Status)
		return apperr
	}

	return nil
}

// GetAPIResponse gets an API response for the given url request
func getAPIResponse(url string) (*http.Response, error) {
	log.Println("[DEBUG] Creating a request for ", url)

	//	Create our http client
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	//	Create our request:
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	//	Set our basic auth field:
	log.Println("[DEBUG] Using PAT ", viper.GetString("pat"))
	req.Header.Add("Authorization", "Basic "+basicAuth("", viper.GetString("pat")))

	//	Execute our request:
	return client.Do(req)
}

// PostAPIResponse POSTs to the API and then gets an API response for the given url request and JSON body
func postAPIResponse(url, jsonBody string) (*http.Response, error) {
	log.Printf("[DEBUG] Creating a POST request for %s\n with post body:\n%s\n", url, jsonBody)

	//	Create our http client
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	//	Create our request:
	req, err := http.NewRequest("POST", url, strings.NewReader(jsonBody))
	if err != nil {
		log.Fatal(err)
	}

	//	Set the request content type:
	req.Header.Add("Content-Type", "application/json")

	//	Set our basic auth field:
	log.Println("[DEBUG] Using PAT ", viper.GetString("pat"))
	req.Header.Add("Authorization", "Basic "+basicAuth("", viper.GetString("pat")))

	//	Execute our request:
	return client.Do(req)
}

//	The redirect policy func
func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", "Basic "+basicAuth("", viper.GetString("pat")))
	return nil
}

//	Format our basic auth header
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
