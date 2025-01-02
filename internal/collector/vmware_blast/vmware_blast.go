//go:build windows

package vmware_blast

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/alecthomas/kingpin/v2"
	"github.com/prometheus-community/windows_exporter/internal/mi"
	"github.com/prometheus-community/windows_exporter/internal/pdh"
	"github.com/prometheus-community/windows_exporter/internal/types"
	"github.com/prometheus/client_golang/prometheus"
)

const Name = "vmware_blast"

type Config struct{}

//nolint:gochecknoglobals
var ConfigDefaults = Config{}

// A Collector is a Prometheus Collector for:
//Win32_PerfRawData_Counters_VMwareBlastAudioCounters
//Win32_PerfRawData_Counters_VMwareBlastCDRCounters
//Win32_PerfRawData_Counters_VMwareBlastClipboardCounters
//Win32_PerfRawData_Counters_VMwareBlastHTML5MMRCounters
//Win32_PerfRawData_Counters_VMwareBlastImagingCounters
//Win32_PerfRawData_Counters_VMwareBlastOtherFeatureCounters
//Win32_PerfRawData_Counters_VMwareBlastPrintingCounters
//Win32_PerfRawData_Counters_VMwareBlastRdeServerCounters
//Win32_PerfRawData_Counters_VMwareBlastRTAVCounters
//Win32_PerfRawData_Counters_VMwareBlastSDRCounters
//Win32_PerfRawData_Counters_VMwareBlastSerialPortandScannerCounters
//Win32_PerfRawData_Counters_VMwareBlastSessionCounters
//Win32_PerfRawData_Counters_VMwareBlastSmartCardCounters
//Win32_PerfRawData_Counters_VMwareBlastUSBCounters
//Win32_PerfRawData_Counters_VMwareBlastViewScannerCounters
//Win32_PerfRawData_Counters_VMwareBlastWindowsMediaMMRCounters

type Collector struct {
	config                                Config
	perfDataCollectorAudio                *pdh.Collector
	perfDataCollectorCDR                  *pdh.Collector
	perfDataCollectorClipboard            *pdh.Collector
	perfDataCollectorHTML5MMR             *pdh.Collector
	perfDataCollectorImaging              *pdh.Collector
	perfDataCollectorOtherFeature         *pdh.Collector
	perfDataCollectorPrinting             *pdh.Collector
	perfDataCollectorRdeServer            *pdh.Collector
	perfDataCollectorRTAV                 *pdh.Collector
	perfDataCollectorSDR                  *pdh.Collector
	perfDataCollectorSerialPortandScanner *pdh.Collector
	perfDataCollectorSession              *pdh.Collector
	perfDataCollectorSmartCard            *pdh.Collector
	perfDataCollectorUSB                  *pdh.Collector
	perfDataCollectorViewScanner          *pdh.Collector
	perfDataCollectorWindowsMediaMMR      *pdh.Collector
	perfDataObjectAudio                   []perfDataCounterValuesAudio
	perfDataObjectCDR                     []perfDataCounterValuesCDR
	perfDataObjectClipboard               []perfDataCounterValuesClipboard
	perfDataObjectHTML5MMR                []perfDataCounterValuesHTML5MMR
	perfDataObjectImaging                 []perfDataCounterValuesImaging
	perfDataObjectOtherFeature            []perfDataCounterValuesOtherFeature
	perfDataObjectPrinting                []perfDataCounterValuesPrinting
	perfDataObjectRdeServer               []perfDataCounterValuesRdeServer
	perfDataObjectRTAV                    []perfDataCounterValuesRTAV
	perfDataObjectSDR                     []perfDataCounterValuesSDR
	perfDataObjectSerialPortandScanner    []perfDataCounterValuesSerialPortandScanner
	perfDataObjectSession                 []perfDataCounterValuesSession
	perfDataObjectSmartCard               []perfDataCounterValuesSmartCard
	perfDataObjectUSB                     []perfDataCounterValuesUSB
	perfDataObjectViewScanner             []perfDataCounterValuesViewScanner
	perfDataObjectWindowsMediaMMR         []perfDataCounterValuesWindowsMediaMMR

	// Win32_PerfRawData_Counters_VMwareBlastAudioCounters
	audioInboundBandwidthKbps  *prometheus.Desc
	audioOutboundBandwidthKbps *prometheus.Desc
	audioOutQueueingTimeUs     *prometheus.Desc
	audioReceivedBytes         *prometheus.Desc
	audioReceivedPackets       *prometheus.Desc
	audioTransmittedBytes      *prometheus.Desc
	audioTransmittedPackets    *prometheus.Desc

	// Win32_PerfRawData_Counters_VMwareBlastCDRCounters
	cdrInboundBandwidthKbps  *prometheus.Desc
	cdrOutboundBandwidthKbps *prometheus.Desc
	cdrOutQueueingTimeUs     *prometheus.Desc
	cdrReceivedBytes         *prometheus.Desc
	cdrReceivedPackets       *prometheus.Desc
	cdrTransmittedBytes      *prometheus.Desc
	cdrTransmittedPackets    *prometheus.Desc

	// Win32_PerfRawData_Counters_VMwareBlastClipboardCounters
	clipboardInboundBandwidthKbps  *prometheus.Desc
	clipboardOutboundBandwidthKbps *prometheus.Desc
	clipboardOutQueueingTimeUs     *prometheus.Desc
	clipboardReceivedBytes         *prometheus.Desc
	clipboardReceivedPackets       *prometheus.Desc
	clipboardTransmittedBytes      *prometheus.Desc
	clipboardTransmittedPackets    *prometheus.Desc

	// Win32_PerfRawData_Counters_VMwareBlastHTML5MMRCounters
	html5MMRInboundBandwidthKbps  *prometheus.Desc
	html5MMROutboundBandwidthKbps *prometheus.Desc
	html5MMROutQueueingTimeUs     *prometheus.Desc
	html5MMRReceivedBytes         *prometheus.Desc
	html5MMRReceivedPackets       *prometheus.Desc
	html5MMRTransmittedBytes      *prometheus.Desc
	html5MMRTransmittedPackets    *prometheus.Desc

	// Win32_PerfRawData_Counters_VMwareBlastImagingCounters
	imagingDirtyFramesPerSecond  *prometheus.Desc
	imagingEncoderType           *prometheus.Desc
	imagingFBCRate               *prometheus.Desc
	imagingFramesPerSecond       *prometheus.Desc
	imagingInboundBandwidthKbps  *prometheus.Desc
	imagingOutboundBandwidthKbps *prometheus.Desc
	imagingOutQueueingTimeUs     *prometheus.Desc
	imagingPollRate              *prometheus.Desc
	imagingReceivedBytes         *prometheus.Desc
	imagingReceivedPackets       *prometheus.Desc
	imagingTotalDirtyFrames      *prometheus.Desc
	imagingTotalFBC              *prometheus.Desc
	imagingTotalFrames           *prometheus.Desc
	imagingTotalPoll             *prometheus.Desc
	imagingTransmittedBytes      *prometheus.Desc
	imagingTransmittedPackets    *prometheus.Desc

	// Win32_PerfRawData_Counters_VMwareBlastOtherFeatureCounters
	otherFeatureInboundBandwidthKbps  *prometheus.Desc
	otherFeatureOutboundBandwidthKbps *prometheus.Desc
	otherFeatureOutQueueingTimeUs     *prometheus.Desc
	otherFeatureReceivedBytes         *prometheus.Desc
	otherFeatureReceivedPackets       *prometheus.Desc
	otherFeatureTransmittedBytes      *prometheus.Desc
	otherFeatureTransmittedPackets    *prometheus.Desc

	// Win32_PerfRawData_Counters_VMwareBlastPrintingCounters
	printingInboundBandwidthKbps  *prometheus.Desc
	printingOutboundBandwidthKbps *prometheus.Desc
	printingOutQueueingTimeUs     *prometheus.Desc
	printingReceivedBytes         *prometheus.Desc
	printingReceivedPackets       *prometheus.Desc
	printingTransmittedBytes      *prometheus.Desc
	printingTransmittedPackets    *prometheus.Desc

	// Win32_PerfRawData_Counters_VMwareBlastRdeServerCounters
	rdeServerInboundBandwidthKbps  *prometheus.Desc
	rdeServerOutboundBandwidthKbps *prometheus.Desc
	rdeServerOutQueueingTimeUs     *prometheus.Desc
	rdeServerReceivedBytes         *prometheus.Desc
	rdeServerReceivedPackets       *prometheus.Desc
	rdeServerTransmittedBytes      *prometheus.Desc
	rdeServerTransmittedPackets    *prometheus.Desc

	// Win32_PerfRawData_Counters_VMwareBlastRTAVCounters
	rtavInboundBandwidthKbps  *prometheus.Desc
	rtavOutboundBandwidthKbps *prometheus.Desc
	rtavOutQueueingTimeUs     *prometheus.Desc
	rtavReceivedBytes         *prometheus.Desc
	rtavReceivedPackets       *prometheus.Desc
	rtavTransmittedBytes      *prometheus.Desc
	rtavTransmittedPackets    *prometheus.Desc

	// Win32_PerfRawData_Counters_VMwareBlastSDRCounters
	sdrInboundBandwidthKbps  *prometheus.Desc
	sdrOutboundBandwidthKbps *prometheus.Desc
	sdrOutQueueingTimeUs     *prometheus.Desc
	sdrReceivedBytes         *prometheus.Desc
	sdrReceivedPackets       *prometheus.Desc
	sdrTransmittedBytes      *prometheus.Desc
	sdrTransmittedPackets    *prometheus.Desc

	// Win32_PerfRawData_Counters_VMwareBlastSerialPortandScannerCounters
	serialPortandScannerInboundBandwidthKbps  *prometheus.Desc
	serialPortandScannerOutboundBandwidthKbps *prometheus.Desc
	serialPortandScannerOutQueueingTimeUs     *prometheus.Desc
	serialPortandScannerReceivedBytes         *prometheus.Desc
	serialPortandScannerReceivedPackets       *prometheus.Desc
	serialPortandScannerTransmittedBytes      *prometheus.Desc
	serialPortandScannerTransmittedPackets    *prometheus.Desc

	// Win32_PerfRawData_Counters_VMwareBlastSessionCounters
	sessionAutomaticReconnectCount              *prometheus.Desc
	sessionCumulativeReceivedBytesOverTCP       *prometheus.Desc
	sessionCumulativeReceivedBytesOverUDP       *prometheus.Desc
	sessionCumulativeTransmittedBytesOverTCP    *prometheus.Desc
	sessionCumulativeTransmittedBytesOverUDP    *prometheus.Desc
	sessionEstimatedBandwidthUplink             *prometheus.Desc
	sessionInstantaneousReceivedBytesOverTCP    *prometheus.Desc
	sessionInstantaneousReceivedBytesOverUDP    *prometheus.Desc
	sessionInstantaneousTransmittedBytesOverTCP *prometheus.Desc
	sessionInstantaneousTransmittedBytesOverUDP *prometheus.Desc
	sessionJitterUplink                         *prometheus.Desc
	sessionPacketLossUplink                     *prometheus.Desc
	sessionReceivedBytes                        *prometheus.Desc
	sessionReceivedPackets                      *prometheus.Desc
	sessionRTT                                  *prometheus.Desc
	sessionTransmittedBytes                     *prometheus.Desc
	sessionTransmittedPackets                   *prometheus.Desc

	// Win32_PerfRawData_Counters_VMwareBlastSmartCardCounters
	smartCardInboundBandwidthKbps  *prometheus.Desc
	smartCardOutboundBandwidthKbps *prometheus.Desc
	smartCardOutQueueingTimeUs     *prometheus.Desc
	smartCardReceivedBytes         *prometheus.Desc
	smartCardReceivedPackets       *prometheus.Desc
	smartCardTransmittedBytes      *prometheus.Desc
	smartCardTransmittedPackets    *prometheus.Desc

	// Win32_PerfRawData_Counters_VMwareBlastUSBCounters
	usbInboundBandwidthKbps  *prometheus.Desc
	usbOutboundBandwidthKbps *prometheus.Desc
	usbOutQueueingTimeUs     *prometheus.Desc
	usbReceivedBytes         *prometheus.Desc
	usbReceivedPackets       *prometheus.Desc
	usbTransmittedBytes      *prometheus.Desc
	usbTransmittedPackets    *prometheus.Desc

	// Win32_PerfRawData_Counters_VMwareBlastViewScannerCounters
	viewScannerInboundBandwidthKbps  *prometheus.Desc
	viewScannerOutboundBandwidthKbps *prometheus.Desc
	viewScannerOutQueueingTimeUs     *prometheus.Desc
	viewScannerReceivedBytes         *prometheus.Desc
	viewScannerReceivedPackets       *prometheus.Desc
	viewScannerTransmittedBytes      *prometheus.Desc
	viewScannerTransmittedPackets    *prometheus.Desc

	// Win32_PerfRawData_Counters_VMwareBlastWindowsMediaMMRCounters
	windowsMediaMMRInboundBandwidthKbps  *prometheus.Desc
	windowsMediaMMROutboundBandwidthKbps *prometheus.Desc
	windowsMediaMMROutQueueingTimeUs     *prometheus.Desc
	windowsMediaMMRReceivedBytes         *prometheus.Desc
	windowsMediaMMRReceivedPackets       *prometheus.Desc
	windowsMediaMMRTransmittedBytes      *prometheus.Desc
	windowsMediaMMRTransmittedPackets    *prometheus.Desc
}

