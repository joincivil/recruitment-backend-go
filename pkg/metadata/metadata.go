package metadata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// SludgeName is the string that represents Sludge
	SludgeName = "sludge"
	// ColoradoSunName is the string that represents the ColoradoSun
	ColoradoSunName = "coloradosun"
	// EcowurdName is the string that represents Ecowurd
	EcowurdName = "ecowurd"
)

var (
	revisionNameToURLs = map[string]string{
		SludgeName:      "https://storage.googleapis.com/civil-media-misc/code-test/data/sludge.json",
		ColoradoSunName: "https://storage.googleapis.com/civil-media-misc/code-test/data/coloradosun.json",
		EcowurdName:     "https://storage.googleapis.com/civil-media-misc/code-test/data/ecowurd.json",
	}
)

// Retriever is an interface for retrieval of metadata from
// a source
type Retriever interface {
	Revisions(newsroomName string) ([]*Metadata, error)
}

// Metadata represents a metadata revision for a newsroom
type Metadata struct {
	Version   int    `json:"version"`
	Mission   string `json:"mission"`
	Ownership string `json:"ownership"`
	Revenue   string `json:"revenue"`
	Conflicts string `json:"conflicts"`
}

// HTTPResponse represents the data format that wraps the metadata
// revisions
type HTTPResponse struct {
	Name      string      `json:"name"`
	Revisions []*Metadata `json:"revisions"`
}

// HTTPRetriever implements Retriever to fetch metadata
// contents from an HTTP source
type HTTPRetriever struct {
}

// Revisions returns the feed for a newsroom given the newsroom name
func (r *HTTPRetriever) Revisions(newsroomName string) ([]*Metadata, error) {
	client := http.Client{}

	url, ok := revisionNameToURLs[newsroomName]
	if !ok {
		return nil, fmt.Errorf("newsroom not found")
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	revsResp := &HTTPResponse{}
	err = json.Unmarshal(body, &revsResp)
	if err != nil {
		return nil, err
	}

	return revsResp.Revisions, nil
}
