package fingerprint

import (
	"testing"

	"github.com/hashicorp/nomad/client/config"
	"github.com/hashicorp/nomad/nomad/structs"
)

func TestArchFingerprint(t *testing.T) {
	f := NewArchFingerprint(testLogger())
	node := &structs.Node{
		Attributes: make(map[string]string),
	}
	nodeAttributesDiff, err := f.Fingerprint(&config.Config{}, node)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if nodeAttributesDiff["cpu.arch"] == "" {
		t.Fatalf("missing arch")
	}
}