func New(config *Config) *Collector {
	if config == nil {
		config = &ConfigDefaults
	}

	c := &Collector{
		config: *config,
	}

	return c
}

func NewWithFlags(_ *kingpin.Application) *Collector {
	return &Collector{}
}

func (c *Collector) GetName() string {
	return Name
}

func (c *Collector) Close() error {
	c.perfDataCollectorAudio.Close()
	c.perfDataCollectorCDR.Close()
	c.perfDataCollectorClipboard.Close()
	c.perfDataCollectorHTML5MMR.Close()
	c.perfDataCollectorImaging.Close()
	c.perfDataCollectorOtherFeature.Close()
	c.perfDataCollectorPrinting.Close()
	c.perfDataCollectorRdeServer.Close()
	c.perfDataCollectorRTAV.Close()
	c.perfDataCollectorSDR.Close()
	c.perfDataCollectorSerialPortandScanner.Close()
	c.perfDataCollectorSession.Close()
	c.perfDataCollectorSmartCard.Close()
	c.perfDataCollectorUSB.Close()
	c.perfDataCollectorViewScanner.Close()
	c.perfDataCollectorWindowsMediaMMR.Close()

	return nil
}

func (c *Collector) Build(_ *slog.Logger, _ *mi.Session) error {
	var err error

	c.perfDataCollectorAudio, err = pdh.NewCollector[perfDataCounterValuesAudio]("Blast Audio", pdh.InstancesTotal)
	if err != nil {
		return fmt.Errorf("failed to create Blast Audio collector: %w", err)
	}

	c.audioInboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "audio_inband_bandwidth_kbps"),
		"Channel Receive Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.audioOutboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "audio_outbound_bandwidth_kbps"),
		"Channel Send Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.audioOutQueueingTimeUs = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "audio_out_queuing_time_us"),
		"Out queueing time in VVC (us)",
		nil,
		nil,
	)
	c.audioReceivedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "audio_received_bytes_total"),
		"Received Bytes",
		nil,
		nil,
	)
	c.audioReceivedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "audio_received_packets_total"),
		"Received Packets",
		nil,
		nil,
	)
	c.audioTransmittedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "audio_transmitted_bytes_total"),
		"Transmitted Bytes",
		nil,
		nil,
	)
	c.audioTransmittedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "audio_transmitted_packets_total"),
		"Transmitted Packets",
		nil,
		nil,
	)

	c.perfDataCollectorCDR, err = pdh.NewCollector[perfDataCounterValuesCDR]("Blast CDR", pdh.InstancesTotal)
	if err != nil {
		return fmt.Errorf("failed to create Blast CDR collector: %w", err)
	}

	c.cdrInboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "cdr_inband_bandwidth_kbps"),
		"Channel Receive Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.cdrOutboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "cdr_outbound_bandwidth_kbps"),
		"Channel Send Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.cdrOutQueueingTimeUs = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "cdr_out_queuing_time_us"),
		"Out queueing time in VVC (us)",
		nil,
		nil,
	)
	c.cdrReceivedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "cdr_received_bytes_total"),
		"Received Bytes",
		nil,
		nil,
	)
	c.cdrReceivedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "cdr_received_packets_total"),
		"Received Packets",
		nil,
		nil,
	)
	c.cdrTransmittedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "cdr_transmitted_bytes_total"),
		"Transmitted Bytes",
		nil,
		nil,
	)
	c.cdrTransmittedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "cdr_transmitted_packets_total"),
		"Transmitted Packets",
		nil,
		nil,
	)

	c.perfDataCollectorClipboard, err = pdh.NewCollector[perfDataCounterValuesClipboard]("Blast Clipboard", pdh.InstancesTotal)
	if err != nil {
		return fmt.Errorf("failed to create Blast Clipboard collector: %w", err)
	}

	c.clipboardInboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "clipboard_inband_bandwidth_kbps"),
		"Channel Receive Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.clipboardOutboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "clipboard_outbound_bandwidth_kbps"),
		"Channel Send Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.clipboardOutQueueingTimeUs = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "clipboard_out_queuing_time_us"),
		"Out queueing time in VVC (us)",
		nil,
		nil,
	)
	c.clipboardReceivedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "clipboard_received_bytes_total"),
		"Received Bytes",
		nil,
		nil,
	)
	c.clipboardReceivedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "clipboard_received_packets_total"),
		"Received Packets",
		nil,
		nil,
	)
	c.clipboardTransmittedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "clipboard_transmitted_bytes_total"),
		"Transmitted Bytes",
		nil,
		nil,
	)
	c.clipboardTransmittedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "clipboard_transmitted_packets_total"),
		"Transmitted Packets",
		nil,
		nil,
	)

	c.perfDataCollectorHTML5MMR, err = pdh.NewCollector[perfDataCounterValuesHTML5MMR]("Blast HTML5 MMR ", pdh.InstancesTotal)
	if err != nil {
		return fmt.Errorf("failed to create Blast HTML5 MMR collector: %w", err)
	}

	c.html5MMRInboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "html5_mmr_inband_bandwidth_kbps"),
		"Channel Receive Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.html5MMROutboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "html5_mmr_outbound_bandwidth_kbps"),
		"Channel Send Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.html5MMROutQueueingTimeUs = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "html5_mmr_out_queuing_time_us"),
		"Out queueing time in VVC (us)",
		nil,
		nil,
	)
	c.html5MMRReceivedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "html5_mmr_received_bytes_total"),
		"Received Bytes",
		nil,
		nil,
	)
	c.html5MMRReceivedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "html5_mmr_received_packets_total"),
		"Received Packets",
		nil,
		nil,
	)
	c.html5MMRTransmittedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "html5_mmr_transmitted_bytes_total"),
		"Transmitted Bytes",
		nil,
		nil,
	)
	c.html5MMRTransmittedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "html5_mmr_transmitted_packets_total"),
		"Transmitted Packets",
		nil,
		nil,
	)

	c.perfDataCollectorImaging, err = pdh.NewCollector[perfDataCounterValuesImaging]("Blast HTML5 MMR", pdh.InstancesTotal)
	if err != nil {
		return fmt.Errorf("failed to create Blast Imaging collector: %w", err)
	}

	c.imagingDirtyFramesPerSecond = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "imaging_dirty_frames_per_second"),
		"Dirty frame rate per second",
		nil,
		nil,
	)
	c.imagingEncoderType = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "imaging_encoder_type"),
		"Current encoder type",
		nil,
		nil,
	)
	c.imagingFBCRate = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "imaging_fbc_rate"),
		"FBC rate",
		nil,
		nil,
	)
	c.imagingFramesPerSecond = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "imaging_frames_per_second"),
		"Frame rate per second",
		nil,
		nil,
	)
	c.imagingInboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "imaging_inband_bandwidth_kbps"),
		"Channel Receive Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.imagingOutboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "imaging_outbound_bandwidth_kbps"),
		"Channel Send Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.imagingOutQueueingTimeUs = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "imaging_out_queuing_time_us"),
		"Out queueing time in VVC (us)",
		nil,
		nil,
	)
	c.imagingPollRate = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "imaging_poll_rate"),
		"Poll rate",
		nil,
		nil,
	)
	c.imagingReceivedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "imaging_received_bytes_total"),
		"Received Bytes",
		nil,
		nil,
	)
	c.imagingReceivedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "imaging_received_packets_total"),
		"Received Packets",
		nil,
		nil,
	)
	c.imagingTotalDirtyFrames = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "imaging_total_dirty_frames_total"),
		"Total dirty frames for session",
		nil,
		nil,
	)
	c.imagingTotalFBC = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "imaging_fbc_total"),
		"Total FBC for session",
		nil,
		nil,
	)
	c.imagingTotalFrames = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "imaging_frames_total"),
		"Total frames for session",
		nil,
		nil,
	)
	c.imagingTotalPoll = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "imaging_poll_total"),
		"Total poll for session",
		nil,
		nil,
	)
	c.imagingTransmittedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "imaging_transmitted_bytes_total"),
		"Transmitted Bytes",
		nil,
		nil,
	)
	c.imagingTransmittedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "imaging_transmitted_packets_total"),
		"Transmitted Packets",
		nil,
		nil,
	)

	c.perfDataCollectorOtherFeature, err = pdh.NewCollector[perfDataCounterValuesOtherFeature]("Blast Other Feature", pdh.InstancesTotal)
	if err != nil {
		return fmt.Errorf("failed to create Blast Other Feature collector: %w", err)
	}

	c.otherFeatureInboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "other_feature_inband_bandwidth_kbps"),
		"Channel Receive Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.otherFeatureOutboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "other_feature_outbound_bandwidth_kbps"),
		"Channel Send Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.otherFeatureOutQueueingTimeUs = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "other_feature_out_queuing_time_us"),
		"Out queueing time in VVC (us)",
		nil,
		nil,
	)
	c.otherFeatureReceivedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "other_feature_received_bytes_total"),
		"Received Bytes",
		nil,
		nil,
	)
	c.otherFeatureReceivedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "other_feature_received_packets_total"),
		"Received Packets",
		nil,
		nil,
	)
	c.otherFeatureTransmittedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "other_feature_transmitted_bytes_total"),
		"Transmitted Bytes",
		nil,
		nil,
	)
	c.otherFeatureTransmittedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "other_feature_transmitted_packets_total"),
		"Transmitted Packets",
		nil,
		nil,
	)

	c.perfDataCollectorPrinting, err = pdh.NewCollector[perfDataCounterValuesPrinting]("Blast Printing", pdh.InstancesTotal)
	if err != nil {
		return fmt.Errorf("failed to create Blast Printing collector: %w", err)
	}

	c.printingInboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "printing_inband_bandwidth_kbps"),
		"Channel Receive Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.printingOutboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "printing_outbound_bandwidth_kbps"),
		"Channel Send Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.printingOutQueueingTimeUs = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "printing_out_queuing_time_us"),
		"Out queueing time in VVC (us)",
		nil,
		nil,
	)
	c.printingReceivedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "printing_received_bytes_total"),
		"Received Bytes",
		nil,
		nil,
	)
	c.printingReceivedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "printing_received_packets_total"),
		"Received Packets",
		nil,
		nil,
	)
	c.printingTransmittedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "printing_transmitted_bytes_total"),
		"Transmitted Bytes",
		nil,
		nil,
	)
	c.printingTransmittedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "printing_transmitted_packets_total"),
		"Transmitted Packets",
		nil,
		nil,
	)

	c.perfDataCollectorRdeServer, err = pdh.NewCollector[perfDataCounterValuesRdeServer]("Blast RdeServer", pdh.InstancesTotal)
	if err != nil {
		return fmt.Errorf("failed to create Blast RdeServer collector: %w", err)
	}

	c.rdeServerInboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "rdeserver_inband_bandwidth_kbps"),
		"Channel Receive Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.rdeServerOutboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "rdeserver_outbound_bandwidth_kbps"),
		"Channel Send Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.rdeServerOutQueueingTimeUs = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "rdeserver_out_queuing_time_us"),
		"Out queueing time in VVC (us)",
		nil,
		nil,
	)
	c.rdeServerReceivedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "rdeserver_received_bytes_total"),
		"Received Bytes",
		nil,
		nil,
	)
	c.rdeServerReceivedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "rdeserver_received_packets_total"),
		"Received Packets",
		nil,
		nil,
	)
	c.rdeServerTransmittedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "rdeserver_transmitted_bytes_total"),
		"Transmitted Bytes",
		nil,
		nil,
	)
	c.rdeServerTransmittedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "rdeserver_transmitted_packets_total"),
		"Transmitted Packets",
		nil,
		nil,
	)

	c.perfDataCollectorRTAV, err = pdh.NewCollector[perfDataCounterValuesRTAV]("Blast RTAV", pdh.InstancesTotal)
	if err != nil {
		return fmt.Errorf("failed to create Blast RTAV collector: %w", err)
	}

	c.rtavInboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "rtav_inband_bandwidth_kbps"),
		"Channel Receive Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.rtavOutboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "rtav_outbound_bandwidth_kbps"),
		"Channel Send Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.rtavOutQueueingTimeUs = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "rtav_out_queuing_time_us"),
		"Out queueing time in VVC (us)",
		nil,
		nil,
	)
	c.rtavReceivedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "rtav_received_bytes_total"),
		"Received Bytes",
		nil,
		nil,
	)
	c.rtavReceivedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "rtav_received_packets_total"),
		"Received Packets",
		nil,
		nil,
	)
	c.rtavTransmittedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "rtav_transmitted_bytes_total"),
		"Transmitted Bytes",
		nil,
		nil,
	)
	c.rtavTransmittedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "rtav_transmitted_packets_total"),
		"Transmitted Packets",
		nil,
		nil,
	)

	c.perfDataCollectorSDR, err = pdh.NewCollector[perfDataCounterValuesSDR]("Blast SDR", pdh.InstancesTotal)
	if err != nil {
		return fmt.Errorf("failed to create Blast SDR collector: %w", err)
	}

	c.sdrInboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "sdr_inband_bandwidth_kbps"),
		"Channel Receive Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.sdrOutboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "sdr_outbound_bandwidth_kbps"),
		"Channel Send Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.sdrOutQueueingTimeUs = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "sdr_out_queuing_time_us"),
		"Out queueing time in VVC (us)",
		nil,
		nil,
	)
	c.sdrReceivedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "sdr_received_bytes_total"),
		"Received Bytes",
		nil,
		nil,
	)
	c.sdrReceivedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "sdr_received_packets_total"),
		"Received Packets",
		nil,
		nil,
	)
	c.sdrTransmittedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "sdr_transmitted_bytes_total"),
		"Transmitted Bytes",
		nil,
		nil,
	)
	c.sdrTransmittedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "sdr_transmitted_packets_total"),
		"Transmitted Packets",
		nil,
		nil,
	)

	c.perfDataCollectorSerialPortandScanner, err = pdh.NewCollector[perfDataCounterValuesSerialPortandScanner]("Blast Serial Port and Scanner", pdh.InstancesTotal)
	if err != nil {
		return fmt.Errorf("failed to create Blast Serial Port and Scanner collector: %w", err)
	}

	c.serialPortandScannerReceivedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "serial_port_and_scanner_inband_bandwidth_kbps"),
		"Channel Receive Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.serialPortandScannerOutboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "serial_port_and_scanner_outbound_bandwidth_kbps"),
		"Channel Send Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.serialPortandScannerOutQueueingTimeUs = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "serial_port_and_scanner_out_queuing_time_us"),
		"Out queueing time in VVC (us)",
		nil,
		nil,
	)
	c.serialPortandScannerReceivedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "serial_port_and_scanner_received_bytes_total"),
		"Received Bytes",
		nil,
		nil,
	)
	c.serialPortandScannerReceivedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "serial_port_and_scanner_received_packets_total"),
		"Received Packets",
		nil,
		nil,
	)
	c.serialPortandScannerTransmittedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "serial_port_and_scanner_transmitted_bytes_total"),
		"Transmitted Bytes",
		nil,
		nil,
	)
	c.serialPortandScannerTransmittedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "serial_port_and_scanner_transmitted_packets_total"),
		"Transmitted Packets",
		nil,
		nil,
	)

	c.perfDataCollectorSession, err = pdh.NewCollector[perfDataCounterValuesSession]("Blast Session", pdh.InstancesTotal)
	if err != nil {
		return fmt.Errorf("failed to create Blast Session collector: %w", err)
	}

	c.sessionAutomaticReconnectCount = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "session_automatic_reconnect_count_total"),
		"Total number of reconnects that happened after session interruptions",
		nil,
		nil,
	)
	c.sessionCumulativeReceivedBytesOverTCP = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "session_cumlative_received_bytes_over_tcp_total"),
		"Cumulative received bytes on the connection over TCP",
		nil,
		nil,
	)
	c.sessionCumulativeReceivedBytesOverUDP = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "session_cumlative_received_bytes_over_udp_total"),
		"Cumulative received bytes on the connection over UDP",
		nil,
		nil,
	)
	c.sessionCumulativeTransmittedBytesOverTCP = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "session_cumlative_transmitted_bytes_over_tcp_total"),
		"Cumulative transmitted bytes on the connection over TCP",
		nil,
		nil,
	)
	c.sessionCumulativeTransmittedBytesOverUDP = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "session_cumlative_transmitted_bytes_over_udp_total"),
		"Cumulative transmitted bytes on the connection over UDP",
		nil,
		nil,
	)
	c.sessionEstimatedBandwidthUplink = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "session_estimated_bandwidth_uplink_kbps"),
		"Estimated network bandwidth for uplink in Kbps",
		nil,
		nil,
	)
	c.sessionInstantaneousReceivedBytesOverTCP = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "session_instantaneous_received_bytes_over_tcp_total"),
		"Instantaneous received bytes on the connection over TCP",
		nil,
		nil,
	)
	c.sessionInstantaneousReceivedBytesOverUDP = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "session_instantaneous_received_bytes_over_udp_total"),
		"Instantaneous received bytes on the connection over UDP",
		nil,
		nil,
	)
	c.sessionInstantaneousTransmittedBytesOverTCP = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "session_instantaneous_transmitted_bytes_over_tcp_total"),
		"Instantaneous transmitted bytes on the connection over TCP",
		nil,
		nil,
	)
	c.sessionInstantaneousTransmittedBytesOverUDP = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "session_instantaneous_transmitted_bytes_over_udp_total"),
		"Instantaneous transmitted bytes on the connection over UDP",
		nil,
		nil,
	)
	c.sessionJitterUplink = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "session_jitter_uplink_ms"),
		"Network jitter for uplink in Milliseconds",
		nil,
		nil,
	)
	c.sessionPacketLossUplink = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "session_packet_loss_uplink"),
		"Network packet loss for uplink in percentage",
		nil,
		nil,
	)
	c.sessionReceivedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "session_received_bytes_total"),
		"Received bytes on the connection",
		nil,
		nil,
	)
	c.sessionReceivedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "session_received_packets_total"),
		"Received packets",
		nil,
		nil,
	)
	c.sessionRTT = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "session_rtt_ms"),
		"Estimated RTT in Milliseconds",
		nil,
		nil,
	)
	c.sessionTransmittedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "session_transmitted_bytes_total"),
		"Transmitted bytes on the connection",
		nil,
		nil,
	)
	c.sessionTransmittedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "session_transmitted_packets_total"),
		"Transmitted Packets",
		nil,
		nil,
	)

	c.perfDataCollectorSmartCard, err = pdh.NewCollector[perfDataCounterValuesSmartCard]("Blast Smart Card", pdh.InstancesTotal)
	if err != nil {
		return fmt.Errorf("failed to create Blast Smart Card collector: %w", err)
	}

	c.smartCardInboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "smart_card_inband_bandwidth_kbps"),
		"Channel Receive Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.smartCardOutboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "smart_card_outbound_bandwidth_kbps"),
		"Channel Send Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.smartCardOutQueueingTimeUs = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "smart_card_out_queuing_time_us"),
		"Out queueing time in VVC (us)",
		nil,
		nil,
	)
	c.smartCardReceivedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "smart_card_received_bytes_total"),
		"Received Bytes",
		nil,
		nil,
	)
	c.smartCardReceivedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "smart_card_received_packets_total"),
		"Received Packets",
		nil,
		nil,
	)
	c.smartCardTransmittedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "smart_card_transmitted_bytes_total"),
		"Transmitted Bytes",
		nil,
		nil,
	)
	c.smartCardTransmittedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "smart_card_transmitted_packets_total"),
		"Transmitted Packets",
		nil,
		nil,
	)

	c.perfDataCollectorUSB, err = pdh.NewCollector[perfDataCounterValuesUSB]("Blast USB", pdh.InstancesTotal)
	if err != nil {
		return fmt.Errorf("failed to create Blast USB collector: %w", err)
	}

	c.usbInboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "usb_inband_bandwidth_kbps"),
		"Channel Receive Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.usbOutboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "usb_outbound_bandwidth_kbps"),
		"Channel Send Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.usbOutQueueingTimeUs = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "usb_out_queuing_time_us"),
		"Out queueing time in VVC (us)",
		nil,
		nil,
	)
	c.usbReceivedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "usb_received_bytes_total"),
		"Received Bytes",
		nil,
		nil,
	)
	c.usbReceivedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "usb_received_packets_total"),
		"Received Packets",
		nil,
		nil,
	)
	c.usbTransmittedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "usb_transmitted_bytes_total"),
		"Transmitted Bytes",
		nil,
		nil,
	)
	c.usbTransmittedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "usb_transmitted_packets_total"),
		"Transmitted Packets",
		nil,
		nil,
	)

	c.perfDataCollectorViewScanner, err = pdh.NewCollector[perfDataCounterValuesViewScanner]("Blast View Scanner", pdh.InstancesTotal)
	if err != nil {
		return fmt.Errorf("failed to create Blast View Scanner collector: %w", err)
	}

	c.viewScannerInboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "view_scanner_inband_bandwidth_kbps"),
		"Channel Receive Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.viewScannerOutboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "view_scanner_outbound_bandwidth_kbps"),
		"Channel Send Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.viewScannerOutQueueingTimeUs = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "view_scanner_out_queuing_time_us"),
		"Out queueing time in VVC (us)",
		nil,
		nil,
	)
	c.viewScannerReceivedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "view_scanner_received_bytes_total"),
		"Received Bytes",
		nil,
		nil,
	)
	c.viewScannerReceivedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "view_scanner_received_packets_total"),
		"Received Packets",
		nil,
		nil,
	)
	c.viewScannerTransmittedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "view_scanner_transmitted_bytes_total"),
		"Transmitted Bytes",
		nil,
		nil,
	)
	c.viewScannerTransmittedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "view_scanner_transmitted_packets_total"),
		"Transmitted Packets",
		nil,
		nil,
	)

	c.perfDataCollectorWindowsMediaMMR, err = pdh.NewCollector[perfDataCounterValuesWindowsMediaMMR]("Blast Windows Media MMR", pdh.InstancesTotal)
	if err != nil {
		return fmt.Errorf("failed to create Blast Windows Media MMR collector: %w", err)
	}

	c.windowsMediaMMRInboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "windows_media_mmr_inband_bandwidth_kbps"),
		"Channel Receive Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.windowsMediaMMROutboundBandwidthKbps = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "windows_media_mmr_outbound_bandwidth_kbps"),
		"Channel Send Bandwidth (Kbps)",
		nil,
		nil,
	)
	c.windowsMediaMMROutQueueingTimeUs = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "windows_media_mmr_out_queuing_time_us"),
		"Out queueing time in VVC (us)",
		nil,
		nil,
	)
	c.windowsMediaMMRReceivedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "windows_media_mmr_received_bytes_total"),
		"Received Bytes",
		nil,
		nil,
	)
	c.windowsMediaMMRReceivedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "windows_media_mmr_received_packets_total"),
		"Received Packets",
		nil,
		nil,
	)
	c.windowsMediaMMRTransmittedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "windows_media_mmr_transmitted_bytes_total"),
		"Transmitted Bytes",
		nil,
		nil,
	)
	c.windowsMediaMMRTransmittedPackets = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "windows_media_mmr_transmitted_packets_total"),
		"Transmitted Packets",
		nil,
		nil,
	)

	return nil
}

