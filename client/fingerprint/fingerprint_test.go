package fingerprint

// This file contains helper methods for testing fingerprinters

import (
	"log"
	"os"
	"testing"

	"github.com/hashicorp/nomad/client/config"
	"github.com/hashicorp/nomad/nomad/structs"
)

func testLogger() *log.Logger {
	return log.New(os.Stderr, "", log.LstdFlags)
}

func assertFingerprintOK(t *testing.T, fp Fingerprint, node *structs.Node) map[string]string {
	nodeAttributesDiff, err := fp.Fingerprint(new(config.Config), node)
	if err != nil {
		t.Fatalf("Failed to fingerprint: %s", err)
	}

	if len(nodeAttributesDiff) == 0 {
		t.Fatalf("Failed to apply node attributes")
	}

	return nodeAttributesDiff
}

func assertNodeAttributeContains(t *testing.T, nodeAttributes map[string]string, attribute string) {
	actual, found := nodeAttributes[attribute]
	if !found {
		t.Errorf("Expected to find Attribute `%s`\n\n[DEBUG] %#v", attribute, nodeAttributes)
		return
	}
	if actual == "" {
		t.Errorf("Expected non-empty Attribute value for `%s`\n\n[DEBUG] %#v", attribute, nodeAttributes)
	}
}

func assertNodeAttributeEquals(t *testing.T, nodeAttributes map[string]string, attribute string, expected string) {
	actual, found := nodeAttributes[attribute]
	if !found {
		t.Errorf("Expected to find Attribute `%s`; unable to check value\n\n[DEBUG] %#v", attribute, nodeAttributes)
		return
	}
	if expected != actual {
		t.Errorf("Expected `%s` Attribute to be `%s`, found `%s`\n\n[DEBUG] %#v", attribute, expected, actual, nodeAttributes)
	}
}

func assertNodeLinksContains(t *testing.T, node *structs.Node, link string) {
	actual, found := node.Links[link]
	if !found {
		t.Errorf("Expected to find Link `%s`\n\n[DEBUG] %#v", link, node)
		return
	}
	if actual == "" {
		t.Errorf("Expected non-empty Link value for `%s`\n\n[DEBUG] %#v", link, node)
	}
}
