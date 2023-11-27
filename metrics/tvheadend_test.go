package metrics_test

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/davidborzek/tvhgo/metrics"
	mock_tvheadend "github.com/davidborzek/tvhgo/mock/tvheadend"
	"github.com/davidborzek/tvhgo/tvheadend"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"go.uber.org/mock/gomock"
)

var (
	connectionStatus = tvheadend.Status[tvheadend.ConnectionStatus]{
		Entries: []tvheadend.ConnectionStatus{
			{
				ID:         1,
				Server:     "server1",
				ServerPort: 1,
				Peer:       "peer1",
				PeerPort:   1,
				Started:    1,
				Streaming:  1,
				Type:       "someType",
				User:       "someUser",
			},
			{
				ID:         2,
				Server:     "server2",
				ServerPort: 2,
				Peer:       "peer2",
				PeerPort:   2,
				Started:    2,
				Streaming:  2,
				Type:       "someOtherType",
				User:       "someOtherUser",
			},
		},
		TotalCount: 2,
	}

	subscriptionStatus = tvheadend.Status[tvheadend.SubscriptionStatus]{
		Entries: []tvheadend.SubscriptionStatus{
			{
				ID:       1,
				Start:    1,
				Errors:   1,
				State:    "someState",
				Hostname: "someHostname",
				Username: "someUsername",
				Client:   "someClient",
				Title:    "someTitle",
				Channel:  "someChannel",
				Service:  "someService",
				Profile:  "someProfile",
				In:       1,
				Out:      2,
				TotalIn:  3,
				TotalOut: 4,
			},
			{
				ID:       2,
				Start:    2,
				Errors:   2,
				State:    "someOtherState",
				Hostname: "someOtherHostname",
				Username: "someOtherUsername",
				Client:   "someOtherClient",
				Title:    "someOtherTitle",
				Channel:  "someOtherChannel",
				Service:  "someOtherService",
				Profile:  "someOtherProfile",
				In:       5,
				Out:      6,
				TotalIn:  7,
				TotalOut: 8,
			},
		},
		TotalCount: 2,
	}

	inputStatus = tvheadend.Status[tvheadend.InputStatus]{
		Entries: []tvheadend.InputStatus{
			{
				UUID:        "someUUID",
				Input:       "someInput",
				Stream:      "someStream",
				Subs:        1,
				Weight:      2,
				Signal:      3,
				SignalScale: 4,
				Ber:         5,
				Snr:         6,
				SnrScale:    7,
				Unc:         8,
				Bps:         9,
				Te:          10,
				Cc:          11,
				EcBit:       12,
				TcBit:       13,
				EcBlock:     14,
				TcBlock:     15,
			},
			{
				UUID:        "someOtherUUID",
				Input:       "someOtherInput",
				Stream:      "someOtherStream",
				Subs:        16,
				Weight:      17,
				Signal:      18,
				SignalScale: 19,
				Ber:         20,
				Snr:         21,
				SnrScale:    22,
				Unc:         23,
				Bps:         24,
				Te:          25,
				Cc:          26,
				EcBit:       27,
				TcBit:       28,
				EcBlock:     29,
				TcBlock:     30,
			},
		},
		TotalCount: 2,
	}
)