// Collect sends the metric values for each metric
// to the provided prometheus Metric channel.
func (c *Collector) Collect(ch chan<- prometheus.Metric) error {
	errs := make([]error, 0, 16)

	if err := c.collectAudio(ch); err != nil {
		errs = append(errs, fmt.Errorf("failed collecting blast audio metrics: %w", err))
	}

	if err := c.collectCdr(ch); err != nil {
		errs = append(errs, fmt.Errorf("failed collecting blast cdr metrics: %w", err))
	}

	if err := c.collectClipboard(ch); err != nil {
		errs = append(errs, fmt.Errorf("failed collecting blast clipboard metrics: %w", err))
	}

	if err := c.collectHtml5Mmr(ch); err != nil {
		errs = append(errs, fmt.Errorf("failed collecting blast html5 mmr metrics: %w", err))
	}

	if err := c.collectImaging(ch); err != nil {
		errs = append(errs, fmt.Errorf("failed collecting blast imaging metrics: %w", err))
	}

	if err := c.collectOtherFeature(ch); err != nil {
		errs = append(errs, fmt.Errorf("failed collecting blast other feature metrics: %w", err))
	}

	if err := c.collectPrinting(ch); err != nil {
		errs = append(errs, fmt.Errorf("failed collecting blast printing metrics: %w", err))
	}

	if err := c.collectRdeServer(ch); err != nil {
		errs = append(errs, fmt.Errorf("failed collecting blast rdeserver metrics: %w", err))
	}

	if err := c.collectRtav(ch); err != nil {
		errs = append(errs, fmt.Errorf("failed collecting blast rtav metrics: %w", err))
	}

	if err := c.collectSdr(ch); err != nil {
		errs = append(errs, fmt.Errorf("failed collecting blast sdr metrics: %w", err))
	}

	if err := c.collectSerialPortandScanner(ch); err != nil {
		errs = append(errs, fmt.Errorf("failed collecting blast serial port and scanner metrics: %w", err))
	}

	if err := c.collectSession(ch); err != nil {
		errs = append(errs, fmt.Errorf("failed collecting blast session metrics: %w", err))
	}

	if err := c.collectSmartCard(ch); err != nil {
		errs = append(errs, fmt.Errorf("failed collecting blast smart card metrics: %w", err))
	}

	if err := c.collectUsb(ch); err != nil {
		errs = append(errs, fmt.Errorf("failed collecting blast usb metrics: %w", err))
	}

	if err := c.collectViewScanner(ch); err != nil {
		errs = append(errs, fmt.Errorf("failed collecting blast view scanner metrics: %w", err))
	}

	if err := c.collectWindowsMediaMmr(ch); err != nil {
		errs = append(errs, fmt.Errorf("failed collecting blast windows media mmr metrics: %w", err))
	}

	return errors.Join(errs...)
}

