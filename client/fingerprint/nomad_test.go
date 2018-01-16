package fingerprint

import (
	"testing"

	"github.com/hashicorp/nomad/client/config"
	"github.com/hashicorp/nomad/nomad/structs"
	"github.com/hashicorp/nomad/version"
)

func TestNomadFingerprint(t *testing.T) {
	f := NewNomadFingerprint(testLogger())
	node := &structs.Node{
		Attributes: make(map[string]string),
	}
	v := "foo"
	r := "123"
	c := &config.Config{
		Version: &version.VersionInfo{
			Revision: r,
			Version:  v,
		},
	}
	nodeAttributesDiff, err := f.Fingerprint(c, node)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if len(nodeAttributesDiff) == 0 {
		t.Fatalf("should apply")
	}
	if nodeAttributesDiff["nomad.version"] != v {
		t.Fatalf("incorrect version")
	}
	if nodeAttributesDiff["nomad.revision"] != r {
		t.Fatalf("incorrect revision")
	}
}
