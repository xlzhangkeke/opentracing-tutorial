# OpenTracing Tutorial - Go

## Prerequisites
````
docker run --rm -p 6831:6831/udp -p 6832:6832/udp -p 16686:16686 jaegertracing/all-in-one:1.7 --log-level=debug

docker run -d -p 5775:5775/udp -p 16686:16686 jaegertracing/all-in-one:latest

docker run -d -p 5775:5775/udp -p 16686:16686 -p 14250:14250 -p 14268:14268 jaegertracing/all-in-one:latest

docker run --name jaegertracing -d -p 5775:5775/udp -p 16686:16686 -p 14250:14250 -p 14268:14268 jaegertracing/all-in-one:latest

docker exec -it jaegertracing /bin/sh
netstat -anp
````
view http://10.95.84.100:16686/search

## Installing

The tutorials are using CNCF Jaeger (https://github.com/jaegertracing/jaeger) as the tracing backend,

```
mkdir -p $GOPATH/src/github.com/yurishkuro/
cd $GOPATH/src/github.com/yurishkuro/
git clone https://github.com/yurishkuro/opentracing-tutorial.git opentracing-tutorial
```

After that, install the dependencies:

```
cd $GOPATH/src/github.com/yurishkuro/opentracing-tutorial/go
make install
```

The rest of the commands in the Go tutorials should be executed relative to this directory.

## Lessons

* [Lesson 01 - Hello World](./lesson01)
  * Instantiate a Tracer
  * Create a simple trace
  * Annotate the trace
* [Lesson 02 - Context and Tracing Functions](./lesson02)
  * Trace individual functions
  * Combine multiple spans into a single trace
  * Propagate the in-process context
* [Lesson 03 - Tracing RPC Requests](./lesson03)
  * Trace a transaction across more than one microservice
  * Pass the context between processes using `Inject` and `Extract`
  * Apply OpenTracing-recommended tags
* [Lesson 04 - Baggage](./lesson04)
  * Understand distributed context propagation
  * Use baggage to pass data through the call graph
* [Extra Credit](./extracredit)
  * Use existing open source instrumentation
