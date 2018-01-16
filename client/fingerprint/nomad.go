package fingerprint

import (
	"log"

	client "github.com/hashicorp/nomad/client/config"
	"github.com/hashicorp/nomad/nomad/structs"
)

// NomadFingerprint is used to fingerprint the Nomad version
type NomadFingerprint struct {
	StaticFingerprinter
	logger *log.Logger
}

// NewNomadFingerprint is used to create a Nomad fingerprint
func NewNomadFingerprint(logger *log.Logger) Fingerprint {
	f := &NomadFingerprint{logger: logger}
	return f
}

func (f *NomadFingerprint) Fingerprint(config *client.Config, node *structs.Node) (map[string]string, error) {
	nodeAttributes := make(map[string]string, 0)
	nodeAttributes["nomad.version"] = config.Version.VersionNumber()
	nodeAttributes["nomad.revision"] = config.Version.Revision
	return nodeAttributes, nil
}
