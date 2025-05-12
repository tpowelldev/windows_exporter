//go:build windows

package vmware_blast

type perfDataCounterValuesAudio struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesCDR struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesClipboard struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesHTML5MMR struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesImaging struct {
	Dirtyframespersecond  float64 `perfdata:"Dirty frames per second"`
	EncoderType           float64 `perfdata:"Encoder Type"`
	FBCRate               float64 `perfdata:"FBC Rate"`
	Framespersecond       float64 `perfdata:"Frames per second"`
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`
	PollRate              float64 `perfdata:"Poll Rate"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	Totaldirtyframes      float64 `perfdata:"Total dirty frames"`
	TotalFBC              float64 `perfdata:"Total FBC"`
	Totalframes           float64 `perfdata:"Total frames"`
	Totalpoll             float64 `perfdata:"Total poll"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesOtherFeature struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesPrinting struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesRdeServer struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesRTAV struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesSDR struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesSerialPortandScanner struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesSession struct {
	AutomaticReconnectCount                 float64 `perfdata:"Automatic Reconnect Count"`
	CumulativeReceivedBytesoverTCP          float64 `perfdata:"Cumulative Received Bytes over TCP"`
	CumulativeReceivedBytesoverUDP          float64 `perfdata:"Cumulative Received Bytes over UDP"`
	CumulativeTransmittedBytesoverTCP       float64 `perfdata:"Cumulative Transmitted Bytes over TCP"`
	CumulativeTransmittedBytesoverUDP       float64 `perfdata:"Cumulative Transmitted Bytes over UDP"`
	EstimatedBandwidthUplink                float64 `perfdata:"Estimated Bandwidth (Uplink)"`
	InstantaneousReceivedBytesoverTCP       float64 `perfdata:"Instantaneous Received Bytes over TCP"`
	InstantaneousReceivedBytesoverUDP       float64 `perfdata:"Instantaneous Received Bytes over UDP"`
	InstantaneousTransmittedBytesoverTCP    float64 `perfdata:"Instantaneous Transmitted Bytes over TCP"`
	InstantaneousTransmittedBytesoverUDP    float64 `perfdata:"Instantaneous Transmitted Bytes over UDP"`
	JitterUplink                            float64 `perfdata:"Jitter (Uplink)"`
	PacketLossUplink                        float64 `perfdata:"Packet Loss (Uplink)"`
	ReceivedBytes                           float64 `perfdata:"Received Bytes"`
	ReceivedPackets                         float64 `perfdata:"Received Packets"`
	RTT                                     float64 `perfdata:"RTT"`
	TransmittedBytes                        float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets                      float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesSmartCard struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesUSB struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesViewScanner struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesWindowsMediaMMR struct {
	InboundBandwidthKbps  float64 `perfdata:"Inbound Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Outbound Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out Queueing time (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}
