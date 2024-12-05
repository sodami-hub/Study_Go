package main

import (
	"fmt"
	"net/http"

	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var PORT = ":1234"

// 새로운 couter 변수를 정의하고 원하는 옵션을 설정했다. Namespace 필드는 메트릭들을 그룹화하는 데 사용하므로 매우 중요하다.
var counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "mtsouk",
		Name:      "my_counter",
		Help:      "This is my counter",
	})

var gauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "mtsouk",
		Name:      "my_gauge",
		Help:      "This is my gauge",
	})

var histogram = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Namespace: "mtsouk",
		Name:      "my_histogram",
		Help:      "This is my histogram",
	})

var summary = prometheus.NewSummary(
	prometheus.SummaryOpts{
		Namespace: "mtsouk",
		Name:      "my_summary",
		Help:      "This is my summary",
	})

// 위와 같이 메트릭 변수들을 정의하는 것만으로는 충분하지 않고 해당 변수들을 프로메테우스에서 수집할 수 있게 등록해야 한다.
func main() {
	rand.Seed(time.Now().Unix())

	prometheus.MustRegister(counter)
	prometheus.MustRegister(gauge)
	prometheus.MustRegister(histogram)
	prometheus.MustRegister(summary)

	// 아래 고루틴은 무한 for루프의 존재로 인해 웹 서버를 실행하는 동안 종료되지 않고 계속 실행된다.
	// 또한 고루틴에서 time.Sleep()을 사용했기 때문에 메트릭 값은 2초마다 업데이트된다.(메트릭 값으로는 난수를 사용했다.)
	go func() {
		for {
			counter.Add(rand.Float64() * 5)
			gauge.Add(rand.Float64()*15 - 5)
			histogram.Observe(rand.Float64() * 10)
			summary.Observe(rand.Float64() * 10)
			time.Sleep(2 * time.Second)
		}
	}()

	// 핸들러를 직접 작성하지 않고 "github.com/prometheus/client_golang/prometheus/promhttp" 패키지의 promhttp.Handler() 핸들러 함수를 이용한다.
	// 이를 이용하면 직접 코드를 사용하지 않아도 된다. 하지만 http.Handle()을 이용해 핸들러 함수를 등록해야 된다.
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Listening to port", PORT)
	fmt.Println(http.ListenAndServe(PORT, nil))
}
