// +build linux

package fingerprint

import (
	"fmt"

	client "github.com/hashicorp/nomad/client/config"
	"github.com/hashicorp/nomad/nomad/structs"
	"github.com/opencontainers/runc/libcontainer/cgroups"
)

// FindCgroupMountpointDir is used to find the cgroup mount point on a Linux
// system.
func FindCgroupMountpointDir() (string, error) {
	mount, err := cgroups.FindCgroupMountpointDir()
	if err != nil {
		switch e := err.(type) {
		case *cgroups.NotFoundError:
			// It's okay if the mount point is not discovered
			return "", nil
		default:
			// All other errors are passed back as is
			return "", e
		}
	}
	return mount, nil
}

// Fingerprint tries to find a valid cgroup moint point
func (f *CGroupFingerprint) Fingerprint(cfg *client.Config, node *structs.Node) (map[string]string, error) {
	nodeAttributes := make(map[string]string, 0)

	mount, err := f.mountPointDetector.MountPoint()
	if err != nil {
		nodeAttributes = f.clearCGroupAttributes(nodeAttributes)
		return nodeAttributes, fmt.Errorf("Failed to discover cgroup mount point: %s", err)
	}

	// Check if a cgroup mount point was found
	if mount == "" {
		// Clear any attributes from the previous fingerprint.
		nodeAttributes = f.clearCGroupAttributes(nodeAttributes)

		if f.lastState == cgroupAvailable {
			f.logger.Printf("[INFO] fingerprint.cgroups: cgroups are unavailable")
		}
		f.lastState = cgroupUnavailable
		return nodeAttributes, nil
	}

	nodeAttributes["unique.cgroup.mountpoint"] = mount

	if f.lastState == cgroupUnavailable {
		f.logger.Printf("[INFO] fingerprint.cgroups: cgroups are available")
	}
	f.lastState = cgroupAvailable
	return nodeAttributes, nil
}
