package tfs

import (
	"encoding/base64"
	"errors"
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
	DefaultProject    string
}

// GetFormattedBaseURL gets the formatted TFS base url to use
func (client Client) GetFormattedBaseURL(collection, project string) (string, error) {

	//	If the collection or project are not blank, use them.  Otherwise, use defaults
	urlcol := client.DefaultCollection
	if collection != "" {
		urlcol = collection
	}

	urlproj := client.DefaultProject
	if project != "" {
		urlproj = project
	}

	//	If urlcol or urlproj is blank, we have a problem.  Return an error:
	if urlcol == "" {
		return "", errors.New("TFS project isn't specified, but is required")
	}

	if urlproj == "" {
		return "", errors.New("TFS collection isn't specified, but is required")
	}

	//	Parse the base url
	u, err := url.Parse(client.TfsURL)
	if err != nil {
		return "", err
	}

	//	Assemble the component parts
	u.Path = path.Join(u.Path, urlcol, urlproj, "_apis")

	//	Return the full url
	return u.String(), nil
}

// GetAPIResponse gets an API response for the given url request
func GetAPIResponse(url string) (*http.Response, error) {
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
func PostAPIResponse(url, jsonBody string) (*http.Response, error) {
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
