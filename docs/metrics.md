tvhgo can optionally export metrics of the tvheadend server in the Prometheus format.

See [Metrics config](/configuration/#metrics-config-metrics) to enable the metrics enpoint.

## Exported metrics

The following metrics will be exported:

| Name                                   | Description                                     |
| -------------------------------------- | ----------------------------------------------- |
| tvhgo_tvheadend_total_connections      | Number of connected client devices.             |
| tvhgo_tvheadend_connection             | Connected client device.                        |
| tvhgo_tvheadend_total_subscriptions    | Number of active subscriptions.                 |
| tvhgo_tvheadend_subscription_errors    | Total errors of an active subscription.         |
| tvhgo_tvheadend_subscription_in        | Incoming bytes of an active subscription.       |
| tvhgo_tvheadend_subscription_out       | Outgoing bytes of an active subscription.       |
| tvhgo_tvheadend_subscription_total_in  | Total incoming bytes of an active subscription. |
| tvhgo_tvheadend_subscription_total_out | Total outgoing bytes of an active subscription. |
| tvhgo_tvheadend_total_inputs           | Number of input devices.                        |
| tvhgo_tvheadend_input_subscriptions    | Total subscriptions of an input.                |
| tvhgo_tvheadend_input_weight           | Weight of an input.                             |
| tvhgo_tvheadend_input_signal           | Signal of an input.                             |
| tvhgo_tvheadend_input_signal_scale     | Signal scale of an input.                       |
| tvhgo_tvheadend_input_ber              | Bit error rate of an input.                     |
| tvhgo_tvheadend_input_snr              | Signal-to-noise ratio of an input.              |
| tvhgo_tvheadend_input_snr_scale        | Signal-to-noise ratio scale of an input.        |
| tvhgo_tvheadend_input_unc              | Uncorrected blocks of an input.                 |
| tvhgo_tvheadend_input_bps              | Bandwidth (bit/second) of an input.             |
| tvhgo_tvheadend_input_te               | Transport errors of an input.                   |
| tvhgo_tvheadend_input_cc               | Continuity errors of an input.                  |
| tvhgo_tvheadend_input_ec_bit           | Bit error count of an input.                    |
| tvhgo_tvheadend_input_tc_bit           | Total bit error count of an input.              |
| tvhgo_tvheadend_input_ec_block         | Block error count of an input.                  |
| tvhgo_tvheadend_input_tc_block         | Total block error count of an input.            |