func (c *Collector) collectAudio(ch chan<- prometheus.Metric) error {
	err := c.perfDataCollectorAudio.Collect(&c.perfDataObjectAudio)
	if err != nil {
		return fmt.Errorf("failed to collect Blast Audio metrics: %w", err)
	}

	ch <- prometheus.MustNewConstMetric(
		c.audioInboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectAudio[0].InboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.audioOutboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectAudio[0].OutboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.audioOutQueueingTimeUs,
		prometheus.GaugeValue,
		c.perfDataObjectAudio[0].OutQueueingtimeus,
	)

	ch <- prometheus.MustNewConstMetric(
		c.audioReceivedBytes,
		prometheus.CounterValue,
		c.perfDataObjectAudio[0].ReceivedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.audioReceivedPackets,
		prometheus.CounterValue,
		c.perfDataObjectAudio[0].ReceivedPackets,
	)

	ch <- prometheus.MustNewConstMetric(
		c.audioTransmittedBytes,
		prometheus.CounterValue,
		c.perfDataObjectAudio[0].TransmittedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.audioTransmittedPackets,
		prometheus.CounterValue,
		c.perfDataObjectAudio[0].TransmittedPackets,
	)

	return nil
}

func (c *Collector) collectCdr(ch chan<- prometheus.Metric) error {
	err := c.perfDataCollectorCDR.Collect(&c.perfDataObjectCDR)
	if err != nil {
		return fmt.Errorf("failed to collect Blast CDR metrics: %w", err)
	}

	ch <- prometheus.MustNewConstMetric(
		c.cdrInboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectCDR[0].InboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.cdrOutboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectCDR[0].OutboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.cdrOutQueueingTimeUs,
		prometheus.GaugeValue,
		c.perfDataObjectCDR[0].OutQueueingtimeus,
	)

	ch <- prometheus.MustNewConstMetric(
		c.cdrReceivedBytes,
		prometheus.CounterValue,
		c.perfDataObjectCDR[0].ReceivedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.cdrReceivedPackets,
		prometheus.CounterValue,
		c.perfDataObjectCDR[0].ReceivedPackets,
	)

	ch <- prometheus.MustNewConstMetric(
		c.cdrTransmittedBytes,
		prometheus.CounterValue,
		c.perfDataObjectCDR[0].TransmittedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.cdrTransmittedPackets,
		prometheus.CounterValue,
		c.perfDataObjectCDR[0].TransmittedPackets,
	)

	return nil
}

