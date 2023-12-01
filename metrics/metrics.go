package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	totalConnectionsMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_total_connections",
		"Number of connected client devices.",
		nil,
		nil,
	)

	connectionMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_connection",
		"Connected client device.",
		[]string{"server", "server_port", "peer", "peer_port", "started", "type", "user"},
		nil,
	)

	totalSubscriptionsMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_total_subscriptions",
		"Number of active subscriptions.",
		nil,
		nil,
	)

	subscriptionLabels = []string{
		"start",
		"hostname",
		"username",
		"client",
		"title",
		"channel",
		"service",
		"profile",
	}

	subscriptionErrorMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_subscription_errors",
		"Total errors of an active subscription.",
		subscriptionLabels,
		nil,
	)

	subscriptionInMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_subscription_in",
		"Incoming bytes of an active subscription.",
		subscriptionLabels,
		nil,
	)

	subscriptionOutMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_subscription_out",
		"Outgoing bytes of an active subscription.",
		subscriptionLabels,
		nil,
	)

	subscriptionTotalInMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_subscription_total_in",
		"Total incoming bytes of an active subscription.",
		subscriptionLabels,
		nil,
	)

	subscriptionTotalOutMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_subscription_total_out",
		"Total outgoing bytes of an active subscription.",
		subscriptionLabels,
		nil,
	)

	totalInputsMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_total_inputs",
		"Number of input devices.",
		nil,
		nil,
	)

	inputLabels = []string{"uuid", "input", "stream"}

	inputSubscriptionsMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_input_subscriptions",
		"Total subscriptions of an input.",
		inputLabels,
		nil,
	)

	inputWeightMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_input_weight",
		"Weight of an input.",
		inputLabels,
		nil,
	)

	inputSignalMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_input_signal",
		"Signal of an input.",
		inputLabels,
		nil,
	)

	inputSignalScaleMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_input_signal_scale",
		"Signal scale of an input.",
		inputLabels,
		nil,
	)

	inputBerMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_input_ber",
		"Bit error rate of an input.",
		inputLabels,
		nil,
	)

	inputSnrMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_input_snr",
		"Signal-to-noise ratio of an input.",
		inputLabels,
		nil,
	)

	inputSnrScaleMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_input_snr_scale",
		"Signal-to-noise ratio scale of an input.",
		inputLabels,
		nil,
	)

	inputUncMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_input_unc",
		"Uncorrected blocks of an input.",
		inputLabels,
		nil,
	)

	inputBpsMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_input_bps",
		"Bandwidth (bit/second) of an input.",
		inputLabels,
		nil,
	)

	inputTeMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_input_te",
		"Transport errors of an input.",
		inputLabels,
		nil,
	)

	inputCcMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_input_cc",
		"Continuity errors of an input.",
		inputLabels,
		nil,
	)

	inputEcBitMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_input_ec_bit",
		"Bit error count of an input.",
		inputLabels,
		nil,
	)

	inputTcBitMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_input_tc_bit",
		"Total bit error count of an input.",
		inputLabels,
		nil,
	)

	inputEcBlockMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_input_ec_block",
		"Block error count of an input.",
		inputLabels,
		nil,
	)

	inputTcBlockMetric = prometheus.NewDesc(
		"tvhgo_tvheadend_input_tc_block",
		"Total block error count of an input.",
		inputLabels,
		nil,
	)
)
