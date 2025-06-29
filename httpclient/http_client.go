package httpclient

import (
	"crypto/tls"
	"log/slog"
	"net"
	"net/http"

	"github.com/unlimited-budget-ecommerce/httpz"
	"github.com/unlimited-budget-ecommerce/microservice-template/config"
)

func New(cfg *config.Cfg, targetService string) *httpz.Client {
	clientCfg, ok := cfg.Outbound[targetService]
	if !ok {
		panic("target service not found")
	}

	dialer := net.Dialer{
		Timeout:   clientCfg.DialTimeout,
		KeepAlive: clientCfg.KeepAliveInterval,
	}

	client := httpz.NewClient(
		cfg.Service.Name,
		clientCfg.BaseURL,
		httpz.WithTransport(&http.Transport{
			Proxy:                 nil,
			DialContext:           dialer.DialContext,
			DisableKeepAlives:     false,
			ForceAttemptHTTP2:     clientCfg.ForceAttemptHTTP2,
			MaxIdleConns:          clientCfg.MaxIdleConns,
			MaxIdleConnsPerHost:   clientCfg.MaxIdleConns, // one client only connects to one host
			MaxConnsPerHost:       clientCfg.MaxConns,
			IdleConnTimeout:       clientCfg.IdleConnTimeout,
			TLSHandshakeTimeout:   clientCfg.TLSHandshakeTimeout,
			ResponseHeaderTimeout: clientCfg.ResponseHeaderTimeout,
			ExpectContinueTimeout: 0, // we don't use the "Expect: 100-continue" header
			TLSClientConfig: &tls.Config{ // TODO
				InsecureSkipVerify: clientCfg.InsecureSkipVerify,
			},
		}),
		httpz.WithBaseHeaders(clientCfg.BaseHeaders),
		httpz.WithPaths(clientCfg.Paths),
		httpz.WithLogger(slog.Default()),
		httpz.WithLogMWEnabled(cfg.HttpClientMW.IsEnableLog),
		httpz.WithTracer(nil),     // TODO
		httpz.WithPropagator(nil), // TODO
		httpz.WithOtelMWEnabled(cfg.HttpClientMW.IsEnableOtel),
		httpz.WithServiceVersion(cfg.Service.Version),
		httpz.WithCircuitBreaker(
			clientCfg.CircuitBreaker.Timeout,
			clientCfg.CircuitBreaker.FailureThreshold,
			clientCfg.CircuitBreaker.SuccessThreshold,
			nil, // default to HTTP Status Code 500 and above
		),
		httpz.WithCircuitBreakerEnabled(clientCfg.CircuitBreaker.Enabled),
	)
	client.SetTimeout(clientCfg.Timeout)

	return client
}
