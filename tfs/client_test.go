package tfs_test

import (
	"testing"

	"github.com/danesparza/tfsutil/tfs"
)

//	Arrange our test table
var urltests = []struct {
	url        string
	collection string
	project    string
	expected   string
}{
	{"http://tfsrepository.mydomain.com:8080/tfs/", "", "", "http://tfsrepository.mydomain.com:8080/tfs/defcol/defproj/_apis"},
	{"http://tfsrepository.mydomain.com:8080/tfs/", "colone", "projone", "http://tfsrepository.mydomain.com:8080/tfs/colone/projone/_apis"},
	{"http://tfsrepository.mydomain.com:8080/tfs", "colone", "projone", "http://tfsrepository.mydomain.com:8080/tfs/colone/projone/_apis"},
	{"http://tfsrepository.mydomain.com:8080/tfs/", "col-one", "proj-one", "http://tfsrepository.mydomain.com:8080/tfs/col-one/proj-one/_apis"},
	{"http://simpleintranet:8080", "col-one", "proj-one", "http://simpleintranet:8080/col-one/proj-one/_apis"},
}

// We should be able to get a formatted url back
func TestClient_ValidDefaults_GetFormattedBaseURL_ReturnsFormattedUrl(t *testing.T) {

	//	Arrange
	client := tfs.Client{
		TfsURL:            "defaulturl",
		DefaultCollection: "defcol",
		DefaultProject:    "defproj",
	}

	//	Act
	for _, tt := range urltests {
		//	Set the base url on the client:
		client.TfsURL = tt.url

		//	Call the method with the test parameters
		actual, err := client.GetFormattedBaseURL(tt.collection, tt.project)
		if err != nil {
			t.Errorf("GetFormattedBaseUrl('%s', '%s') with base url: %s expected: %s but got error %s", tt.collection, tt.project, tt.url, tt.expected, err)
		}

		//	Compare expected with actual and report an error if they don't match
		if tt.expected != actual {
			t.Errorf("GetFormattedBaseUrl('%s', '%s') with base url: %s expected: %s but got %s", tt.collection, tt.project, tt.url, tt.expected, actual)
		}
	}

}

// If we use bad defaults and don't pass args, it should throw an error
func TestClient_BadDefaultsNoArgs_GetFormattedBaseURL_ThowsError(t *testing.T) {

	//	Arrange
	client := tfs.Client{
		TfsURL:            "http://tfsrepository.mydomain.com:8080/tfs/",
		DefaultCollection: "",
		DefaultProject:    "",
	}

	//	Act
	_, err := client.GetFormattedBaseURL("", "")

	//	Assert
	if err == nil {
		t.Errorf("GetFormattedBaseUrl with no defaults and no parameters should throw an error, but didn't")
	}

}

// If we use bad defaults but have valid args, it should return a formatted url
func TestClient_BadDefaultsValidArgs_GetFormattedBaseURL_ThowsError(t *testing.T) {

	//	Arrange
	client := tfs.Client{
		TfsURL:            "http://tfsrepository.mydomain.com:8080/tfs",
		DefaultCollection: "",
		DefaultProject:    "",
	}

	collection := "awesomecol"
	project := "awesomeproj"
	expected := "http://tfsrepository.mydomain.com:8080/tfs/awesomecol/awesomeproj/_apis"

	//	Act
	actual, err := client.GetFormattedBaseURL(collection, project)

	//	Assert
	if err != nil {
		t.Errorf("GetFormattedBaseUrl('%s', '%s') with base url: %s expected: %s but got error %s", collection, project, client.TfsURL, expected, err)
	}

	if expected != actual {
		t.Errorf("GetFormattedBaseUrl('%s', '%s') with base url: %s expected: %s but got %s", collection, project, client.TfsURL, expected, actual)
	}

}
