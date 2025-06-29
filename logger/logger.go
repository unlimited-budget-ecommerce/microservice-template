package logger

import (
	"log/slog"
	"os"
	"strings"

	"github.com/unlimited-budget-ecommerce/logz"
	"github.com/unlimited-budget-ecommerce/microservice-template/config"
)

func Init(cfg *config.Cfg) {
	logz.Init(
		cfg.Service.Name,
		logz.WithWriter(os.Stdout),
		logz.WithSourceEnabled(cfg.Logger.IsEnableSource),
		logz.WithLevel(cfg.Logger.Level),
		logz.WithReplacer(func(groups []string, a slog.Attr) slog.Attr {
			switch strings.ToLower(a.Key) {
			}
			return a
		}),
		logz.WithReplacerEnabled(cfg.Logger.IsEnableReplacer),
		logz.WithServiceVersion(cfg.Service.Version),
		logz.WithEnv(cfg.Service.Env),
	)
}
