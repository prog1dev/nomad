package fingerprint

import (
	"fmt"
	"log"

	"github.com/hashicorp/nomad/client/config"
	"github.com/hashicorp/nomad/nomad/structs"
	"github.com/shirou/gopsutil/mem"
)

// MemoryFingerprint is used to fingerprint the available memory on the node
type MemoryFingerprint struct {
	StaticFingerprinter
	logger *log.Logger
}

// NewMemoryFingerprint is used to create a Memory fingerprint
func NewMemoryFingerprint(logger *log.Logger) Fingerprint {
	f := &MemoryFingerprint{
		logger: logger,
	}
	return f
}

func (f *MemoryFingerprint) Fingerprint(cfg *config.Config, node *structs.Node) (map[string]string, error) {
	nodeAttributes := make(map[string]string, 0)

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		f.logger.Printf("[WARN] Error reading memory information: %s", err)
		return nodeAttributes, err
	}

	if memInfo.Total > 0 {
		nodeAttributes["memory.totalbytes"] = fmt.Sprintf("%d", memInfo.Total)

		if node.Resources == nil {
			node.Resources = &structs.Resources{}
		}
		node.Resources.MemoryMB = int(memInfo.Total / 1024 / 1024)
	}

	return nodeAttributes, nil
}
