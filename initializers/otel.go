package initializers

import (
	"github.com/honeycombio/otel-config-go/otelconfig"
	"log"
)

var otelShutdown func()

// ConfigureOpenTelemetry sets up OpenTelemetry instrumentation
func ConfigureOpenTelemetry() {
	var err error
	otelShutdown, err = otelconfig.ConfigureOpenTelemetry()
	if err != nil {
		log.Fatalf("Error setting up OpenTelemetry: %v", err)
	}
}

// ShutdownOpenTelemetry shuts down OpenTelemetry gracefully
func ShutdownOpenTelemetry() {
	if otelShutdown != nil {
		otelShutdown()
	}
}