func (c *Collector) collectClipboard(ch chan<- prometheus.Metric) error {
	err := c.perfDataCollectorClipboard.Collect(&c.perfDataObjectClipboard)
	if err != nil {
		return fmt.Errorf("failed to collect Blast Clipboard metrics: %w", err)
	}

	ch <- prometheus.MustNewConstMetric(
		c.clipboardInboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectClipboard[0].InboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.clipboardOutboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectClipboard[0].OutboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.clipboardOutQueueingTimeUs,
		prometheus.GaugeValue,
		c.perfDataObjectClipboard[0].OutQueueingtimeus,
	)

	ch <- prometheus.MustNewConstMetric(
		c.clipboardReceivedBytes,
		prometheus.CounterValue,
		c.perfDataObjectClipboard[0].ReceivedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.clipboardReceivedPackets,
		prometheus.CounterValue,
		c.perfDataObjectClipboard[0].ReceivedPackets,
	)

	ch <- prometheus.MustNewConstMetric(
		c.clipboardTransmittedBytes,
		prometheus.CounterValue,
		c.perfDataObjectClipboard[0].TransmittedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.clipboardTransmittedPackets,
		prometheus.CounterValue,
		c.perfDataObjectClipboard[0].TransmittedPackets,
	)

	return nil
}

