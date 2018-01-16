package fingerprint

import (
	"testing"

	"github.com/hashicorp/nomad/client/config"
	"github.com/hashicorp/nomad/nomad/structs"
)

func TestHostFingerprint(t *testing.T) {
	f := NewHostFingerprint(testLogger())
	node := &structs.Node{
		Attributes: make(map[string]string),
	}
	nodeAttributesDiff, err := f.Fingerprint(&config.Config{}, node)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if len(nodeAttributesDiff) == 0 {
		t.Fatalf("should generate a diff of node attributes")
	}

	// Host info
	for _, key := range []string{"os.name", "os.version", "unique.hostname", "kernel.name"} {
		assertNodeAttributeContains(t, nodeAttributesDiff, key)
	}
}
