package internal

import "time"

type Config struct {
	Service      Service
	Logger       Logger
	HttpClientMW HttpClientMW
	Outbound     map[string]HttpClient
}

type Service struct {
	Name    string
	Env     string
	Version string
}

type Logger struct {
	Level            string
	IsEnableSource   bool
	IsEnableReplacer bool
}

type HttpClientMW struct {
	IsEnableLog  bool
	IsEnableOtel bool
}

type HttpClient struct {
	BaseURL               string
	BaseHeaders           map[string]string // optional
	DialTimeout           time.Duration
	KeepAliveInterval     time.Duration
	MaxConns              int
	MaxIdleConns          int
	IdleConnTimeout       time.Duration
	TLSHandshakeTimeout   time.Duration // optional
	ResponseHeaderTimeout time.Duration
	Timeout               time.Duration
	Paths                 map[string]string
	CircuitBreaker        CircuitBreaker
	ForceAttemptHTTP2     bool // optional
	InsecureSkipVerify    bool // optional
}

type CircuitBreaker struct {
	Timeout          time.Duration
	FailureThreshold uint32
	SuccessThreshold uint32
	Enabled          bool
}

func NewConfig() *Config {
	// TODO: viper unmarshal to struct
	return &Config{}
}
