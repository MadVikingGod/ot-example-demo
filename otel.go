package main

import (
	"context"
	"math/rand"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func OTELExample(ctx context.Context) int {
	_, span := otel.Tracer("bridge-demo").Start(ctx, "OTELExample")
	defer span.End()

	// Give the span a bit of time.
	time.Sleep(10 * time.Millisecond)
	val := rand.Intn(100)
	span.SetAttributes(attribute.Int("val", val))

	return val
}
