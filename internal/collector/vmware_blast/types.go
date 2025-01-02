//go:build windows

package vmware_blast

type perfDataCounterValuesAudio struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`  // \VMware Blast Audio Counters(*)\Inbound Bandwidth (Kbps)
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"` // \VMware Blast Audio Counters(*)\Outbound Bandwidth (Kbps)
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`    // \VMware Blast Audio Counters(*)\Out Queueing time (us)
	ReceivedBytes         float64 `perfdata:"Received Bytes"`            // \VMware Blast Audio Counters(*)\Received Bytes
	ReceivedPackets       float64 `perfdata:"Received Packets"`          // \VMware Blast Audio Counters(*)\Received Packets
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`         // \VMware Blast Audio Counters(*)\Transmitted Bytes
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`       // \VMware Blast Audio Counters(*)\Transmitted Packets
}

type perfDataCounterValuesCDR struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`  // \VMware Blast CDR Counters(*)\Inbound Bandwidth (Kbps)
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"` // \VMware Blast CDR Counters(*)\Outbound Bandwidth (Kbps)
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`    // \VMware Blast CDR Counters(*)\Out Queueing time (us)
	ReceivedBytes         float64 `perfdata:"Received Bytes"`            // \VMware Blast CDR Counters(*)\Received Bytes
	ReceivedPackets       float64 `perfdata:"Received Packets"`          // \VMware Blast CDR Counters(*)\Received Packets
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`         // \VMware Blast CDR Counters(*)\Transmitted Bytes
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`       // \VMware Blast CDR Counters(*)\Transmitted Packets
}

type perfDataCounterValuesClipboard struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`  // \VMware Blast Clipboard Counters(*)\Inbound Bandwidth (Kbps)
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"` // \VMware Blast Clipboard Counters(*)\Outbound Bandwidth (Kbps)
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`    // \VMware Blast Clipboard Counters(*)\Out Queueing time (us)
	ReceivedBytes         float64 `perfdata:"Received Bytes"`            // \VMware Blast Clipboard Counters(*)\Received Bytes
	ReceivedPackets       float64 `perfdata:"Received Packets"`          // \VMware Blast Clipboard Counters(*)\Received Packets
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`         // \VMware Blast Clipboard Counters(*)\Transmitted Bytes
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`       // \VMware Blast Clipboard Counters(*)\Transmitted Packets
}

type perfDataCounterValuesHTML5MMR struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`  // \VMware Blast HTML5 MMR Counters(*)\Inbound Bandwidth (Kbps)
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"` // \VMware Blast HTML5 MMR Counters(*)\Outbound Bandwidth (Kbps)
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`    // \VMware Blast HTML5 MMR Counters(*)\Out Queueing time (us)
	ReceivedBytes         float64 `perfdata:"Received Bytes"`            // \VMware Blast HTML5 MMR Counters(*)\Received Bytes
	ReceivedPackets       float64 `perfdata:"Received Packets"`          // \VMware Blast HTML5 MMR Counters(*)\Received Packets
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`         // \VMware Blast HTML5 MMR Counters(*)\Transmitted Bytes
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`       // \VMware Blast HTML5 MMR Counters(*)\Transmitted Packets
}

type perfDataCounterValuesImaging struct {
	Dirtyframespersecond  float64 `perfdata:"Dirty frames per second"`   // \VMware Blast Imaging Counters(*)\Dirty frames per second
	EncoderType           float64 `perfdata:"Encoder Type"`              // \VMware Blast Imaging Counters(*)\Encoder Type
	FBCRate               float64 `perfdata:"FBC Rate"`                  // \VMware Blast Imaging Counters(*)\FBC Rate
	Framespersecond       float64 `perfdata:"Frame per second"`          // \VMware Blast Imaging Counters(*)\Frames per second
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`  // \VMware Blast Imaging Counters(*)\Inbound Bandwidth (Kbps)
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"` // \VMware Blast Imaging Counters(*)\Outbound Bandwidth (Kbps)
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`    // \VMware Blast Imaging Counters(*)\Out Queueing time (us)
	PollRate              float64 `perfdata:"Poll Rate"`                 // \VMware Blast Imaging Counters(*)\Poll Rate
	ReceivedBytes         float64 `perfdata:"Received Bytes"`            // \VMware Blast Imaging Counters(*)\Received Bytes
	ReceivedPackets       float64 `perfdata:"Received Packets"`          // \VMware Blast Imaging Counters(*)\Received Packets
	Totaldirtyframes      float64 `perfdata:"Total dirty frames"`        // \VMware Blast Imaging Counters(*)\Total dirty frames
	TotalFBC              float64 `perfdata:"Total FBC"`                 // \VMware Blast Imaging Counters(*)\Total FBC
	Totalframes           float64 `perfdata:"Total frames"`              // \VMware Blast Imaging Counters(*)\Total frames
	Totalpoll             float64 `perfdata:"Total poll"`                // \VMware Blast Imaging Counters(*)\Total poll
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`         // \VMware Blast Imaging Counters(*)\Transmitted Bytes
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`       // \VMware Blast Imaging Counters(*)\Transmitted Packets
}

type perfDataCounterValuesOtherFeature struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`  // \VMware Blast Other Feature Counters(*)\Inbound Bandwidth (Kbps)
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"` // \VMware Blast Other Feature Counters(*)\Outbound Bandwidth (Kbps)
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`    // \VMware Blast Other Feature Counters(*)\Out Queueing time (us)
	ReceivedBytes         float64 `perfdata:"Received Bytes"`            // \VMware Blast Other Feature Counters(*)\Received Bytes
	ReceivedPackets       float64 `perfdata:"Received Packets"`          // \VMware Blast Other Feature Counters(*)\Received Packets
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`         // \VMware Blast Other Feature Counters(*)\Transmitted Bytes
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`       // \VMware Blast Other Feature Counters(*)\Transmitted Packets
}

type perfDataCounterValuesPrinting struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`  // \VMware Blast Printing Counters(*)\Inbound Bandwidth (Kbps)
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"` // \VMware Blast Printing Counters(*)\Outbound Bandwidth (Kbps)
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`    // \VMware Blast Printing Counters(*)\Out Queueing time (us)
	ReceivedBytes         float64 `perfdata:"Received Bytes"`            // \VMware Blast Printing Counters(*)\Received Bytes
	ReceivedPackets       float64 `perfdata:"Received Packets"`          // \VMware Blast Printing Counters(*)\Received Packets
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`         // \VMware Blast Printing Counters(*)\Transmitted Bytes
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`       // \VMware Blast Printing Counters(*)\Transmitted Packets
}

type perfDataCounterValuesRdeServer struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`  // \VMware Blast RdeServer Counters(*)\Inbound Bandwidth (Kbps)
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"` // \VMware Blast RdeServer Counters(*)\Outbound Bandwidth (Kbps)
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`    // \VMware Blast RdeServer Counters(*)\Out Queueing time (us)
	ReceivedBytes         float64 `perfdata:"Received Bytes"`            // \VMware Blast RdeServer Counters(*)\Received Bytes
	ReceivedPackets       float64 `perfdata:"Received Packets"`          // \VMware Blast RdeServer Counters(*)\Received Packets
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`         // \VMware Blast RdeServer Counters(*)\Transmitted Bytes
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`       // \VMware Blast RdeServer Counters(*)\Transmitted Packets
}

type perfDataCounterValuesRTAV struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`  // \VMware Blast RTAV Counters(*)\Inbound Bandwidth (Kbps)
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"` // \VMware Blast RTAV Counters(*)\Outbound Bandwidth (Kbps)
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`    // \VMware Blast RTAV Counters(*)\Out Queueing time (us)
	ReceivedBytes         float64 `perfdata:"Received Bytes"`            // \VMware Blast RTAV Counters(*)\Received Bytes
	ReceivedPackets       float64 `perfdata:"Received Packets"`          // \VMware Blast RTAV Counters(*)\Received Packets
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`         // \VMware Blast RTAV Counters(*)\Transmitted Bytes
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`       // \VMware Blast RTAV Counters(*)\Transmitted Packets
}

type perfDataCounterValuesSDR struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`  // \VMware Blast SDR Counters(*)\Inbound Bandwidth (Kbps)
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"` // \VMware Blast SDR Counters(*)\Outbound Bandwidth (Kbps)
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`    // \VMware Blast SDR Counters(*)\Out Queueing time (us)
	ReceivedBytes         float64 `perfdata:"Received Bytes"`            // \VMware Blast SDR Counters(*)\Received Bytes
	ReceivedPackets       float64 `perfdata:"Received Packets"`          // \VMware Blast SDR Counters(*)\Received Packets
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`         // \VMware Blast SDR Counters(*)\Transmitted Bytes
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`       // \VMware Blast SDR Counters(*)\Transmitted Packets
}

type perfDataCounterValuesSerialPortandScanner struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`  // \VMware Blast Serial Port and Scanner Counters(*)\Inbound Bandwidth (Kbps)
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"` // \VMware Blast Serial Port and Scanner Counters(*)\Outbound Bandwidth (Kbps)
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`    // \VMware Blast Serial Port and Scanner Counters(*)\Out Queueing time (us)
	ReceivedBytes         float64 `perfdata:"Received Bytes"`            // \VMware Blast Serial Port and Scanner Counters(*)\Received Bytes
	ReceivedPackets       float64 `perfdata:"Received Packets"`          // \VMware Blast Serial Port and Scanner Counters(*)\Received Packets
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`         // \VMware Blast Serial Port and Scanner Counters(*)\Transmitted Bytes
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`       // \VMware Blast Serial Port and Scanner Counters(*)\Transmitted Packets
}

type perfDataCounterValuesSession struct {
	AutomaticReconnectCount              float64 `perfdata:"Automatic Reconnect Count"`                // \VMware Blast Session Counters(*)\Automatic Reconnect Count
	CumulativeReceivedBytesoverTCP       float64 `perfdata:"Cumulative Received Bytes over TCP"`       // \VMware Blast Session Counters(*)\Cumulative Received Bytes over TCP
	CumulativeReceivedBytesoverUDP       float64 `perfdata:"Cumulative Received Bytes over UDP"`       // \VMware Blast Session Counters(*)\Cumulative Received Bytes over UDP
	CumulativeTransmittedBytesoverTCP    float64 `perfdata:"Cumulative Transmitted Bytes over TCP"`    // \VMware Blast Session Counters(*)\Cumulative Transmitted Bytes over TCP
	CumulativeTransmittedBytesoverUDP    float64 `perfdata:"Cumulative Transmitted Bytes over UDP"`    // \VMware Blast Session Counters(*)\Cumulative Transmitted Bytes over UDP
	EstimatedBandwidthUplink             float64 `perfdata:"Estimated Bandwidth (Uplink)"`             // \VMware Blast Session Counters(*)\Estimated Bandwidth (Uplink)
	InstantaneousReceivedBytesoverTCP    float64 `perfdata:"Instantaneous Received Bytes over TCP"`    // \VMware Blast Session Counters(*)\Instantaneous Received Bytes over TCP
	InstantaneousReceivedBytesoverUDP    float64 `perfdata:"Instantaneous Received Bytes over UDP"`    // \VMware Blast Session Counters(*)\Instantaneous Received Bytes over UDP
	InstantaneousTransmittedBytesoverTCP float64 `perfdata:"Instantaneous Transmitted Bytes over TCP"` // \VMware Blast Session Counters(*)\Instantaneous Transmitted Bytes over TCP
	InstantaneousTransmittedBytesoverUDP float64 `perfdata:"Instantaneous Transmitted Bytes over UDP"` // \VMware Blast Session Counters(*)\Instantaneous Transmitted Bytes over UDP
	JitterUplink                         float64 `perfdata:"Jitter (Uplink)"`                          // \VMware Blast Session Counters(*)\Jitter (Uplink)
	PacketLossUplink                     float64 `perfdata:"Packet Loss (Uplink)"`                     // \VMware Blast Session Counters(*)\Packet Loss (Uplink)
	ReceivedBytes                        float64 `perfdata:"Received Bytes"`                           // \VMware Blast Session Counters(*)\Received Bytes
	ReceivedPackets                      float64 `perfdata:"Received Packets"`                         // \VMware Blast Session Counters(*)\Received Packets
	RTT                                  float64 `perfdata:"RTT"`                                      // \VMware Blast Session Counters(*)\RTT
	TransmittedBytes                     float64 `perfdata:"Transmitted Bytes"`                        // \VMware Blast Session Counters(*)\Transmitted Bytes
	TransmittedPackets                   float64 `perfdata:"Transmitted Packets"`                      // \VMware Blast Session Counters(*)\Transmitted Packets
}

type perfDataCounterValuesSmartCard struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`  // \VMware Blast Smart Card Counters(*)\Inbound Bandwidth (Kbps)
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"` // \VMware Blast Smart Card Counters(*)\Outbound Bandwidth (Kbps)
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`    // \VMware Blast Smart Card Counters(*)\Out Queueing time (us)
	ReceivedBytes         float64 `perfdata:"Received Bytes"`            // \VMware Blast Smart Card Counters(*)\Received Bytes
	ReceivedPackets       float64 `perfdata:"Received Packets"`          // \VMware Blast Smart Card Counters(*)\Received Packets
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`         // \VMware Blast Smart Card Counters(*)\Transmitted Bytes
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`       // \VMware Blast Smart Card Counters(*)\Transmitted Packets
}

type perfDataCounterValuesUSB struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`  // \VMware Blast USB Counters(*)\Inbound Bandwidth (Kbps)
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"` // \VMware Blast USB Counters(*)\Outbound Bandwidth (Kbps)
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`    // \VMware Blast USB Counters(*)\Out Queueing time (us)
	ReceivedBytes         float64 `perfdata:"Received Bytes"`            // \VMware Blast USB Counters(*)\Received Bytes
	ReceivedPackets       float64 `perfdata:"Received Packets"`          // \VMware Blast USB Counters(*)\Received Packets
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`         // \VMware Blast USB Counters(*)\Transmitted Bytes
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`       // \VMware Blast USB Counters(*)\Transmitted Packets
}

type perfDataCounterValuesViewScanner struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`  // \VMware Blast View Scanner Counters(*)\Inbound Bandwidth (Kbps)
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"` // \VMware Blast View Scanner Counters(*)\Outbound Bandwidth (Kbps)
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`    // \VMware Blast View Scanner Counters(*)\Out Queueing time (us)
	ReceivedBytes         float64 `perfdata:"Received Bytes"`            // \VMware Blast View Scanner Counters(*)\Received Bytes
	ReceivedPackets       float64 `perfdata:"Received Packets"`          // \VMware Blast View Scanner Counters(*)\Received Packets
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`         // \VMware Blast View Scanner Counters(*)\Transmitted Bytes
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`       // \VMware Blast View Scanner Counters(*)\Transmitted Packets
}

type perfDataCounterValuesWindowsMediaMMR struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`  // \VMware Blast Windows Media MMR Counters(*)\Inbound Bandwidth (Kbps)
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"` // \VMware Blast Windows Media MMR Counters(*)\Outbound Bandwidth (Kbps)
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`    // \VMware Blast Windows Media MMR Counters(*)\Out Queueing time (us)
	ReceivedBytes         float64 `perfdata:"Received Bytes"`            // \VMware Blast Windows Media MMR Counters(*)\Received Bytes
	ReceivedPackets       float64 `perfdata:"Received Packets"`          // \VMware Blast Windows Media MMR Counters(*)\Received Packets
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`         // \VMware Blast Windows Media MMR Counters(*)\Transmitted Bytes
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`       // \VMware Blast Windows Media MMR Counters(*)\Transmitted Packets
}
