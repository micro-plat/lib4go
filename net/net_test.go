package net

import (
	"testing"
)

func TestIsTCPPortAvailable(t *testing.T) {
	if IsTCPPortAvailable(443) {
		t.Error("IsTCPPortAvailable is fail")
	}

	if !IsTCPPortAvailable(6666) {
		t.Error("IsTCPPortAvailable is fail")
	}
}

func TestRandomTCPPort(t *testing.T) {
	if RandomTCPPort() == -1 {
		t.Error("RandomTCPPort is fail")
	}
}

func TestGetAvailablePort(t *testing.T) {
	ports := []int{443, 6666}
	if GetAvailablePort(ports) == -1 || GetAvailablePort(ports) != 6666 {
		t.Error("GetAvailablePort is fail")
	}

	ports = []int{443, 22}
	if GetAvailablePort(ports) != -1 {
		t.Error("GetAvailablePort is fail")
	}
}
