package metrics

import (
	"context"
	"strconv"
	"sync"

	"github.com/davidborzek/tvhgo/tvheadend"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

type TvheadendCollector struct {
	client tvheadend.Client
}

func NewTvheadendCollector(client tvheadend.Client) *TvheadendCollector {
	return &TvheadendCollector{
		client: client,
	}
}

func (c *TvheadendCollector) Describe(_ chan<- *prometheus.Desc) {}

func (c *TvheadendCollector) Collect(ch chan<- prometheus.Metric) {
	ctx := context.Background()

	var wg sync.WaitGroup
	wg.Add(3)

	go c.connectionsMetrics(ctx, ch, &wg)
	go c.subscriptionMetrics(ctx, ch, &wg)
	go c.inputsMetrics(ctx, ch, &wg)

	wg.Wait()
}

func (c *TvheadendCollector) connectionsMetrics(
	ctx context.Context,
	ch chan<- prometheus.Metric,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	var connections tvheadend.Status[tvheadend.ConnectionStatus]
	res, err := c.client.Exec(ctx, "/api/status/connections", &connections)
	if err != nil {
		log.WithError(err).
			Error("failed to collect tvheadend connections metrics")
		return
	}

	if res.StatusCode >= 400 {
		log.WithField("status", res.StatusCode).
			Error("tvheadend connections metrics request failed with erroneous status code")
		return
	}

	ch <- prometheus.MustNewConstMetric(
		totalConnectionsMetric,
		prometheus.GaugeValue,
		float64(connections.TotalCount),
	)

	for _, connection := range connections.Entries {
		ch <- prometheus.MustNewConstMetric(
			connectionMetric,
			prometheus.GaugeValue,
			float64(connection.Streaming),
			connection.Server,
			strconv.Itoa(connection.ServerPort),
			connection.Peer,
			strconv.Itoa(connection.PeerPort),
			strconv.Itoa(connection.Started),
			connection.Type,
			connection.User,
		)
	}
}

func (c *TvheadendCollector) subscriptionMetrics(
	ctx context.Context,
	ch chan<- prometheus.Metric,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	var subscriptions tvheadend.Status[tvheadend.SubscriptionStatus]
	res, err := c.client.Exec(ctx, "/api/status/subscriptions", &subscriptions)
	if err != nil {
		log.WithError(err).
			Error("failed to collect tvheadend subscriptions metrics")
		return
	}

	if res.StatusCode >= 400 {
		log.WithField("status", res.StatusCode).
			Error("tvheadend subscriptions metrics request failed with erroneous status code")
		return
	}

	ch <- prometheus.MustNewConstMetric(
		totalSubscriptionsMetric,
		prometheus.GaugeValue,
		float64(subscriptions.TotalCount),
	)

	for _, subscription := range subscriptions.Entries {
		labels := []string{
			strconv.Itoa(subscription.Start),
			subscription.Hostname,
			subscription.Username,
			subscription.Client,
			subscription.Title,
			subscription.Channel,
			subscription.Service,
			subscription.Profile,
		}

		ch <- prometheus.MustNewConstMetric(
			subscriptionErrorMetric,
			prometheus.GaugeValue,
			float64(subscription.Errors),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			subscriptionInMetric,
			prometheus.GaugeValue,
			float64(subscription.In),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			subscriptionOutMetric,
			prometheus.GaugeValue,
			float64(subscription.Out),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			subscriptionTotalInMetric,
			prometheus.GaugeValue,
			float64(subscription.TotalIn),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			subscriptionTotalOutMetric,
			prometheus.GaugeValue,
			float64(subscription.TotalOut),
			labels...,
		)
	}
}

func (c *TvheadendCollector) inputsMetrics(
	ctx context.Context,
	ch chan<- prometheus.Metric,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	var inputs tvheadend.Status[tvheadend.InputStatus]
	res, err := c.client.Exec(ctx, "/api/status/inputs", &inputs)
	if err != nil {
		log.WithError(err).
			Error("failed to collect tvheadend inputs metrics")
		return
	}

	if res.StatusCode >= 400 {
		log.WithField("status", res.StatusCode).
			Error("tvheadend inputs metrics request failed with erroneous status code")
		return
	}

	ch <- prometheus.MustNewConstMetric(
		totalInputsMetric,
		prometheus.GaugeValue,
		float64(inputs.TotalCount),
	)

	for _, input := range inputs.Entries {
		labels := []string{input.UUID, input.Input, input.Stream}

		ch <- prometheus.MustNewConstMetric(
			inputSubscriptionsMetric,
			prometheus.GaugeValue,
			float64(input.Subs),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			inputWeightMetric,
			prometheus.GaugeValue,
			float64(input.Weight),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			inputSignalMetric,
			prometheus.GaugeValue,
			float64(input.Signal),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			inputSignalScaleMetric,
			prometheus.GaugeValue,
			float64(input.SignalScale),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			inputBerMetric,
			prometheus.GaugeValue,
			float64(input.Ber),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			inputSnrMetric,
			prometheus.GaugeValue,
			float64(input.Snr),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			inputSnrScaleMetric,
			prometheus.GaugeValue,
			float64(input.SnrScale),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			inputUncMetric,
			prometheus.GaugeValue,
			float64(input.Unc),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			inputBpsMetric,
			prometheus.GaugeValue,
			float64(input.Bps),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			inputTeMetric,
			prometheus.GaugeValue,
			float64(input.Te),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			inputCcMetric,
			prometheus.GaugeValue,
			float64(input.Cc),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			inputEcBitMetric,
			prometheus.GaugeValue,
			float64(input.EcBit),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			inputTcBitMetric,
			prometheus.GaugeValue,
			float64(input.TcBit),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			inputEcBlockMetric,
			prometheus.GaugeValue,
			float64(input.EcBlock),
			labels...,
		)

		ch <- prometheus.MustNewConstMetric(
			inputTcBlockMetric,
			prometheus.GaugeValue,
			float64(input.TcBlock),
			labels...,
		)
	}
}
