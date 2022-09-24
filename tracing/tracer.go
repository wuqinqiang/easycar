package tracing

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	serviceName = "easycar"
	version     = "unknown"
)

func MustLoad(url string) {
	tp, err := tracerProvider(url)
	if err != nil {
		panic(err)
	}
	otel.SetTracerProvider(tp)
}

func Tracer(ctx context.Context, spanName string, kvs ...string) (bctx context.Context, span trace.Span) {
	tracer := otel.Tracer(serviceName)
	bctx, span = tracer.Start(ctx, spanName,
		trace.WithTimestamp(time.Now()),
		trace.WithLinks(trace.LinkFromContext(ctx)),
	)

	kvsLen := len(kvs)
	if kvsLen == 0 || kvsLen%2 != 0 {
		return
	}

	var (
		attrs []attribute.KeyValue
	)

	for i := 0; i < kvsLen; i += 2 {
		attrs = append(attrs, attribute.String(kvs[i], kvs[i+1]))
	}

	span.SetAttributes(attrs...)
	return
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
func getJaegerExporter(url string) (*jaeger.Exporter, error) {
	return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
}
