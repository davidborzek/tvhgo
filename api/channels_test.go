package api_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/davidborzek/tvhgo/api"
	"github.com/davidborzek/tvhgo/config"
	"github.com/davidborzek/tvhgo/core"
	mock_core "github.com/davidborzek/tvhgo/mock/core"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("Channels", func() {
	var mockCtrl *gomock.Controller
	var mockChannelService *mock_core.MockChannelService
	var mockTokenService *mock_core.MockTokenService
	var sut http.Handler

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockChannelService = mock_core.NewMockChannelService(mockCtrl)
		mockTokenService = mock_core.NewMockTokenService(mockCtrl)

		mockTokenService.EXPECT().
			Validate(gomock.Any(), gomock.Any()).
			Return(&core.AuthContext{}, nil).
			AnyTimes()

		sut = api.New(&config.Config{}, mockChannelService, nil, nil, nil, nil, nil, nil, nil, nil, nil, mockTokenService, nil).
			Handler()

	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("GetChannels", func() {
		It("should return 500 on error", func() {
			req, err := http.NewRequest("GET", "/channels", nil)
			if err != nil {
				Fail(err.Error())
			}
			req.Header.Set("Authorization", "Bearer token")

			rr := httptest.NewRecorder()

			mockChannelService.EXPECT().
				GetAll(gomock.Any(), core.PaginationSortQueryParams{}).
				Return(nil, errors.New("unexpected error")).
				Times(1)

			sut.ServeHTTP(rr, req)

			Expect(rr.Code).To(Equal(http.StatusInternalServerError))
			Expect(rr.Body.String()).To(MatchJSON(`{"message":"unexpected error"}`))
		})

		DescribeTable(
			"should return 400 on invalid query parameters",
			func(limit, offset, sortKey, sortDir, errMsg string) {
				u := url.Values{}
				u.Add("limit", limit)
				u.Add("offset", offset)
				u.Add("sort_key", sortKey)
				u.Add("sort_dir", sortDir)

				req, err := http.NewRequest("GET", fmt.Sprintf("/channels?%s", u.Encode()), nil)
				if err != nil {
					Fail(err.Error())
				}
				req.Header.Set("Authorization", "Bearer token")

				rr := httptest.NewRecorder()

				mockChannelService.EXPECT().
					GetAll(gomock.Any(), gomock.Any()).
					Times(0)

				sut.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusBadRequest))
				Expect(rr.Body.String()).To(MatchJSON(errMsg))

			},
			Entry(
				"limit not a number",
				"invalid",
				"10",
				"foo",
				"asc",
				`{"message":"schema: error converting value for \"limit\""}`,
			),
			Entry(
				"limit less than 0",
				"-1",
				"10",
				"foo",
				"asc",
				`{"message":"pagination limit invalid"}`,
			),
			Entry(
				"offset not a number",
				"10",
				"invalid",
				"foo",
				"asc",
				`{"message":"schema: error converting value for \"offset\""}`,
			),
			Entry(
				"offset less than 0",
				"10",
				"-1",
				"foo",
				"asc",
				`{"message":"pagination offset invalid"}`,
			),
			Entry(
				"invalid sort direction",
				"10",
				"20",
				"foo",
				"invalid",
				`{"message":"sort direction invalid"}`,
			),
		)
	})

	It("should return a list result of channels", func() {
		u := url.Values{}
		u.Add("limit", "10")
		u.Add("offset", "20")
		u.Add("sort_key", "name")
		u.Add("sort_dir", "desc")

		req, err := http.NewRequest("GET", fmt.Sprintf("/channels?%s", u.Encode()), nil)
		if err != nil {
			Fail(err.Error())
		}
		req.Header.Set("Authorization", "Bearer token")

		rr := httptest.NewRecorder()

		channels := []*core.Channel{
			{
				ID:      "1",
				Name:    "channel1",
				Enabled: true,
				Number:  1,
				PiconID: 2,
			},
			{
				ID:      "2",
				Name:    "channel2",
				Enabled: false,
				Number:  2,
				PiconID: 3,
			},
		}

		mockChannelService.EXPECT().
			GetAll(gomock.Any(), core.PaginationSortQueryParams{
				PaginationQueryParams: core.PaginationQueryParams{
					Limit:  10,
					Offset: 20,
				},
				SortQueryParams: core.SortQueryParams{
					SortKey:       "name",
					SortDirection: "desc",
				},
			}).
			Return(channels, nil).
			Times(1)

		sut.ServeHTTP(rr, req)

		expected, _ := json.Marshal(channels)

		Expect(rr.Code).To(Equal(http.StatusOK))
		Expect(rr.Body.String()).To(MatchJSON(string(expected)))
	})

	Describe("GetChannel", func() {
		DescribeTable(
			"should return erroneous status code",
			func(targetError error, status int, msg string) {
				req, err := http.NewRequest("GET", "/channels/foobar", nil)
				if err != nil {
					Fail(err.Error())
				}
				req.Header.Set("Authorization", "Bearer token")

				rr := httptest.NewRecorder()

				mockChannelService.EXPECT().
					Get(gomock.Any(), "foobar").
					Return(nil, targetError).
					Times(1)

				sut.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(status))
				Expect(rr.Body.String()).To(MatchJSON(msg))
			},
			Entry(
				"should return 500 on error",
				errors.New("unexpected error"),
				http.StatusInternalServerError,
				`{"message":"unexpected error"}`,
			),
			Entry(
				"should return 404 on when channel is absent",
				core.ErrChannelNotFound,
				http.StatusNotFound,
				`{"message":"channel not found"}`,
			),
		)

		It("should return channel", func() {
			req, err := http.NewRequest("GET", "/channels/foobar", nil)
			if err != nil {
				Fail(err.Error())
			}
			req.Header.Set("Authorization", "Bearer token")

			rr := httptest.NewRecorder()

			channel := &core.Channel{
				ID:      "foobar",
				Name:    "channel1",
				Enabled: true,
				Number:  1,
				PiconID: 2,
			}

			mockChannelService.EXPECT().
				Get(gomock.Any(), "foobar").
				Return(channel, nil).
				Times(1)

			sut.ServeHTTP(rr, req)

			expected, _ := json.Marshal(channel)

			Expect(rr.Code).To(Equal(http.StatusOK))
			Expect(rr.Body.String()).To(MatchJSON(string(expected)))
		})
	})
})
