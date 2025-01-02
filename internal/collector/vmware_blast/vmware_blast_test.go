//go:build windows

package vmware_blast_test

import (
	"testing"

	"github.com/prometheus-community/windows_exporter/internal/collector/vmware_blast"
	"github.com/prometheus-community/windows_exporter/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, vmware_blast.Name, vmware_blast.NewWithFlags)
}

func TestCollector(t *testing.T) {
	testutils.TestCollector(t, vmware_blast.New, nil)
}
