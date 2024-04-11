package api_test

import (
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/davidborzek/tvhgo/api"
	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/config"
	"github.com/davidborzek/tvhgo/core"
	mock_core "github.com/davidborzek/tvhgo/mock/core"
	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HandleAuthentication", func() {
	var mockCtrl *gomock.Controller
	var mockUserRepo *mock_core.MockUserRepository
	var mockTokenService *mock_core.MockTokenService
	var mockSessionManager *mock_core.MockSessionManager

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockUserRepo = mock_core.NewMockUserRepository(mockCtrl)
		mockTokenService = mock_core.NewMockTokenService(mockCtrl)
		mockSessionManager = mock_core.NewMockSessionManager(mockCtrl)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("returns status unauthorized", func() {
		sut := api.New(&config.Config{}, nil, nil, nil, nil, nil, nil, nil, mockUserRepo, nil, nil, nil, nil)

		middleware := sut.HandleAuthentication(nil)

		req, err := http.NewRequest("GET", "/foobar", nil)
		if err != nil {
			Fail(err.Error())
		}

		rr := httptest.NewRecorder()

		middleware.ServeHTTP(rr, req)

		Expect(rr.Code).To(Equal(http.StatusUnauthorized))
		Expect(rr.Body.String()).To(MatchJSON(`{"message":"token invalid"}`))
	})

	Describe("forward auth", func() {
		var cfg *config.Config
		var allowedProxies = []string{"10.0.0.0/24", "10.0.69.69"}

		const (
			remoteAddr = "10.0.0.1"
			username   = "foobar"
			userHeader = "Remote-User"
		)

		BeforeEach(func() {
			cfg = &config.Config{
				Auth: config.AuthConfig{
					ReverseProxy: config.ReverseProxyAuthConfig{
						Enabled:        true,
						AllowedProxies: allowedProxies,
						UserHeader:     userHeader,
					},
				},
			}
		})

		DescribeTable("remote addr is not allowed",
			func(remoteAddr string, allowedAddresses []string) {
				cfg.Auth.ReverseProxy.AllowedProxies = allowedAddresses
				sut := api.New(cfg, nil, nil, nil, nil, nil, nil, nil, mockUserRepo, nil, nil, nil, nil)
				m := sut.HandleAuthentication(nil)

				req, err := http.NewRequest("GET", "/foobar", nil)
				if err != nil {
					Fail(err.Error())
				}

				req.RemoteAddr = remoteAddr
				req.Header.Set(userHeader, username)

				rr := httptest.NewRecorder()
				m.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusUnauthorized))
				Expect(rr.Body.String()).To(MatchJSON(`{"message":"token invalid"}`))
			},
			Entry("remote addr is empty", "", allowedProxies),
			Entry("allowed addresses are empty", "10.0.0.1", []string{}),
			Entry("remote addr is invalid", "invalid", allowedProxies),
			Entry("remote ip is invalid", "invalid:1234", allowedProxies),
			Entry("remote ip is not contained in cidr", "10.69.69.69", allowedProxies),
			Entry("it fails to parse the allowed cidr", "10.69.69.69", []string{"invalid"}),
		)

		DescribeTable("remote addr is allowed and user is found",
			func(remoteAddr string) {
				sut := api.New(cfg, nil, nil, nil, nil, nil, nil, nil, mockUserRepo, nil, nil, nil, nil)

				nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					authCtx, ok := request.GetAuthContext(r.Context())
					Expect(ok).To(BeTrue())
					Expect(authCtx.ForwardAuth).To(BeTrue())
					Expect(authCtx.SessionID).To(BeNil())

					w.WriteHeader(http.StatusOK)
				})

				middleware := sut.HandleAuthentication(nextHandler)

				req, err := http.NewRequest("GET", "/foobar", nil)
				if err != nil {
					Fail(err.Error())
				}

				req.RemoteAddr = remoteAddr
				req.Header.Set(userHeader, username)

				rr := httptest.NewRecorder()

				user := &core.User{
					ID:          1,
					Username:    username,
					Email:       "foobar@example.de",
					DisplayName: "Foo Bar",
				}

				mockUserRepo.EXPECT().
					FindByUsername(req.Context(), "foobar").
					Return(user, nil).
					Times(1)

				middleware.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusOK))
			},
			Entry("remote ip is equal to allowed ip", "10.0.69.69"),
			Entry("remote ip is in allowed cidr", "10.0.0.1"),
		)

		When("remote addr is allowed", func() {
			Context("and user header is empty", func() {
				It("returns status unauthorized", func() {
					sut := api.New(cfg, nil, nil, nil, nil, nil, nil, nil, mockUserRepo, nil, nil, nil, nil)
					m := sut.HandleAuthentication(nil)

					req, err := http.NewRequest("GET", "/foobar", nil)
					if err != nil {
						Fail(err.Error())
					}

					req.RemoteAddr = remoteAddr
					req.Header.Set(userHeader, "")

					rr := httptest.NewRecorder()
					m.ServeHTTP(rr, req)

					Expect(rr.Code).To(Equal(http.StatusUnauthorized))
					Expect(rr.Body.String()).To(MatchJSON(`{"message":"token invalid"}`))
				})
			})

			Context("user is not found", func() {
				Context("and registration is disabled", func() {
					It("returns status unauthorized", func() {
						sut := api.New(cfg, nil, nil, nil, nil, nil, nil, nil, mockUserRepo, nil, nil, nil, nil)
						m := sut.HandleAuthentication(nil)

						req, err := http.NewRequest("GET", "/foobar", nil)
						if err != nil {
							Fail(err.Error())
						}

						req.RemoteAddr = remoteAddr
						req.Header.Set(userHeader, username)

						rr := httptest.NewRecorder()

						mockUserRepo.EXPECT().
							FindByUsername(req.Context(), username).
							Return(nil, nil).
							Times(1)

						m.ServeHTTP(rr, req)

						Expect(rr.Code).To(Equal(http.StatusUnauthorized))
						Expect(rr.Body.String()).To(MatchJSON(`{"message":"token invalid"}`))
					})
				})

				Context("and registration is enabled", func() {
					It("creates a new user", func() {
						cfg.Auth.ReverseProxy.AllowRegistration = true
						sut := api.New(cfg, nil, nil, nil, nil, nil, nil, nil, mockUserRepo, nil, nil, nil, nil)

						nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
							authCtx, ok := request.GetAuthContext(r.Context())

							Expect(ok).To(BeTrue())
							Expect(authCtx.ForwardAuth).To(BeTrue())
							Expect(authCtx.SessionID).To(BeNil())

							w.WriteHeader(http.StatusOK)
						})

						middleware := sut.HandleAuthentication(nextHandler)
						req, err := http.NewRequest("GET", "/foobar", nil)
						if err != nil {
							Fail(err.Error())
						}

						req.RemoteAddr = remoteAddr
						req.Header.Set(userHeader, username)

						rr := httptest.NewRecorder()
						user := &core.User{
							Username:    "foobar",
							Email:       "foobar@tvhgo.local",
							DisplayName: "foobar",
						}

						mockUserRepo.EXPECT().
							FindByUsername(req.Context(), "foobar").
							Return(nil, nil).
							Times(1)

						mockUserRepo.EXPECT().
							Create(req.Context(), user).
							Return(nil).
							Times(1)

						middleware.ServeHTTP(rr, req)
						Expect(rr.Code).To(Equal(http.StatusOK))
					})

					It("fails to create a new user", func() {
						cfg.Auth.ReverseProxy.AllowRegistration = true
						sut := api.New(cfg, nil, nil, nil, nil, nil, nil, nil, mockUserRepo, nil, nil, nil, nil)

						middleware := sut.HandleAuthentication(nil)
						req, err := http.NewRequest("GET", "/foobar", nil)
						if err != nil {
							Fail(err.Error())
						}

						req.RemoteAddr = remoteAddr
						req.Header.Set(userHeader, username)

						rr := httptest.NewRecorder()
						user := &core.User{
							Username:    "foobar",
							Email:       "foobar@tvhgo.local",
							DisplayName: "foobar",
						}

						mockUserRepo.EXPECT().
							FindByUsername(req.Context(), "foobar").
							Return(nil, nil).
							Times(1)

						mockUserRepo.EXPECT().
							Create(req.Context(), user).
							Return(errors.New("some error")).
							Times(1)

						middleware.ServeHTTP(rr, req)

						Expect(rr.Code).To(Equal(http.StatusUnauthorized))
						Expect(rr.Body.String()).To(MatchJSON(`{"message":"token invalid"}`))
					})
				})
			})

			It("fails to find user", func() {
				sut := api.New(cfg, nil, nil, nil, nil, nil, nil, nil, mockUserRepo, nil, nil, nil, nil)
				middleware := sut.HandleAuthentication(nil)

				req, err := http.NewRequest("GET", "/foobar", nil)
				if err != nil {
					Fail(err.Error())
				}

				req.RemoteAddr = remoteAddr
				req.Header.Set(userHeader, username)

				rr := httptest.NewRecorder()

				mockUserRepo.EXPECT().
					FindByUsername(req.Context(), "foobar").
					Return(nil, errors.New("some error")).
					Times(1)

				middleware.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusUnauthorized))
				Expect(rr.Body.String()).To(MatchJSON(`{"message":"token invalid"}`))
			})
		})
	})

	Describe("authorization header", func() {
		When("token is valid", func() {
			It("returns status ok", func() {
				sut := api.New(&config.Config{}, nil, nil, nil, nil, nil, nil, nil, mockUserRepo, nil, nil, mockTokenService, nil)

				nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					authCtx, ok := request.GetAuthContext(r.Context())

					Expect(ok).To(BeTrue())
					Expect(authCtx.ForwardAuth).To(BeFalse())
					Expect(authCtx.SessionID).To(BeNil())

					w.WriteHeader(http.StatusOK)
				})

				m := sut.HandleAuthentication(nextHandler)

				req, err := http.NewRequest("GET", "/foobar", nil)
				if err != nil {
					Fail(err.Error())
				}

				req.Header.Set("Authorization", "Bearer valid")

				rr := httptest.NewRecorder()

				mockTokenService.EXPECT().
					Validate(req.Context(), "valid").
					Return(&core.AuthContext{
						UserID:      1,
						SessionID:   nil,
						ForwardAuth: false,
					}, nil).
					Times(1)

				m.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusOK))
			})
		})

		When("token is invalid", func() {
			It("returns status unauthorized", func() {
				sut := api.New(&config.Config{}, nil, nil, nil, nil, nil, nil, nil, mockUserRepo, nil, nil, mockTokenService, nil)
				m := sut.HandleAuthentication(nil)

				req, err := http.NewRequest("GET", "/foobar", nil)
				if err != nil {
					Fail(err.Error())
				}

				req.Header.Set("Authorization", "Bearer invalid")

				rr := httptest.NewRecorder()

				mockTokenService.EXPECT().
					Validate(req.Context(), "invalid").
					Return(nil, core.InvalidOrExpiredTokenError{
						Reason: core.ErrTokenInvalid,
					}).
					Times(1)

				m.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusUnauthorized))
				Expect(rr.Body.String()).To(MatchJSON(`{"message":"invalid or expired token"}`))
			})
		})

		When("token service returns error", func() {
			It("returns status internal server error", func() {
				sut := api.New(&config.Config{}, nil, nil, nil, nil, nil, nil, nil, mockUserRepo, nil, nil, mockTokenService, nil)
				m := sut.HandleAuthentication(nil)

				req, err := http.NewRequest("GET", "/foobar", nil)
				if err != nil {
					Fail(err.Error())
				}

				req.Header.Set("Authorization", "Bearer invalid")

				rr := httptest.NewRecorder()

				mockTokenService.EXPECT().
					Validate(req.Context(), "invalid").
					Return(nil, errors.New("unexpected error")).
					Times(1)

				m.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusInternalServerError))
				Expect(rr.Body.String()).To(MatchJSON(`{"message":"unexpected error"}`))
			})
		})
	})

	Describe("session cookie", func() {
		var cfg *config.Config

		const (
			cookieName = "session"
			sessionID  = int64(1234)
		)

		BeforeEach(func() {
			cfg = &config.Config{
				Auth: config.AuthConfig{
					Session: config.SessionConfig{
						CookieName: cookieName,
					},
				},
			}
		})

		When("token is valid", func() {
			It("returns status ok", func() {
				sessionID := int64(1234)

				sut := api.New(cfg, nil, nil, nil, nil, nil, mockSessionManager, nil, mockUserRepo, nil, nil, mockTokenService, nil)

				nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					authCtx, ok := request.GetAuthContext(r.Context())

					Expect(ok).To(BeTrue())
					Expect(authCtx.ForwardAuth).To(BeFalse())
					Expect(*authCtx.SessionID).To(Equal(sessionID))

					w.WriteHeader(http.StatusOK)
				})

				m := sut.HandleAuthentication(nextHandler)

				req, err := http.NewRequest("GET", "/foobar", nil)
				if err != nil {
					Fail(err.Error())
				}

				req.AddCookie(&http.Cookie{
					Name:  cookieName,
					Value: "valid",
				})

				rr := httptest.NewRecorder()

				mockSessionManager.EXPECT().
					Validate(req.Context(), "valid").
					Return(&core.AuthContext{
						UserID:      1,
						SessionID:   &sessionID,
						ForwardAuth: false,
					}, nil, nil).
					Times(1)

				m.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusOK))
			})

			It("returns status ok and set the rotated token", func() {
				sessionID := int64(1234)
				rotatedToken := "rotatedToken"

				sut := api.New(cfg, nil, nil, nil, nil, nil, mockSessionManager, nil, mockUserRepo, nil, nil, mockTokenService, nil)

				nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					authCtx, ok := request.GetAuthContext(r.Context())

					Expect(ok).To(BeTrue())
					Expect(authCtx.ForwardAuth).To(BeFalse())
					Expect(*authCtx.SessionID).To(Equal(sessionID))

					w.WriteHeader(http.StatusOK)
				})

				m := sut.HandleAuthentication(nextHandler)

				req, err := http.NewRequest("GET", "/foobar", nil)
				if err != nil {
					Fail(err.Error())
				}

				req.AddCookie(&http.Cookie{
					Name:  cookieName,
					Value: "valid",
				})

				rr := httptest.NewRecorder()

				mockSessionManager.EXPECT().
					Validate(req.Context(), "valid").
					Return(&core.AuthContext{
						UserID:      1,
						SessionID:   &sessionID,
						ForwardAuth: false,
					}, &rotatedToken, nil).
					Times(1)

				m.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusOK))
				Expect(rr.Header().Get("Set-Cookie")).To(ContainSubstring("rotatedToken"))
			})
		})

		When("token is invalid", func() {
			It("returns status unauthorized", func() {
				sut := api.New(cfg, nil, nil, nil, nil, nil, mockSessionManager, nil, mockUserRepo, nil, nil, mockTokenService, nil)
				m := sut.HandleAuthentication(nil)

				req, err := http.NewRequest("GET", "/foobar", nil)
				if err != nil {
					Fail(err.Error())
				}

				req.AddCookie(&http.Cookie{
					Name:  cookieName,
					Value: "invalid",
				})

				rr := httptest.NewRecorder()

				mockSessionManager.EXPECT().
					Validate(req.Context(), "invalid").
					Return(nil, nil, core.InvalidOrExpiredTokenError{
						Reason: core.ErrTokenInvalid,
					}).
					Times(1)

				m.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusUnauthorized))
				Expect(rr.Body.String()).To(MatchJSON(`{"message":"invalid or expired token"}`))
			})
		})

		When("session manager returns error", func() {
			It("returns status internal server error", func() {
				sut := api.New(cfg, nil, nil, nil, nil, nil, mockSessionManager, nil, mockUserRepo, nil, nil, mockTokenService, nil)
				m := sut.HandleAuthentication(nil)

				req, err := http.NewRequest("GET", "/foobar", nil)
				if err != nil {
					Fail(err.Error())
				}

				req.AddCookie(&http.Cookie{
					Name:  cookieName,
					Value: "invalid",
				})

				rr := httptest.NewRecorder()

				mockSessionManager.EXPECT().
					Validate(req.Context(), "invalid").
					Return(nil, nil, errors.New("unexpected error")).
					Times(1)

				m.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusInternalServerError))
				Expect(rr.Body.String()).To(MatchJSON(`{"message":"unexpected error"}`))
			})
		})
	})
})
