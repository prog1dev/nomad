package fingerprint

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/hashicorp/nomad/client/config"
	"github.com/hashicorp/nomad/nomad/structs"
)

const bytesPerMegabyte = 1024 * 1024

// StorageFingerprint is used to measure the amount of storage free for
// applications that the Nomad agent will run on this machine.
type StorageFingerprint struct {
	StaticFingerprinter
	logger *log.Logger
}

func NewStorageFingerprint(logger *log.Logger) Fingerprint {
	fp := &StorageFingerprint{logger: logger}
	return fp
}

func (f *StorageFingerprint) Fingerprint(cfg *config.Config, node *structs.Node) (map[string]string, error) {
	nodeAttributes := make(map[string]string, 0)

	// Initialize these to empty defaults
	nodeAttributes["unique.storage.volume"] = ""
	nodeAttributes["unique.storage.bytestotal"] = ""
	nodeAttributes["unique.storage.bytesfree"] = ""
	if node.Resources == nil {
		node.Resources = &structs.Resources{}
	}

	// Guard against unset AllocDir
	storageDir := cfg.AllocDir
	if storageDir == "" {
		var err error
		storageDir, err = os.Getwd()
		if err != nil {
			return nodeAttributes, fmt.Errorf("unable to get CWD from filesystem: %s", err)
		}
	}

	volume, total, free, err := f.diskFree(storageDir)
	if err != nil {
		return nodeAttributes, fmt.Errorf("failed to determine disk space for %s: %v", storageDir, err)
	}

	nodeAttributes["unique.storage.volume"] = volume
	nodeAttributes["unique.storage.bytestotal"] = strconv.FormatUint(total, 10)
	nodeAttributes["unique.storage.bytesfree"] = strconv.FormatUint(free, 10)

	node.Resources.DiskMB = int(free / bytesPerMegabyte)

	return nodeAttributes, nil
}
