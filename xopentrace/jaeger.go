package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

type Trace struct {
	TraceId string `json:"trace_id"`
	SpanId string `json:"span_id"`
}

func initJaeger(service string) io.Closer {
	cfg := jaegercfg.Configuration{
		// 将采样频率设置为1，每个span都记录，方便查看测试结果
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
			// 将span发往jaeger-collector的服务地址
			CollectorEndpoint: "http://localhost:14268/api/traces",
		},
	}
	closer, err := cfg.InitGlobalTracer(service, jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return closer
}

func Foo(ctx context.Context) {
	// 开始一个span, 设置span的operation_name=Foo
	span, _ := opentracing.StartSpanFromContext(ctx, "Foo")
	span.Context()
	fun := func(k, v string) bool {
		fmt.Printf("k:%v, v:%v", k, v)
		return true
	}
	span.Context().ForeachBaggageItem(fun)
	defer span.Finish()
	// 将context传递给Bar
	go Bar(ctx)
	// 模拟执行耗时
	time.Sleep(1 * time.Second)
}
func Bar(ctx context.Context) {
	// 开始一个span，设置span的operation_name=Bar
	span, ctx := opentracing.StartSpanFromContext(ctx, "Bar")
	sp := span.Context().(jaeger.SpanContext)
	//defer span.Finish()
	// 模拟执行耗时
	time.Sleep(2 * time.Second)

	// 假设Bar发生了某些错误
	err := errors.New("something wrong")
	span.LogFields(
		log.String("event", "error"),
		log.String("message", err.Error()),
	)
	span.SetTag("trace_id", sp.TraceID())
	span.SetTag("error", true)
}

func main() {
	closer := initJaeger("in-process")
	defer closer.Close()
	// 获取jaeger tracer
	t := opentracing.GlobalTracer()
	// 建立root span
	sp := t.StartSpan("in-process-service")
	// main执行完结束这个span
	defer sp.Finish()
	// 将span传递给Foo
	ctx := opentracing.ContextWithSpan(context.Background(), sp)
	Foo(ctx)
}