const expectedMetricOutput = `
# HELP tvhgo_tvheadend_connection Connected client device.
# TYPE tvhgo_tvheadend_connection gauge
tvhgo_tvheadend_connection{peer="peer1",peer_port="1",server="server1",server_port="1",started="1",type="someType",user="someUser"} 1
tvhgo_tvheadend_connection{peer="peer2",peer_port="2",server="server2",server_port="2",started="2",type="someOtherType",user="someOtherUser"} 2
# HELP tvhgo_tvheadend_input_ber Bit error rate of an input.
# TYPE tvhgo_tvheadend_input_ber gauge
tvhgo_tvheadend_input_ber{input="someInput",stream="someStream",uuid="someUUID"} 5
tvhgo_tvheadend_input_ber{input="someOtherInput",stream="someOtherStream",uuid="someOtherUUID"} 20
# HELP tvhgo_tvheadend_input_bps Bandwidth (bit/second) of an input.
# TYPE tvhgo_tvheadend_input_bps gauge
tvhgo_tvheadend_input_bps{input="someInput",stream="someStream",uuid="someUUID"} 9
tvhgo_tvheadend_input_bps{input="someOtherInput",stream="someOtherStream",uuid="someOtherUUID"} 24
# HELP tvhgo_tvheadend_input_cc Continuity errors of an input.
# TYPE tvhgo_tvheadend_input_cc gauge
tvhgo_tvheadend_input_cc{input="someInput",stream="someStream",uuid="someUUID"} 11
tvhgo_tvheadend_input_cc{input="someOtherInput",stream="someOtherStream",uuid="someOtherUUID"} 26
# HELP tvhgo_tvheadend_input_ec_bit Bit error count of an input.
# TYPE tvhgo_tvheadend_input_ec_bit gauge
tvhgo_tvheadend_input_ec_bit{input="someInput",stream="someStream",uuid="someUUID"} 12
tvhgo_tvheadend_input_ec_bit{input="someOtherInput",stream="someOtherStream",uuid="someOtherUUID"} 27
# HELP tvhgo_tvheadend_input_ec_block Block error count of an input.
# TYPE tvhgo_tvheadend_input_ec_block gauge
tvhgo_tvheadend_input_ec_block{input="someInput",stream="someStream",uuid="someUUID"} 14
tvhgo_tvheadend_input_ec_block{input="someOtherInput",stream="someOtherStream",uuid="someOtherUUID"} 29
# HELP tvhgo_tvheadend_input_signal Signal of an input.
# TYPE tvhgo_tvheadend_input_signal gauge
tvhgo_tvheadend_input_signal{input="someInput",stream="someStream",uuid="someUUID"} 3
tvhgo_tvheadend_input_signal{input="someOtherInput",stream="someOtherStream",uuid="someOtherUUID"} 18
# HELP tvhgo_tvheadend_input_signal_scale Signal scale of an input.
# TYPE tvhgo_tvheadend_input_signal_scale gauge
tvhgo_tvheadend_input_signal_scale{input="someInput",stream="someStream",uuid="someUUID"} 4
tvhgo_tvheadend_input_signal_scale{input="someOtherInput",stream="someOtherStream",uuid="someOtherUUID"} 19
# HELP tvhgo_tvheadend_input_snr Signal-to-noise ratio of an input.
# TYPE tvhgo_tvheadend_input_snr gauge
tvhgo_tvheadend_input_snr{input="someInput",stream="someStream",uuid="someUUID"} 6
tvhgo_tvheadend_input_snr{input="someOtherInput",stream="someOtherStream",uuid="someOtherUUID"} 21
# HELP tvhgo_tvheadend_input_snr_scale Signal-to-noise ratio scale of an input.
# TYPE tvhgo_tvheadend_input_snr_scale gauge
tvhgo_tvheadend_input_snr_scale{input="someInput",stream="someStream",uuid="someUUID"} 7
tvhgo_tvheadend_input_snr_scale{input="someOtherInput",stream="someOtherStream",uuid="someOtherUUID"} 22
# HELP tvhgo_tvheadend_input_subscriptions Total subscriptions of an input.
# TYPE tvhgo_tvheadend_input_subscriptions gauge
tvhgo_tvheadend_input_subscriptions{input="someInput",stream="someStream",uuid="someUUID"} 1
tvhgo_tvheadend_input_subscriptions{input="someOtherInput",stream="someOtherStream",uuid="someOtherUUID"} 16
# HELP tvhgo_tvheadend_input_tc_bit Total bit error count of an input.
# TYPE tvhgo_tvheadend_input_tc_bit gauge
tvhgo_tvheadend_input_tc_bit{input="someInput",stream="someStream",uuid="someUUID"} 13
tvhgo_tvheadend_input_tc_bit{input="someOtherInput",stream="someOtherStream",uuid="someOtherUUID"} 28
# HELP tvhgo_tvheadend_input_tc_block Total block error count of an input.
# TYPE tvhgo_tvheadend_input_tc_block gauge
tvhgo_tvheadend_input_tc_block{input="someInput",stream="someStream",uuid="someUUID"} 15
tvhgo_tvheadend_input_tc_block{input="someOtherInput",stream="someOtherStream",uuid="someOtherUUID"} 30
# HELP tvhgo_tvheadend_input_te Transport errors of an input.
# TYPE tvhgo_tvheadend_input_te gauge
tvhgo_tvheadend_input_te{input="someInput",stream="someStream",uuid="someUUID"} 10
tvhgo_tvheadend_input_te{input="someOtherInput",stream="someOtherStream",uuid="someOtherUUID"} 25
# HELP tvhgo_tvheadend_input_unc Uncorrected blocks of an input.
# TYPE tvhgo_tvheadend_input_unc gauge
tvhgo_tvheadend_input_unc{input="someInput",stream="someStream",uuid="someUUID"} 8
tvhgo_tvheadend_input_unc{input="someOtherInput",stream="someOtherStream",uuid="someOtherUUID"} 23
# HELP tvhgo_tvheadend_input_weight Weight of an input.
# TYPE tvhgo_tvheadend_input_weight gauge
tvhgo_tvheadend_input_weight{input="someInput",stream="someStream",uuid="someUUID"} 2
tvhgo_tvheadend_input_weight{input="someOtherInput",stream="someOtherStream",uuid="someOtherUUID"} 17
# HELP tvhgo_tvheadend_subscription_errors Total errors of an active subscription.
# TYPE tvhgo_tvheadend_subscription_errors gauge
tvhgo_tvheadend_subscription_errors{channel="someChannel",client="someClient",hostname="someHostname",profile="someProfile",service="someService",start="1",title="someTitle",username="someUsername"} 1
tvhgo_tvheadend_subscription_errors{channel="someOtherChannel",client="someOtherClient",hostname="someOtherHostname",profile="someOtherProfile",service="someOtherService",start="2",title="someOtherTitle",username="someOtherUsername"} 2
# HELP tvhgo_tvheadend_subscription_in Incoming bytes of an active subscription.
# TYPE tvhgo_tvheadend_subscription_in gauge
tvhgo_tvheadend_subscription_in{channel="someChannel",client="someClient",hostname="someHostname",profile="someProfile",service="someService",start="1",title="someTitle",username="someUsername"} 1
tvhgo_tvheadend_subscription_in{channel="someOtherChannel",client="someOtherClient",hostname="someOtherHostname",profile="someOtherProfile",service="someOtherService",start="2",title="someOtherTitle",username="someOtherUsername"} 5
# HELP tvhgo_tvheadend_subscription_out Outgoing bytes of an active subscription.
# TYPE tvhgo_tvheadend_subscription_out gauge
tvhgo_tvheadend_subscription_out{channel="someChannel",client="someClient",hostname="someHostname",profile="someProfile",service="someService",start="1",title="someTitle",username="someUsername"} 2
tvhgo_tvheadend_subscription_out{channel="someOtherChannel",client="someOtherClient",hostname="someOtherHostname",profile="someOtherProfile",service="someOtherService",start="2",title="someOtherTitle",username="someOtherUsername"} 6
# HELP tvhgo_tvheadend_subscription_total_in Total incoming bytes of an active subscription.
# TYPE tvhgo_tvheadend_subscription_total_in gauge
tvhgo_tvheadend_subscription_total_in{channel="someChannel",client="someClient",hostname="someHostname",profile="someProfile",service="someService",start="1",title="someTitle",username="someUsername"} 3
tvhgo_tvheadend_subscription_total_in{channel="someOtherChannel",client="someOtherClient",hostname="someOtherHostname",profile="someOtherProfile",service="someOtherService",start="2",title="someOtherTitle",username="someOtherUsername"} 7
# HELP tvhgo_tvheadend_subscription_total_out Total outgoing bytes of an active subscription.
# TYPE tvhgo_tvheadend_subscription_total_out gauge
tvhgo_tvheadend_subscription_total_out{channel="someChannel",client="someClient",hostname="someHostname",profile="someProfile",service="someService",start="1",title="someTitle",username="someUsername"} 4
tvhgo_tvheadend_subscription_total_out{channel="someOtherChannel",client="someOtherClient",hostname="someOtherHostname",profile="someOtherProfile",service="someOtherService",start="2",title="someOtherTitle",username="someOtherUsername"} 8
# HELP tvhgo_tvheadend_total_connections Number of connected client devices.
# TYPE tvhgo_tvheadend_total_connections gauge
tvhgo_tvheadend_total_connections 2
# HELP tvhgo_tvheadend_total_inputs Number of input devices.
# TYPE tvhgo_tvheadend_total_inputs gauge
tvhgo_tvheadend_total_inputs 2
# HELP tvhgo_tvheadend_total_subscriptions Number of active subscriptions.
# TYPE tvhgo_tvheadend_total_subscriptions gauge
tvhgo_tvheadend_total_subscriptions 2
`

