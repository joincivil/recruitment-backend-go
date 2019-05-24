package metadata

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

// TODO: Implement an HTTP retriever to retrieve the revision data
