package tracing

import (
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
	"log/slog"
)

func Init(logger *slog.Logger, serviceName string) {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		//Reporter: &config.ReporterConfig{
		//	LocalAgentHostPort: "127.0.0.1:6831",
		//},
	}

	_, err := cfg.InitGlobalTracer(serviceName)
	if err != nil {
		logger.Error("failed to init tracing", zap.Error(err))
	}
}
