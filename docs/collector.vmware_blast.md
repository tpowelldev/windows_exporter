# vmware_blast collector

The vmware_blast collector exposes metrics relating to VMware Blast sessions

|                     |                      |
|---------------------|----------------------|
| Metric name prefix  | `vmware_blast`       |
| Source              | Performance counters |
| Enabled by default? | No                   |

## Flags

None

## Metrics

Some of these metrics may not be collected, depending on the installation options chosen when installing the VMware Horizon agent

| Name                                                                          | Description                                                          | Type    | Labels |
|-------------------------------------------------------------------------------|----------------------------------------------------------------------|---------|--------|
| `windows_vmware_blast_audio_inband_bandwidth_kbps`                            | Channel Receive Bandwidth (Kbps)                                     | gauge   | None   |
| `windows_vmware_blast_audio_out_queuing_time_us`                              | Out queueing time in VVC (us)                                        | gauge   | None   |
| `windows_vmware_blast_audio_outbound_bandwidth_kbps`                          | Channel Send Bandwidth (Kbps)                                        | gauge   | None   |
| `windows_vmware_blast_audio_received_bytes_total`                             | Received Bytes                                                       | counter | None   |
| `windows_vmware_blast_audio_received_packets_total`                           | Received Packets                                                     | counter | None   |
| `windows_vmware_blast_audio_transmitted_bytes_total`                          | Transmitted Bytes                                                    | counter | None   |
| `windows_vmware_blast_audio_transmitted_packets_total`                        | Transmitted Packets                                                  | counter | None   |
| `windows_vmware_blast_cdr_inband_bandwidth_kbps`                              | Channel Receive Bandwidth (Kbps)                                     | gauge   | None   |
| `windows_vmware_blast_cdr_out_queuing_time_us`                                | Out queueing time in VVC (us)                                        | gauge   | None   |
| `windows_vmware_blast_cdr_outbound_bandwidth_kbps`                            | Channel Send Bandwidth (Kbps)                                        | gauge   | None   |
| `windows_vmware_blast_cdr_received_bytes_total`                               | Received Bytes                                                       | counter | None   |
| `windows_vmware_blast_cdr_received_packets_total`                             | Received Packets                                                     | counter | None   |
| `windows_vmware_blast_cdr_transmitted_bytes_total`                            | Transmitted Bytes                                                    | counter | None   |
| `windows_vmware_blast_cdr_transmitted_packets_total`                          | Transmitted Packets                                                  | counter | None   |
| `windows_vmware_blast_clipboard_inband_bandwidth_kbps`                        | Channel Receive Bandwidth (Kbps)                                     | gauge   | None   |
| `windows_vmware_blast_clipboard_out_queuing_time_us`                          | Out queueing time in VVC (us)                                        | gauge   | None   |
| `windows_vmware_blast_clipboard_outbound_bandwidth_kbps`                      | Channel Send Bandwidth (Kbps)                                        | gauge   | None   |
| `windows_vmware_blast_clipboard_received_bytes_total`                         | Received Bytes                                                       | counter | None   |
| `windows_vmware_blast_clipboard_received_packets_total`                       | Received Packets                                                     | counter | None   |
| `windows_vmware_blast_clipboard_transmitted_bytes_total`                      | Transmitted Bytes                                                    | counter | None   |
| `windows_vmware_blast_clipboard_transmitted_packets_total`                    | Transmitted Packets                                                  | counter | None   |
| `windows_vmware_blast_html5_mmr_inband_bandwidth_kbps`                        | Channel Receive Bandwidth (Kbps)                                     | gauge   | None   |
| `windows_vmware_blast_html5_mmr_out_queuing_time_us`                          | Out queueing time in VVC (us)                                        | gauge   | None   |
| `windows_vmware_blast_html5_mmr_outbound_bandwidth_kbps`                      | Channel Send Bandwidth (Kbps)                                        | gauge   | None   |
| `windows_vmware_blast_html5_mmr_received_bytes_total`                         | Received Bytes                                                       | counter | None   |
| `windows_vmware_blast_html5_mmr_received_packets_total`                       | Received Packets                                                     | counter | None   |
| `windows_vmware_blast_html5_mmr_transmitted_bytes_total`                      | Transmitted Bytes                                                    | counter | None   |
| `windows_vmware_blast_html5_mmr_transmitted_packets_total`                    | Transmitted Packets                                                  | counter | None   |
| `windows_vmware_blast_imaging_dirty_frames_per_second`                        | Dirty frame rate per second                                          | gauge   | None   |
| `windows_vmware_blast_imaging_encoder_type`                                   | Current encoder type                                                 | gauge   | None   |
| `windows_vmware_blast_imaging_fbc_rate`                                       | FBC rate                                                             | gauge   | None   |
| `windows_vmware_blast_imaging_fbc_total`                                      | Total FBC for session                                                | counter | None   |
| `windows_vmware_blast_imaging_frames_per_second`                              | Frame rate per second                                                | gauge   | None   |
| `windows_vmware_blast_imaging_frames_total`                                   | Total frames for session                                             | counter | None   |
| `windows_vmware_blast_imaging_inband_bandwidth_kbps`                          | Channel Receive Bandwidth (Kbps)                                     | gauge   | None   |
| `windows_vmware_blast_imaging_out_queuing_time_us`                            | Out queueing time in VVC (us)                                        | gauge   | None   |
| `windows_vmware_blast_imaging_outbound_bandwidth_kbps`                        | Channel Send Bandwidth (Kbps)                                        | gauge   | None   |
| `windows_vmware_blast_imaging_poll_rate`                                      | Poll rate                                                            | gauge   | None   |
| `windows_vmware_blast_imaging_poll_total`                                     | Total poll for session                                               | counter | None   |
| `windows_vmware_blast_imaging_received_bytes_total`                           | Received Bytes                                                       | counter | None   |
| `windows_vmware_blast_imaging_received_packets_total`                         | Received Packets                                                     | counter | None   |
| `windows_vmware_blast_imaging_total_dirty_frames_total`                       | Total dirty frames for session                                       | counter | None   |
| `windows_vmware_blast_imaging_transmitted_bytes_total`                        | Transmitted Bytes                                                    | counter | None   |
| `windows_vmware_blast_imaging_transmitted_packets_total`                      | Transmitted Packets                                                  | counter | None   |
| `windows_vmware_blast_other_feature_inband_bandwidth_kbps`                    | Channel Receive Bandwidth (Kbps)                                     | gauge   | None   |
| `windows_vmware_blast_other_feature_out_queuing_time_us`                      | Out queueing time in VVC (us)                                        | gauge   | None   |
| `windows_vmware_blast_other_feature_outbound_bandwidth_kbps`                  | Channel Send Bandwidth (Kbps)                                        | gauge   | None   |
| `windows_vmware_blast_other_feature_received_bytes_total`                     | Received Bytes                                                       | counter | None   |
| `windows_vmware_blast_other_feature_received_packets_total`                   | Received Packets                                                     | counter | None   |
| `windows_vmware_blast_other_feature_transmitted_bytes_total`                  | Transmitted Bytes                                                    | counter | None   |
| `windows_vmware_blast_other_feature_transmitted_packets_total`                | Transmitted Packets                                                  | counter | None   |
| `windows_vmware_blast_printing_inband_bandwidth_kbps`                         | Channel Receive Bandwidth (Kbps)                                     | gauge   | None   |
| `windows_vmware_blast_printing_out_queuing_time_us`                           | Out queueing time in VVC (us)                                        | gauge   | None   |
| `windows_vmware_blast_printing_outbound_bandwidth_kbps`                       | Channel Send Bandwidth (Kbps)                                        | gauge   | None   |
| `windows_vmware_blast_printing_received_bytes_total`                          | Received Bytes                                                       | counter | None   |
| `windows_vmware_blast_printing_received_packets_total`                        | Received Packets                                                     | counter | None   |
| `windows_vmware_blast_printing_transmitted_bytes_total`                       | Transmitted Bytes                                                    | counter | None   |
| `windows_vmware_blast_printing_transmitted_packets_total`                     | Transmitted Packets                                                  | counter | None   |
| `windows_vmware_blast_rdeserver_inband_bandwidth_kbps`                        | Channel Receive Bandwidth (Kbps)                                     | gauge   | None   |
| `windows_vmware_blast_rdeserver_out_queuing_time_us`                          | Out queueing time in VVC (us)                                        | gauge   | None   |
| `windows_vmware_blast_rdeserver_outbound_bandwidth_kbps`                      | Channel Send Bandwidth (Kbps)                                        | gauge   | None   |
| `windows_vmware_blast_rdeserver_received_bytes_total`                         | Received Bytes                                                       | counter | None   |
| `windows_vmware_blast_rdeserver_received_packets_total`                       | Received Packets                                                     | counter | None   |
| `windows_vmware_blast_rdeserver_transmitted_bytes_total`                      | Transmitted Bytes                                                    | counter | None   |
| `windows_vmware_blast_rdeserver_transmitted_packets_total`                    | Transmitted Packets                                                  | counter | None   |
| `windows_vmware_blast_rtav_inband_bandwidth_kbps`                             | Channel Receive Bandwidth (Kbps)                                     | gauge   | None   |
| `windows_vmware_blast_rtav_out_queuing_time_us`                               | Out queueing time in VVC (us)                                        | gauge   | None   |
| `windows_vmware_blast_rtav_outbound_bandwidth_kbps`                           | Channel Send Bandwidth (Kbps)                                        | gauge   | None   |
| `windows_vmware_blast_rtav_received_bytes_total`                              | Received Bytes                                                       | counter | None   |
| `windows_vmware_blast_rtav_received_packets_total`                            | Received Packets                                                     | counter | None   |
| `windows_vmware_blast_rtav_transmitted_bytes_total`                           | Transmitted Bytes                                                    | counter | None   |
| `windows_vmware_blast_rtav_transmitted_packets_total`                         | Transmitted Packets                                                  | counter | None   |
| `windows_vmware_blast_sdr_inband_bandwidth_kbps`                              | Channel Receive Bandwidth (Kbps)                                     | gauge   | None   |
| `windows_vmware_blast_sdr_out_queuing_time_us`                                | Out queueing time in VVC (us)                                        | gauge   | None   |
| `windows_vmware_blast_sdr_outbound_bandwidth_kbps`                            | Channel Send Bandwidth (Kbps)                                        | gauge   | None   |
| `windows_vmware_blast_sdr_received_bytes_total`                               | Received Bytes                                                       | counter | None   |
| `windows_vmware_blast_sdr_received_packets_total`                             | Received Packets                                                     | counter | None   |
| `windows_vmware_blast_sdr_transmitted_bytes_total`                            | Transmitted Bytes                                                    | counter | None   |
| `windows_vmware_blast_sdr_transmitted_packets_total`                          | Transmitted Packets                                                  | counter | None   |
| `windows_vmware_blast_serial_port_and_scanner_inband_bandwidth_kbps`          | Channel Receive Bandwidth (Kbps)                                     | gauge   | None   |
| `windows_vmware_blast_serial_port_and_scanner_out_queuing_time_us`            | Out queueing time in VVC (us)                                        | gauge   | None   |
| `windows_vmware_blast_serial_port_and_scanner_outbound_bandwidth_kbps`        | Channel Send Bandwidth (Kbps)                                        | gauge   | None   |
| `windows_vmware_blast_serial_port_and_scanner_received_bytes_total`           | Received Bytes                                                       | counter | None   |
| `windows_vmware_blast_serial_port_and_scanner_received_packets_total`         | Received Packets                                                     | counter | None   |
| `windows_vmware_blast_serial_port_and_scanner_transmitted_bytes_total`        | Transmitted Bytes                                                    | counter | None   |
| `windows_vmware_blast_serial_port_and_scanner_transmitted_packets_total`      | Transmitted Packets                                                  | counter | None   |
| `windows_vmware_blast_session_automatic_reconnect_count_total`                | Total number of reconnects that happened after session interruptions | counter | None   |
| `windows_vmware_blast_session_cumlative_received_bytes_over_tcp_total`        | Cumulative received bytes on the connection over TCP                 | counter | None   |
| `windows_vmware_blast_session_cumlative_received_bytes_over_udp_total`        | Cumulative received bytes on the connection over UDP                 | counter | None   |
| `windows_vmware_blast_session_cumlative_transmitted_bytes_over_tcp_total`     | Cumulative transmitted bytes on the connection over TCP              | counter | None   |
| `windows_vmware_blast_session_cumlative_transmitted_bytes_over_udp_total`     | Cumulative transmitted bytes on the connection over UDP              | counter | None   |
| `windows_vmware_blast_session_estimated_bandwidth_uplink_kbps`                | Estimated network bandwidth for uplink in Kbps                       | gauge   | None   |
| `windows_vmware_blast_session_instantaneous_received_bytes_over_tcp_total`    | Instantaneous received bytes on the connection over TCP              | gauge   | None   |
| `windows_vmware_blast_session_instantaneous_received_bytes_over_udp_total`    | Instantaneous received bytes on the connection over UDP              | gauge   | None   |
| `windows_vmware_blast_session_instantaneous_transmitted_bytes_over_tcp_total` | Instantaneous transmitted bytes on the connection over TCP           | gauge   | None   |
| `windows_vmware_blast_session_instantaneous_transmitted_bytes_over_udp_total` | Instantaneous transmitted bytes on the connection over UDP           | gauge   | None   |
| `windows_vmware_blast_session_jitter_uplink_ms`                               | Network jitter for uplink in Milliseconds                            | gauge   | None   |
| `windows_vmware_blast_session_packet_loss_uplink`                             | Network packet loss for uplink in percentage                         | gauge   | None   |
| `windows_vmware_blast_session_received_bytes_total`                           | Received Bytes                                                       | counter | None   |
| `windows_vmware_blast_session_received_packets_total`                         | Received Packets                                                     | counter | None   |
| `windows_vmware_blast_session_rtt_ms`                                         | Estimated RTT in Milliseconds                                        | gauge   | None   |
| `windows_vmware_blast_session_transmitted_bytes_total`                        | Transmitted Bytes                                                    | counter | None   |
| `windows_vmware_blast_session_transmitted_packets_total`                      | Transmitted Packets                                                  | counter | None   |
| `windows_vmware_blast_smart_card_inband_bandwidth_kbps`                       | Channel Receive Bandwidth (Kbps)                                     | gauge   | None   |
| `windows_vmware_blast_smart_card_out_queuing_time_us`                         | Out queueing time in VVC (us)                                        | gauge   | None   |
| `windows_vmware_blast_smart_card_outbound_bandwidth_kbps`                     | Channel Send Bandwidth (Kbps)                                        | gauge   | None   |
| `windows_vmware_blast_smart_card_received_bytes_total`                        | Received Bytes                                                       | counter | None   |
| `windows_vmware_blast_smart_card_received_packets_total`                      | Received Packets                                                     | counter | None   |
| `windows_vmware_blast_smart_card_transmitted_bytes_total`                     | Transmitted Bytes                                                    | counter | None   |
| `windows_vmware_blast_smart_card_transmitted_packets_total`                   | Transmitted Packets                                                  | counter | None   |
| `windows_vmware_blast_usb_inband_bandwidth_kbps`                              | Channel Receive Bandwidth (Kbps)                                     | gauge   | None   |
| `windows_vmware_blast_usb_out_queuing_time_us`                                | Out queueing time in VVC (us)                                        | gauge   | None   |
| `windows_vmware_blast_usb_outbound_bandwidth_kbps`                            | Channel Send Bandwidth (Kbps)                                        | gauge   | None   |
| `windows_vmware_blast_usb_received_bytes_total`                               | Received Bytes                                                       | counter | None   |
| `windows_vmware_blast_usb_received_packets_total`                             | Received Packets                                                     | counter | None   |
| `windows_vmware_blast_usb_transmitted_bytes_total`                            | Transmitted Bytes                                                    | counter | None   |
| `windows_vmware_blast_usb_transmitted_packets_total`                          | Transmitted Packets                                                  | counter | None   |
| `windows_vmware_blast_view_scanner_inband_bandwidth_kbps`                     | Channel Receive Bandwidth (Kbps)                                     | gauge   | None   |
| `windows_vmware_blast_view_scanner_out_queuing_time_us`                       | Out queueing time in VVC (us)                                        | gauge   | None   |
| `windows_vmware_blast_view_scanner_outbound_bandwidth_kbps`                   | Channel Send Bandwidth (Kbps)                                        | gauge   | None   |
| `windows_vmware_blast_view_scanner_received_bytes_total`                      | Received Bytes                                                       | counter | None   |
| `windows_vmware_blast_view_scanner_received_packets_total`                    | Received Packets                                                     | counter | None   |
| `windows_vmware_blast_view_scanner_transmitted_bytes_total`                   | Transmitted Bytes                                                    | counter | None   |
| `windows_vmware_blast_view_scanner_transmitted_packets_total`                 | Transmitted Packets                                                  | counter | None   |
| `windows_vmware_blast_windows_media_mmr_inband_bandwidth_kbps`                | Channel Receive Bandwidth (Kbps)                                     | gauge   | None   |
| `windows_vmware_blast_windows_media_mmr_out_queuing_time_us`                  | Out queueing time in VVC (us)                                        | gauge   | None   |
| `windows_vmware_blast_windows_media_mmr_outbound_bandwidth_kbps`              | Channel Send Bandwidth (Kbps)                                        | gauge   | None   |
| `windows_vmware_blast_windows_media_mmr_received_bytes_total`                 | Received Bytes                                                       | counter | None   |
| `windows_vmware_blast_windows_media_mmr_received_packets_total`               | Received Packets                                                     | counter | None   |
| `windows_vmware_blast_windows_media_mmr_transmitted_bytes_total`              | Transmitted Bytes                                                    | counter | None   |
| `windows_vmware_blast_windows_media_mmr_transmitted_packets_total`            | Transmitted Packets                                                  | counter | None   |


### Example metric
_This collector does not yet have explained examples, we would appreciate your help adding them!_

## Useful queries
_This collector does not yet have any useful queries added, we would appreciate your help adding them!_

## Alerting examples
_This collector does not yet have alerting examples, we would appreciate your help adding them!_
