package tracing

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/stats"

	otelgrpc "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
)

func WithTracingInterceptors() []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.StatsHandler(newServerHandler()),
	}
}

func DialOptionsWithTracing() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithStatsHandler(newServerHandler()),
	}
}

func newClientHandler() stats.Handler {
	return otelgrpc.NewClientHandler(otelgrpc.WithTracerProvider(otel.GetTracerProvider()))
}

func newServerHandler() stats.Handler {
	return otelgrpc.NewServerHandler(otelgrpc.WithTracerProvider(otel.GetTracerProvider()))
}
