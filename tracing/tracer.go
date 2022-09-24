package tracing

import (
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

const (
	serviceName = "easycar"
	version     = "unknown"
)

func getJaegerExporter(url string) (*jaeger.Exporter, error) {
	return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
}

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	exp, err := getJaegerExporter(url)
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),

		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
			semconv.ServiceVersionKey.String(version),
		)),
	)
	return tp, err
}

func Tracer() trace.Tracer {
	return otel.Tracer(serviceName)
}

func MustLoad(url string) {
	tp, err := tracerProvider(url)
	if err != nil {
		panic(err)
	}
	otel.SetTracerProvider(tp)
}
