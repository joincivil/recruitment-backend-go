package metadata_test

import (
	"testing"

	"github.com/joincivil/recruitment-backend-go/pkg/metadata"
)

func isRetriever(t *testing.T, r metadata.Retriever) {
}

func TestHttpRetriever(t *testing.T) {
	r := metadata.HTTPRetriever{}

	meta, err := r.Revisions(metadata.SludgeName)
	if err != nil {
		t.Errorf("Should have retrieved some metadata: err: %v", err)
	}

	if len(meta) != 3 {
		t.Errorf("Should have retrieved 3 items from the metadata")
	}

	if meta[0].Version == 0 {
		t.Errorf("Should have retrieved a valid value for version")
	}

	if meta[1].Version == 1 {
		t.Errorf("Should have retrieved a valid value for version")
	}

	if meta[0].Mission == "" {
		t.Errorf("Should have retrieved a value for mission statement")
	}

	if meta[0].Ownership == "" {
		t.Errorf("Should have retrieved a value for ownership statement")
	}

	if meta[0].Revenue == "" {
		t.Errorf("Should have retrieved a value for revenue statement")
	}

	if meta[0].Conflicts == "" {
		t.Errorf("Should have retrieved a value for conflict statement")
	}

	meta, err = r.Revisions(metadata.ColoradoSunName)
	if err != nil {
		t.Errorf("Should have retrieved some metadata: err: %v", err)
	}

	if len(meta) != 3 {
		t.Errorf("Should have retrieved 3 items from the metadata")
	}

	meta, err = r.Revisions(metadata.EcowurdName)
	if err != nil {
		t.Errorf("Should have retrieved some metadata: err: %v", err)
	}

	if len(meta) != 3 {
		t.Errorf("Should have retrieved 3 items from the metadata")
	}
}
