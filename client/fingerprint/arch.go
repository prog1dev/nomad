package fingerprint

import (
	"log"
	"runtime"

	client "github.com/hashicorp/nomad/client/config"
	"github.com/hashicorp/nomad/nomad/structs"
)

// ArchFingerprint is used to fingerprint the architecture
type ArchFingerprint struct {
	StaticFingerprinter
	logger *log.Logger
}

// NewArchFingerprint is used to create an OS fingerprint
func NewArchFingerprint(logger *log.Logger) Fingerprint {
	f := &ArchFingerprint{logger: logger}
	return f
}

func (f *ArchFingerprint) Fingerprint(config *client.Config, node *structs.Node) (map[string]string, error) {
	nodeAttributes := make(map[string]string, 0)
	nodeAttributes["cpu.arch"] = runtime.GOARCH
	return nodeAttributes, nil
}