func (c *Collector) collectHtml5Mmr(ch chan<- prometheus.Metric) error {
	err := c.perfDataCollectorHTML5MMR.Collect(&c.perfDataObjectHTML5MMR)
	if err != nil {
		return fmt.Errorf("failed to collect Blast HTML5 MMR metrics: %w", err)
	}

	ch <- prometheus.MustNewConstMetric(
		c.html5MMRInboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectHTML5MMR[0].InboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.html5MMROutboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectHTML5MMR[0].OutboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.html5MMROutQueueingTimeUs,
		prometheus.GaugeValue,
		c.perfDataObjectHTML5MMR[0].OutQueueingtimeus,
	)

	ch <- prometheus.MustNewConstMetric(
		c.html5MMRReceivedBytes,
		prometheus.CounterValue,
		c.perfDataObjectHTML5MMR[0].ReceivedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.html5MMRReceivedPackets,
		prometheus.CounterValue,
		c.perfDataObjectHTML5MMR[0].ReceivedPackets,
	)

	ch <- prometheus.MustNewConstMetric(
		c.html5MMRTransmittedBytes,
		prometheus.CounterValue,
		c.perfDataObjectHTML5MMR[0].TransmittedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.html5MMRTransmittedPackets,
		prometheus.CounterValue,
		c.perfDataObjectHTML5MMR[0].TransmittedPackets,
	)

	return nil
}

func (c *Collector) collectImaging(ch chan<- prometheus.Metric) error {
	err := c.perfDataCollectorImaging.Collect(&c.perfDataObjectImaging)
	if err != nil {
		return fmt.Errorf("failed to collect Blast Imaging metrics: %w", err)
	}

	ch <- prometheus.MustNewConstMetric(
		c.imagingDirtyFramesPerSecond,
		prometheus.GaugeValue,
		c.perfDataObjectImaging[0].Dirtyframespersecond,
	)

	ch <- prometheus.MustNewConstMetric(
		c.imagingEncoderType,
		prometheus.GaugeValue,
		c.perfDataObjectImaging[0].EncoderType,
	)

	ch <- prometheus.MustNewConstMetric(
		c.imagingFBCRate,
		prometheus.GaugeValue,
		c.perfDataObjectImaging[0].FBCRate,
	)

	ch <- prometheus.MustNewConstMetric(
		c.imagingFramesPerSecond,
		prometheus.GaugeValue,
		c.perfDataObjectImaging[0].Framespersecond,
	)

	ch <- prometheus.MustNewConstMetric(
		c.imagingInboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectImaging[0].InboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.imagingOutboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectImaging[0].OutboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.imagingOutQueueingTimeUs,
		prometheus.GaugeValue,
		c.perfDataObjectImaging[0].OutQueueingtimeus,
	)

	ch <- prometheus.MustNewConstMetric(
		c.imagingPollRate,
		prometheus.GaugeValue,
		c.perfDataObjectImaging[0].PollRate,
	)

	ch <- prometheus.MustNewConstMetric(
		c.imagingReceivedBytes,
		prometheus.CounterValue,
		c.perfDataObjectImaging[0].ReceivedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.imagingReceivedPackets,
		prometheus.CounterValue,
		c.perfDataObjectImaging[0].ReceivedPackets,
	)

	ch <- prometheus.MustNewConstMetric(
		c.imagingTotalDirtyFrames,
		prometheus.CounterValue,
		c.perfDataObjectImaging[0].Totaldirtyframes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.imagingTotalFBC,
		prometheus.CounterValue,
		c.perfDataObjectImaging[0].TotalFBC,
	)

	ch <- prometheus.MustNewConstMetric(
		c.imagingTotalFrames,
		prometheus.CounterValue,
		c.perfDataObjectImaging[0].Totalframes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.imagingTotalPoll,
		prometheus.CounterValue,
		c.perfDataObjectImaging[0].Totalpoll,
	)

	ch <- prometheus.MustNewConstMetric(
		c.imagingTransmittedBytes,
		prometheus.CounterValue,
		c.perfDataObjectImaging[0].TransmittedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.imagingTransmittedPackets,
		prometheus.CounterValue,
		c.perfDataObjectImaging[0].TransmittedPackets,
	)

	return nil
}

