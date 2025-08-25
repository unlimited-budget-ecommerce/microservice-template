package logger

import (
	"os"

	"github.com/unlimited-budget-ecommerce/logz"
	"github.com/unlimited-budget-ecommerce/microservice-template/config"
)

func Init(cfg *config.Cfg) {
	logz.Init(
		cfg.Service.Name,
		logz.WithWriter(os.Stdout),
		logz.WithSourceEnabled(cfg.Logger.IsEnableSource),
		logz.WithLevel(cfg.Logger.Level),
		logz.WithReplacerEnabled(cfg.Logger.IsEnableReplacer),
		logz.WithServiceVersion(cfg.Service.Version),
		logz.WithEnv(cfg.Service.Env),
	)
}
