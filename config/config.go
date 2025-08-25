package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Cfg struct {
	OutboundServices map[string]HttpClient
	Service          Service
	Logger           Logger
	HttpClientMW     HttpClientMW
}

type Service struct {
	Name    string
	Port    string
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
	BaseHeaders           map[string]string // optional
	Paths                 map[string]string
	BaseURL               string
	CircuitBreaker        CircuitBreaker
	IdleConnTimeout       time.Duration
	MaxIdleConns          int
	MaxConns              int
	TLSHandshakeTimeout   time.Duration // optional
	ResponseHeaderTimeout time.Duration
	Timeout               time.Duration
	KeepAliveInterval     time.Duration
	DialTimeout           time.Duration
	ForceAttemptHTTP2     bool // optional
	InsecureSkipVerify    bool // optional
}

type CircuitBreaker struct {
	Timeout          time.Duration
	FailureThreshold uint32
	SuccessThreshold uint32
	Enabled          bool
}

func MustNew(path string) *Cfg {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic("error reading config: " + err.Error())
	}

	var cfg Cfg
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic("error unmarshalling config: " + err.Error())
	}

	return &cfg
}