func (c *Collector) collectOtherFeature(ch chan<- prometheus.Metric) error {
	err := c.perfDataCollectorOtherFeature.Collect(&c.perfDataObjectOtherFeature)
	if err != nil {
		return fmt.Errorf("failed to collect Blast Other Feature metrics: %w", err)
	}

	ch <- prometheus.MustNewConstMetric(
		c.otherFeatureInboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectOtherFeature[0].InboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.otherFeatureOutboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectOtherFeature[0].OutboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.otherFeatureOutQueueingTimeUs,
		prometheus.GaugeValue,
		c.perfDataObjectOtherFeature[0].OutQueueingtimeus,
	)

	ch <- prometheus.MustNewConstMetric(
		c.otherFeatureReceivedBytes,
		prometheus.CounterValue,
		c.perfDataObjectOtherFeature[0].ReceivedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.otherFeatureReceivedPackets,
		prometheus.CounterValue,
		c.perfDataObjectOtherFeature[0].ReceivedPackets,
	)

	ch <- prometheus.MustNewConstMetric(
		c.otherFeatureTransmittedBytes,
		prometheus.CounterValue,
		c.perfDataObjectOtherFeature[0].TransmittedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.otherFeatureTransmittedPackets,
		prometheus.CounterValue,
		c.perfDataObjectOtherFeature[0].TransmittedPackets,
	)

	return nil
}

func (c *Collector) collectPrinting(ch chan<- prometheus.Metric) error {
	err := c.perfDataCollectorPrinting.Collect(&c.perfDataObjectPrinting)
	if err != nil {
		return fmt.Errorf("failed to collect Blast Printing metrics: %w", err)
	}

	ch <- prometheus.MustNewConstMetric(
		c.printingInboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectPrinting[0].InboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.printingOutboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectPrinting[0].OutboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.printingOutQueueingTimeUs,
		prometheus.GaugeValue,
		c.perfDataObjectPrinting[0].OutQueueingtimeus,
	)

	ch <- prometheus.MustNewConstMetric(
		c.printingReceivedBytes,
		prometheus.CounterValue,
		c.perfDataObjectPrinting[0].ReceivedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.printingReceivedPackets,
		prometheus.CounterValue,
		c.perfDataObjectPrinting[0].ReceivedPackets,
	)

	ch <- prometheus.MustNewConstMetric(
		c.printingTransmittedBytes,
		prometheus.CounterValue,
		c.perfDataObjectPrinting[0].TransmittedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.printingTransmittedPackets,
		prometheus.CounterValue,
		c.perfDataObjectPrinting[0].TransmittedPackets,
	)

	return nil
}

func (c *Collector) collectRdeServer(ch chan<- prometheus.Metric) error {
	err := c.perfDataCollectorRdeServer.Collect(&c.perfDataObjectRdeServer)
	if err != nil {
		return fmt.Errorf("failed to collect Blast RdeServer metrics: %w", err)
	}

	ch <- prometheus.MustNewConstMetric(
		c.rdeServerInboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectRdeServer[0].InboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.rdeServerOutboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectRdeServer[0].OutboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.rdeServerOutQueueingTimeUs,
		prometheus.GaugeValue,
		c.perfDataObjectRdeServer[0].OutQueueingtimeus,
	)

	ch <- prometheus.MustNewConstMetric(
		c.rdeServerReceivedBytes,
		prometheus.CounterValue,
		c.perfDataObjectRdeServer[0].ReceivedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.rdeServerReceivedPackets,
		prometheus.CounterValue,
		c.perfDataObjectRdeServer[0].ReceivedPackets,
	)

	ch <- prometheus.MustNewConstMetric(
		c.rdeServerTransmittedBytes,
		prometheus.CounterValue,
		c.perfDataObjectRdeServer[0].TransmittedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.rdeServerTransmittedPackets,
		prometheus.CounterValue,
		c.perfDataObjectRdeServer[0].TransmittedPackets,
	)

	return nil
}

func (c *Collector) collectRtav(ch chan<- prometheus.Metric) error {
	err := c.perfDataCollectorRTAV.Collect(&c.perfDataObjectRTAV)
	if err != nil {
		return fmt.Errorf("failed to collect Blast RTAV metrics: %w", err)
	}

	ch <- prometheus.MustNewConstMetric(
		c.rtavInboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectRTAV[0].InboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.rtavOutboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectRTAV[0].OutboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.rtavOutQueueingTimeUs,
		prometheus.GaugeValue,
		c.perfDataObjectRTAV[0].OutQueueingtimeus,
	)

	ch <- prometheus.MustNewConstMetric(
		c.rtavReceivedBytes,
		prometheus.CounterValue,
		c.perfDataObjectRTAV[0].ReceivedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.rtavReceivedPackets,
		prometheus.CounterValue,
		c.perfDataObjectRTAV[0].ReceivedPackets,
	)

	ch <- prometheus.MustNewConstMetric(
		c.rtavTransmittedBytes,
		prometheus.CounterValue,
		c.perfDataObjectRTAV[0].TransmittedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.rtavTransmittedPackets,
		prometheus.CounterValue,
		c.perfDataObjectRTAV[0].TransmittedPackets,
	)

	return nil
}

func (c *Collector) collectSdr(ch chan<- prometheus.Metric) error {
	err := c.perfDataCollectorSDR.Collect(&c.perfDataObjectSDR)
	if err != nil {
		return fmt.Errorf("failed to collect Blast SDR metrics: %w", err)
	}

	ch <- prometheus.MustNewConstMetric(
		c.sdrInboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectSDR[0].InboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sdrOutboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectSDR[0].OutboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sdrOutQueueingTimeUs,
		prometheus.GaugeValue,
		c.perfDataObjectSDR[0].OutQueueingtimeus,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sdrReceivedBytes,
		prometheus.CounterValue,
		c.perfDataObjectSDR[0].ReceivedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sdrReceivedPackets,
		prometheus.CounterValue,
		c.perfDataObjectSDR[0].ReceivedPackets,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sdrTransmittedBytes,
		prometheus.CounterValue,
		c.perfDataObjectSDR[0].TransmittedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sdrTransmittedPackets,
		prometheus.CounterValue,
		c.perfDataObjectSDR[0].TransmittedPackets,
	)

	return nil
}

func (c *Collector) collectSerialPortandScanner(ch chan<- prometheus.Metric) error {
	err := c.perfDataCollectorSerialPortandScanner.Collect(&c.perfDataObjectSerialPortandScanner)
	if err != nil {
		return fmt.Errorf("failed to collect Blast Serial Port and Scanner metrics: %w", err)
	}

	ch <- prometheus.MustNewConstMetric(
		c.serialPortandScannerInboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectSerialPortandScanner[0].InboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.serialPortandScannerOutboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectSerialPortandScanner[0].OutboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.serialPortandScannerOutQueueingTimeUs,
		prometheus.GaugeValue,
		c.perfDataObjectSerialPortandScanner[0].OutQueueingtimeus,
	)

	ch <- prometheus.MustNewConstMetric(
		c.serialPortandScannerReceivedBytes,
		prometheus.CounterValue,
		c.perfDataObjectSerialPortandScanner[0].ReceivedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.serialPortandScannerReceivedPackets,
		prometheus.CounterValue,
		c.perfDataObjectSerialPortandScanner[0].ReceivedPackets,
	)

	ch <- prometheus.MustNewConstMetric(
		c.serialPortandScannerTransmittedBytes,
		prometheus.CounterValue,
		c.perfDataObjectSerialPortandScanner[0].TransmittedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.serialPortandScannerTransmittedPackets,
		prometheus.CounterValue,
		c.perfDataObjectSerialPortandScanner[0].TransmittedPackets,
	)

	return nil
}

