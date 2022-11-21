package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/xlzhangkeke/opentracing-tutorial/lib/tracing"
)

func main() {
	start := time.Now()
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}
	helloTo := os.Args[1]

	tracer, closer := tracing.Init("data-service3")

	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	span := tracer.StartSpan("handler")
	defer span.Finish()
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	span.SetTag("hello-to", helloTo)

	helloStr := formatString(ctx, helloTo)
	printHello(ctx, helloStr)
	time.Sleep(time.Second)
	span.LogKV("elapsed", time.Now().Sub(start))
}

func formatString(ctx context.Context, helloTo string) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "formatString")
	defer span.Finish()

	helloStr := fmt.Sprintf("Hello, %s!", helloTo)
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)

	return helloStr
}

func printHello(ctx context.Context, helloStr string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "printHello")
	defer span.Finish()

	println(helloStr)
	span.LogKV("event", "println")
}
