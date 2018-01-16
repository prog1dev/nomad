package fingerprint

import (
	"testing"

	"github.com/hashicorp/nomad/client/config"
	"github.com/hashicorp/nomad/nomad/structs"
)

func TestMemoryFingerprint(t *testing.T) {
	f := NewMemoryFingerprint(testLogger())
	node := &structs.Node{
		Attributes: make(map[string]string),
	}
	nodeAttributesDiff, err := f.Fingerprint(&config.Config{}, node)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if len(nodeAttributesDiff) == 0 {
		t.Fatalf("should generate diff of node attributes")
	}

	assertNodeAttributeContains(t, nodeAttributesDiff, "memory.totalbytes")

	if node.Resources == nil {
		t.Fatalf("Node Resources was nil")
	}
	if node.Resources.MemoryMB == 0 {
		t.Errorf("Expected node.Resources.MemoryMB to be non-zero")
	}

}
