package main

import (
	"fmt"
	"os"
	"time"

	"github.com/opentracing/opentracing-go/log"
	"github.com/xlzhangkeke/opentracing-tutorial/lib/tracing"
)

func main() {
	start := time.Now()
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}
	helloTo := os.Args[1]

	tracer, closer := tracing.Init("data-service")

	defer closer.Close()

	span := tracer.StartSpan("handler")
	defer span.Finish()
	span.SetTag("hello-to", helloTo)

	helloStr := fmt.Sprintf("Hello, %s!", helloTo)
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)
	time.Sleep(time.Second)
	println(helloStr)
	span.LogKV("event", "println", "elapsed", time.Now().Sub(start))
}
