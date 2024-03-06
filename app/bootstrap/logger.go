package bootstrap

import (
	"github.com/Zainal21/Ubersnap-backend-test/pkg/config"
	"github.com/Zainal21/Ubersnap-backend-test/pkg/logger"
	"github.com/Zainal21/Ubersnap-backend-test/pkg/util"
	"github.com/sirupsen/logrus"
)

func RegistryLogger(cfg *config.Config) {
	logger.Setup(logger.Config{
		Environment: util.EnvironmentTransform(cfg.AppEnv),
		Debug:       cfg.AppLoggerDebug,
		Level:       cfg.AppLoggerLevel,
		ServiceName: cfg.AppName,
		Hooks:       []logrus.Hook{}, // Add Hook list here
	})
}
