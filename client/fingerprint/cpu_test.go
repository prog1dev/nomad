package fingerprint

import (
	"testing"

	"github.com/hashicorp/nomad/client/config"
	"github.com/hashicorp/nomad/nomad/structs"
)

func TestCPUFingerprint(t *testing.T) {
	f := NewCPUFingerprint(testLogger())
	node := &structs.Node{
		Attributes: make(map[string]string),
	}
	nodeAttributes, err := f.Fingerprint(&config.Config{}, node)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	// CPU info
	if nodeAttributes["cpu.numcores"] == "" {
		t.Fatalf("Missing Num Cores")
	}
	if nodeAttributes["cpu.modelname"] == "" {
		t.Fatalf("Missing Model Name")
	}

	if nodeAttributes["cpu.frequency"] == "" {
		t.Fatalf("Missing CPU Frequency")
	}
	if nodeAttributes["cpu.totalcompute"] == "" {
		t.Fatalf("Missing CPU Total Compute")
	}

	if node.Resources == nil || node.Resources.CPU == 0 {
		t.Fatalf("Expected to find CPU Resources")
	}

}

// TestCPUFingerprint_OverrideCompute asserts that setting cpu_total_compute in
// the client config overrides the detected CPU freq (if any).
func TestCPUFingerprint_OverrideCompute(t *testing.T) {
	f := NewCPUFingerprint(testLogger())
	node := &structs.Node{
		Attributes: make(map[string]string),
	}
	cfg := &config.Config{}

	{
		_, err := f.Fingerprint(cfg, node)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		if node.Resources.CPU == 0 {
			t.Fatalf("expected fingerprint of cpu of but found 0")
		}
	}

	// Get actual system CPU
	origCPU := node.Resources.CPU

	// Override it with a setting
	cfg.CpuCompute = origCPU + 123

	{
		// Make sure the Fingerprinter applies the override to the node resources
		_, err := f.Fingerprint(cfg, node)
		if err != nil {
			t.Fatalf("err: %v", err)
		}

		if node.Resources.CPU != cfg.CpuCompute {
			t.Fatalf("expected override cpu of %d but found %d", cfg.CpuCompute, node.Resources.CPU)
		}
	}
}