func (c *Collector) collectSession(ch chan<- prometheus.Metric) error {
	err := c.perfDataCollectorSession.Collect(&c.perfDataObjectSession)
	if err != nil {
		return fmt.Errorf("failed to collect Blast Session metrics: %w", err)
	}

	ch <- prometheus.MustNewConstMetric(
		c.sessionAutomaticReconnectCount,
		prometheus.CounterValue,
		c.perfDataObjectSession[0].AutomaticReconnectCount,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sessionCumulativeReceivedBytesOverTCP,
		prometheus.CounterValue,
		c.perfDataObjectSession[0].CumulativeReceivedBytesoverTCP,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sessionCumulativeReceivedBytesOverUDP,
		prometheus.CounterValue,
		c.perfDataObjectSession[0].CumulativeReceivedBytesoverUDP,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sessionCumulativeTransmittedBytesOverTCP,
		prometheus.CounterValue,
		c.perfDataObjectSession[0].CumulativeTransmittedBytesoverTCP,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sessionCumulativeTransmittedBytesOverUDP,
		prometheus.CounterValue,
		c.perfDataObjectSession[0].CumulativeTransmittedBytesoverUDP,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sessionEstimatedBandwidthUplink,
		prometheus.GaugeValue,
		c.perfDataObjectSession[0].EstimatedBandwidthUplink,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sessionInstantaneousReceivedBytesOverTCP,
		prometheus.GaugeValue,
		c.perfDataObjectSession[0].InstantaneousReceivedBytesoverTCP,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sessionInstantaneousReceivedBytesOverUDP,
		prometheus.GaugeValue,
		c.perfDataObjectSession[0].InstantaneousReceivedBytesoverUDP,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sessionInstantaneousTransmittedBytesOverTCP,
		prometheus.GaugeValue,
		c.perfDataObjectSession[0].InstantaneousTransmittedBytesoverTCP,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sessionInstantaneousTransmittedBytesOverUDP,
		prometheus.GaugeValue,
		c.perfDataObjectSession[0].InstantaneousTransmittedBytesoverUDP,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sessionJitterUplink,
		prometheus.GaugeValue,
		c.perfDataObjectSession[0].JitterUplink,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sessionPacketLossUplink,
		prometheus.GaugeValue,
		c.perfDataObjectSession[0].PacketLossUplink,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sessionReceivedBytes,
		prometheus.CounterValue,
		c.perfDataObjectSession[0].ReceivedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sessionReceivedPackets,
		prometheus.CounterValue,
		c.perfDataObjectSession[0].ReceivedPackets,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sessionRTT,
		prometheus.GaugeValue,
		c.perfDataObjectSession[0].RTT,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sessionTransmittedBytes,
		prometheus.CounterValue,
		c.perfDataObjectSession[0].TransmittedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.sessionTransmittedPackets,
		prometheus.CounterValue,
		c.perfDataObjectSession[0].TransmittedPackets,
	)

	return nil
}

func (c *Collector) collectSmartCard(ch chan<- prometheus.Metric) error {
	err := c.perfDataCollectorSmartCard.Collect(&c.perfDataObjectSmartCard)
	if err != nil {
		return fmt.Errorf("failed to collect Blast Smart Card metrics: %w", err)
	}

	ch <- prometheus.MustNewConstMetric(
		c.smartCardInboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectSmartCard[0].InboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.smartCardOutboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectSmartCard[0].OutboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.smartCardOutQueueingTimeUs,
		prometheus.GaugeValue,
		c.perfDataObjectSmartCard[0].OutQueueingtimeus,
	)

	ch <- prometheus.MustNewConstMetric(
		c.smartCardReceivedBytes,
		prometheus.CounterValue,
		c.perfDataObjectSmartCard[0].ReceivedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.smartCardReceivedPackets,
		prometheus.CounterValue,
		c.perfDataObjectSmartCard[0].ReceivedPackets,
	)

	ch <- prometheus.MustNewConstMetric(
		c.smartCardTransmittedBytes,
		prometheus.CounterValue,
		c.perfDataObjectSmartCard[0].TransmittedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.smartCardTransmittedPackets,
		prometheus.CounterValue,
		c.perfDataObjectSmartCard[0].TransmittedPackets,
	)

	return nil
}

func (c *Collector) collectUsb(ch chan<- prometheus.Metric) error {
	err := c.perfDataCollectorUSB.Collect(&c.perfDataObjectUSB)
	if err != nil {
		return fmt.Errorf("failed to collect Blast USB metrics: %w", err)
	}

	ch <- prometheus.MustNewConstMetric(
		c.usbInboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectUSB[0].InboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.usbOutboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectUSB[0].OutboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.usbOutQueueingTimeUs,
		prometheus.GaugeValue,
		c.perfDataObjectUSB[0].OutQueueingtimeus,
	)

	ch <- prometheus.MustNewConstMetric(
		c.usbReceivedBytes,
		prometheus.CounterValue,
		c.perfDataObjectUSB[0].ReceivedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.usbReceivedPackets,
		prometheus.CounterValue,
		c.perfDataObjectUSB[0].ReceivedPackets,
	)

	ch <- prometheus.MustNewConstMetric(
		c.usbTransmittedBytes,
		prometheus.CounterValue,
		c.perfDataObjectUSB[0].TransmittedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.usbTransmittedPackets,
		prometheus.CounterValue,
		c.perfDataObjectUSB[0].TransmittedPackets,
	)

	return nil
}

func (c *Collector) collectViewScanner(ch chan<- prometheus.Metric) error {
	err := c.perfDataCollectorViewScanner.Collect(&c.perfDataObjectViewScanner)
	if err != nil {
		return fmt.Errorf("failed to collect Blast View Scanner metrics: %w", err)
	}

	ch <- prometheus.MustNewConstMetric(
		c.viewScannerInboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectViewScanner[0].InboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.viewScannerOutboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectViewScanner[0].OutboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.viewScannerOutQueueingTimeUs,
		prometheus.GaugeValue,
		c.perfDataObjectViewScanner[0].OutQueueingtimeus,
	)

	ch <- prometheus.MustNewConstMetric(
		c.viewScannerReceivedBytes,
		prometheus.CounterValue,
		c.perfDataObjectViewScanner[0].ReceivedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.viewScannerReceivedPackets,
		prometheus.CounterValue,
		c.perfDataObjectViewScanner[0].ReceivedPackets,
	)

	ch <- prometheus.MustNewConstMetric(
		c.viewScannerTransmittedBytes,
		prometheus.CounterValue,
		c.perfDataObjectViewScanner[0].TransmittedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.viewScannerTransmittedPackets,
		prometheus.CounterValue,
		c.perfDataObjectViewScanner[0].TransmittedPackets,
	)

	return nil
}

func (c *Collector) collectWindowsMediaMmr(ch chan<- prometheus.Metric) error {
	err := c.perfDataCollectorWindowsMediaMMR.Collect(&c.perfDataObjectWindowsMediaMMR)
	if err != nil {
		return fmt.Errorf("failed to collect Blast Windows Media MMR metrics: %w", err)
	}

	ch <- prometheus.MustNewConstMetric(
		c.windowsMediaMMRInboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectWindowsMediaMMR[0].InboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.windowsMediaMMROutboundBandwidthKbps,
		prometheus.GaugeValue,
		c.perfDataObjectWindowsMediaMMR[0].OutboundBandwidthKbps,
	)

	ch <- prometheus.MustNewConstMetric(
		c.windowsMediaMMROutQueueingTimeUs,
		prometheus.GaugeValue,
		c.perfDataObjectWindowsMediaMMR[0].OutQueueingtimeus,
	)

	ch <- prometheus.MustNewConstMetric(
		c.windowsMediaMMRReceivedBytes,
		prometheus.CounterValue,
		c.perfDataObjectWindowsMediaMMR[0].ReceivedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.windowsMediaMMRReceivedPackets,
		prometheus.CounterValue,
		c.perfDataObjectWindowsMediaMMR[0].ReceivedPackets,
	)

	ch <- prometheus.MustNewConstMetric(
		c.windowsMediaMMRTransmittedBytes,
		prometheus.CounterValue,
		c.perfDataObjectWindowsMediaMMR[0].TransmittedBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		c.windowsMediaMMRTransmittedPackets,
		prometheus.CounterValue,
		c.perfDataObjectWindowsMediaMMR[0].TransmittedPackets,
	)

	return nil
}
