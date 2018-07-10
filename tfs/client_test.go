package tfs_test

import (
	"testing"

	"github.com/danesparza/tfsutil/tfs"
)

type testparams struct {
	collection string
	project    string
	expected   string
}

//	Our test data struct
type testdata struct {
	baseurl        string
	defaultcol     string
	defaultproject string
	params         []testparams
}

// If we use valid defaults, We should be able to get a formatted url back
func TestClient_ValidDefaults_GetFormattedBaseURL_ReturnsFormattedUrl(t *testing.T) {

	//	Arrange
	urltests := testdata{
		baseurl:        "http://tfsrepository.mydomain.com:8080/tfs",
		defaultcol:     "DefaultCollection",
		defaultproject: "",
		params: []testparams{
			{"", "", "http://tfsrepository.mydomain.com:8080/tfs/DefaultCollection/_apis"},
			{"colone", "projone", "http://tfsrepository.mydomain.com:8080/tfs/colone/projone/_apis"},
			{"colone", "projone", "http://tfsrepository.mydomain.com:8080/tfs/colone/projone/_apis"},
			{"col-one", "proj-one", "http://tfsrepository.mydomain.com:8080/tfs/col-one/proj-one/_apis"},
		},
	}

	client := tfs.Client{
		TfsURL:            urltests.baseurl,
		DefaultCollection: urltests.defaultcol,
		DefaultProject:    urltests.defaultproject,
	}

	//	Act
	for _, tt := range urltests.params {

		//	Call the method with the test parameters
		actual, err := client.GetFormattedBaseURL(tt.collection, tt.project)
		if err != nil {
			t.Errorf("GetFormattedBaseUrl('%s', '%s') with base url: %s expected: %s but got error %s", tt.collection, tt.project, urltests.baseurl, tt.expected, err)
		}

		//	Compare expected with actual and report an error if they don't match
		if tt.expected != actual {
			t.Errorf("GetFormattedBaseUrl('%s', '%s') with base url: %s expected: %s but got %s", tt.collection, tt.project, urltests.baseurl, tt.expected, actual)
		}
	}

}

// If we use blank defaults and don't pass args, it should throw an error
func TestClient_BlankDefaultsNoArgs_GetFormattedBaseURL_ThowsError(t *testing.T) {

	//	Arrange
	client := tfs.Client{
		TfsURL:            "http://tfsrepository.mydomain.com:8080/tfs/",
		DefaultCollection: "",
		DefaultProject:    "",
	}

	collection := ""
	project := ""

	//	Act
	_, err := client.GetFormattedBaseURL(collection, project)

	//	Assert
	if err == nil {
		t.Errorf("GetFormattedBaseUrl with no defaults and no parameters should throw an error, but didn't")
	}

}

// If we use blank defaults but have valid args, it should return a formatted url
func TestClient_BlankDefaultsValidArgs_GetFormattedBaseURL_ReturnsFormattedUrl(t *testing.T) {

	//	Arrange
	urltests := testdata{
		baseurl:        "http://tfsrepository.mydomain.com:8080/tfs",
		defaultcol:     "",
		defaultproject: "",
		params: []testparams{
			{"colone", "projone", "http://tfsrepository.mydomain.com:8080/tfs/colone/projone/_apis"},
			{"colone", "projone", "http://tfsrepository.mydomain.com:8080/tfs/colone/projone/_apis"},
			{"col-one", "proj-one", "http://tfsrepository.mydomain.com:8080/tfs/col-one/proj-one/_apis"},
		},
	}

	client := tfs.Client{
		TfsURL:            urltests.baseurl,
		DefaultCollection: urltests.defaultcol,
		DefaultProject:    urltests.defaultproject,
	}

	//	Act
	for _, tt := range urltests.params {

		//	Call the method with the test parameters
		actual, err := client.GetFormattedBaseURL(tt.collection, tt.project)
		if err != nil {
			t.Errorf("GetFormattedBaseUrl('%s', '%s') with base url: %s expected: %s but got error %s", tt.collection, tt.project, urltests.baseurl, tt.expected, err)
		}

		//	Compare expected with actual and report an error if they don't match
		if tt.expected != actual {
			t.Errorf("GetFormattedBaseUrl('%s', '%s') with base url: %s expected: %s but got %s", tt.collection, tt.project, urltests.baseurl, tt.expected, actual)
		}
	}

}
