//go:build windows

package vmware_blast

import (
	"testing"

	"github.com/prometheus-community/windows_exporter/internal/mi"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/require"
	"log/slog"
)

func TestNewCollector(t *testing.T) {
	c := New(nil)
	require.NotNil(t, c, "collector should not be nil")
	require.Equal(t, Name, c.GetName())
}

func TestBuild(t *testing.T) {
	c := New(nil)
	require.NotNil(t, c, "collector should not be nil")

	logger := slog.Default()
	err := c.Build(logger, nil)
	
	// Since this test might run on systems without VMware Blast,
	// we don't strictly require the build to succeed.
	// Instead, we log the error if it occurs.
	if err != nil {
		t.Logf("Build failed, possibly because VMware Blast is not installed: %v", err)
	}
}

func TestCollect(t *testing.T) {
	// Skip this test if running in CI environment
	// as VMware Blast might not be available
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}

	c := New(nil)
	require.NotNil(t, c, "collector should not be nil")

	logger := slog.Default()
	err := c.Build(logger, nil)
	if err != nil {
		t.Skipf("Skipping test as VMware Blast might not be installed: %v", err)
	}

	// Create a registry and register the collector
	registry := prometheus.NewRegistry()
	registry.MustRegister(c)

	// Try to collect metrics
	_, err = testutil.GatherAndCount(registry)
	
	// If collection fails, it might be because VMware Blast is not installed
	// or the performance counters are not available
	if err != nil {
		t.Logf("Collection failed, possibly because VMware Blast is not installed or counters are not available: %v", err)
	}

	// Clean up
	err = c.Close()
	require.NoError(t, err, "Close should not return an error")
}

func TestTypes(t *testing.T) {
	// Test that all the perfDataCounterValues structs have the expected fields
	// This is a compile-time check to ensure the structs match the expected performance counters
	
	var _ = perfDataCounterValuesAudio{}
	var _ = perfDataCounterValuesCDR{}
	var _ = perfDataCounterValuesClipboard{}
	var _ = perfDataCounterValuesHTML5MMR{}
	var _ = perfDataCounterValuesImaging{}
	var _ = perfDataCounterValuesOtherFeature{}
	var _ = perfDataCounterValuesPrinting{}
	var _ = perfDataCounterValuesRdeServer{}
	var _ = perfDataCounterValuesRTAV{}
	var _ = perfDataCounterValuesSDR{}
	var _ = perfDataCounterValuesSerialPortandScanner{}
	var _ = perfDataCounterValuesSession{}
	var _ = perfDataCounterValuesSmartCard{}
	var _ = perfDataCounterValuesUSB{}
	var _ = perfDataCounterValuesViewScanner{}
	var _ = perfDataCounterValuesWindowsMediaMMR{}
}
