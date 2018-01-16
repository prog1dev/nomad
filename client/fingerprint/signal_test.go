package fingerprint

import (
	"testing"

	"github.com/hashicorp/nomad/nomad/structs"
)

func TestSignalFingerprint(t *testing.T) {
	fp := NewSignalFingerprint(testLogger())
	node := &structs.Node{
		Attributes: make(map[string]string),
	}

	nodeAttributesDiff := assertFingerprintOK(t, fp, node)
	assertNodeAttributeContains(t, nodeAttributesDiff, "os.signals")
}