func mockTvheadendExec(
	ctx context.Context,
	path string,
	dst interface{},
	query ...tvheadend.Query,
) (*tvheadend.Response, error) {
	res := &tvheadend.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	switch path {
	case "/api/status/connections":
		{
			g := dst.(*tvheadend.Status[tvheadend.ConnectionStatus])
			g.Entries = connectionStatus.Entries
			g.TotalCount = connectionStatus.TotalCount
		}
	case "/api/status/subscriptions":
		{
			g := dst.(*tvheadend.Status[tvheadend.SubscriptionStatus])
			g.Entries = subscriptionStatus.Entries
			g.TotalCount = subscriptionStatus.TotalCount
		}
	case "/api/status/inputs":
		{
			g := dst.(*tvheadend.Status[tvheadend.InputStatus])
			g.Entries = inputStatus.Entries
			g.TotalCount = inputStatus.TotalCount
		}
	}

	return res, nil
}

func TestCollectMetrics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)

	mockClient.EXPECT().
		Exec(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(mockTvheadendExec).
		Times(3)

	c := metrics.NewTvheadendCollector(mockClient)

	if err := testutil.CollectAndCompare(c, strings.NewReader(expectedMetricOutput)); err != nil {
		t.Errorf("unexpected collecting result:\n%s", err)
	}
}
