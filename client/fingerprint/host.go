package fingerprint

import (
	"log"
	"runtime"

	"github.com/hashicorp/nomad/client/config"
	"github.com/hashicorp/nomad/nomad/structs"
	"github.com/shirou/gopsutil/host"
)

// HostFingerprint is used to fingerprint the host
type HostFingerprint struct {
	StaticFingerprinter
	logger *log.Logger
}

// NewHostFingerprint is used to create a Host fingerprint
func NewHostFingerprint(logger *log.Logger) Fingerprint {
	f := &HostFingerprint{logger: logger}
	return f
}

func (f *HostFingerprint) Fingerprint(cfg *config.Config, node *structs.Node) (map[string]string, error) {
	nodeAttributes := make(map[string]string, 0)

	hostInfo, err := host.Info()
	if err != nil {
		f.logger.Println("[WARN] Error retrieving host information: ", err)
		return nodeAttributes, err
	}

	nodeAttributes["os.name"] = hostInfo.Platform
	nodeAttributes["os.version"] = hostInfo.PlatformVersion

	nodeAttributes["kernel.name"] = runtime.GOOS
	nodeAttributes["kernel.version"] = hostInfo.KernelVersion

	nodeAttributes["unique.hostname"] = hostInfo.Hostname

	return nodeAttributes, nil
}
