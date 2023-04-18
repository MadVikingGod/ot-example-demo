package main

import (
	"context"
	"time"

	"math/rand"

	"github.com/opentracing/opentracing-go"
)

func OTExample(ctx context.Context) int {
	span, _ := opentracing.StartSpanFromContext(ctx, "OTExample")

	// Give the span a bit of time.
	time.Sleep(10 * time.Millisecond)
	val := rand.Intn(100)
	span.SetTag("val", val)

	span.Finish()
	return val
}
