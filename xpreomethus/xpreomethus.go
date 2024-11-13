package xpreomethus

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RunPromethus() {
	temp := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "home_temperature_celsius",
			Help: "The current temperature in degrees Celsius.",
		},
		// 指定标签名称
		[]string{"task_id"},
	)

	// 注册到全局默认注册表中
	prometheus.MustRegister(temp)
	temp.WithLabelValues("1111").Add(10)
	// 针对不同标签值设置不同的指标值

	// 暴露自定义的指标
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8081", nil)
	fmt.Printf("Http Err:", err)
}
