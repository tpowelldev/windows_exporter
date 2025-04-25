//go:build windows

package vmware_blast

type perfDataCounterValuesAudio struct {
	InboundBandwidthKbps  float64 `perfdata:"Channel Receive Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Channel Send Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out queueing time in VVC (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesCDR struct {
	InboundBandwidthKbps  float64 `perfdata:"Channel Receive Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Channel Send Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out queueing time in VVC (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesClipboard struct {
	InboundBandwidthKbps  float64 `perfdata:"Channel Receive Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Channel Send Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out queueing time in VVC (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesHTML5MMR struct {
	InboundBandwidthKbps  float64 `perfdata:"Channel Receive Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Channel Send Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out queueing time in VVC (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesImaging struct {
	Dirtyframespersecond  float64 `perfdata:"Dirty frame rate per second"`
	EncoderType           float64 `perfdata:"Current encoder type"`
	FBCRate               float64 `perfdata:"FBC rate"`
	Framespersecond       float64 `perfdata:"Frame rate per second"`
	InboundBandwidthKbps  float64 `perfdata:"Channel Receive Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Channel Send Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out queueing time in VVC (us)"`
	PollRate              float64 `perfdata:"Poll rate"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	Totaldirtyframes      float64 `perfdata:"Total dirty frames for session"`
	TotalFBC              float64 `perfdata:"Total FBC for session"`
	Totalframes           float64 `perfdata:"Total frames for session"`
	Totalpoll             float64 `perfdata:"Total poll for session"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesOtherFeature struct {
	InboundBandwidthKbps  float64 `perfdata:"Channel Receive Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Channel Send Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out queueing time in VVC (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesPrinting struct {
	InboundBandwidthKbps  float64 `perfdata:"Channel Receive Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Channel Send Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out queueing time in VVC (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesRdeServer struct {
	InboundBandwidthKbps  float64 `perfdata:"Channel Receive Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Channel Send Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out queueing time in VVC (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesRTAV struct {
	InboundBandwidthKbps  float64 `perfdata:"Channel Receive Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Channel Send Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out queueing time in VVC (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesSDR struct {
	InboundBandwidthKbps  float64 `perfdata:"Channel Receive Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Channel Send Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out queueing time in VVC (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesSerialPortandScanner struct {
	InboundBandwidthKbps  float64 `perfdata:"Channel Receive Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Channel Send Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out queueing time in VVC (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesSession struct {
	AutomaticReconnectCount                 float64 `perfdata:"Total number of reconnects that happened after session interruptions"`
	CumulativeReceivedBytesoverTCP          float64 `perfdata:"Cumulative received bytes on the connection over TCP"`
	CumulativeReceivedBytesoverUDP          float64 `perfdata:"Cumulative received bytes on the connection over UDP"`
	CumulativeTransmittedBytesoverTCP       float64 `perfdata:"Cumulative transmitted bytes on the connection over TCP"`
	CumulativeTransmittedBytesoverUDP       float64 `perfdata:"Cumulative transmitted bytes on the connection over UDP"`
	EstimatedBandwidthUplink                float64 `perfdata:"Estimated network bandwidth for uplink in Kbps"`
	InstantaneousReceivedBytesoverTCP       float64 `perfdata:"Instantaneous received bytes on the connection over TCP"`
	InstantaneousReceivedBytesoverUDP       float64 `perfdata:"Instantaneous received bytes on the connection over UDP"`
	InstantaneousTransmittedBytesoverTCP    float64 `perfdata:"Instantaneous transmitted bytes on the connection over TCP"`
	InstantaneousTransmittedBytesoverUDP    float64 `perfdata:"Instantaneous transmitted bytes on the connection over UDP"`
	JitterUplink                            float64 `perfdata:"Network jitter for uplink in Milliseconds"`
	PacketLossUplink                        float64 `perfdata:"Network packet loss for uplink in percentage"`
	ReceivedBytes                           float64 `perfdata:"Received bytes on the connection"`
	ReceivedPackets                         float64 `perfdata:"Received packets"`
	RTT                                     float64 `perfdata:"Estimated RTT in Milliseconds"`
	TransmittedBytes                        float64 `perfdata:"Transmitted bytes on the connection"`
	TransmittedPackets                      float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesSmartCard struct {
	InboundBandwidthKbps  float64 `perfdata:"Channel Receive Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Channel Send Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out queueing time in VVC (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesUSB struct {
	InboundBandwidthKbps  float64 `perfdata:"Channel Receive Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Channel Send Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out queueing time in VVC (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesViewScanner struct {
	InboundBandwidthKbps  float64 `perfdata:"Channel Receive Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Channel Send Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out queueing time in VVC (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}

type perfDataCounterValuesWindowsMediaMMR struct {
	InboundBandwidthKbps  float64 `perfdata:"Channel Receive Bandwidth (Kbps)"`
	OutboundBandwidthKbps float64 `perfdata:"Channel Send Bandwidth (Kbps)"`
	OutQueueingtimeus     float64 `perfdata:"Out queueing time in VVC (us)"`
	ReceivedBytes         float64 `perfdata:"Received Bytes"`
	ReceivedPackets       float64 `perfdata:"Received Packets"`
	TransmittedBytes      float64 `perfdata:"Transmitted Bytes"`
	TransmittedPackets    float64 `perfdata:"Transmitted Packets"`
}
