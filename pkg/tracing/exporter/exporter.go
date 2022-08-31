package exporter

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"io"
)

type Exporter struct {
	tp *trace.TracerProvider
	exp *stdouttrace.Exporter
}

// NewResource returns a resource describing this application.
func newResource(service, version string) *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			semconv.ServiceVersionKey.String(version),
		),
	)
	return r
}

// NewExporter returns a console exporter.
func NewExporter(w io.Writer, service, version string) (*Exporter, error) {
	stdExpoter, err := stdouttrace.New(
		stdouttrace.WithWriter(w),
		stdouttrace.WithPrettyPrint(),
		stdouttrace.WithoutTimestamps(),
	)
	if err != nil {
		return nil, err
	}

	return &Exporter{
		exp: stdExpoter,
		tp: trace.NewTracerProvider(
			trace.WithBatcher(stdExpoter),
			trace.WithResource(newResource(service, version)),
		),
	}, nil
}

func (e *Exporter) Run() {
	otel.SetTracerProvider(e.tp)
}

func (e *Exporter) DownExporter() {
	_ = e.tp.Shutdown(context.Background())
}

