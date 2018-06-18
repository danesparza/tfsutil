package tfs

import (
	"encoding/base64"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

// GetAPIResponse gets an API response for the given url request
func GetAPIResponse(url string) (*http.Response, error) {
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
