package prom

import (
    "github.com/prometheus/client_golang/prometheus"
	//"github.com/prometheus/client_golang/prometheus/promhttp"

)

var buckets = []float64{300, 1200, 5000}

//definisiin tipe tipe metrics nya
var (
	counter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "golang",
			Name:      "my_counter",
			Help:      "This is my counter",
		},
		[]string{"code", "method", "path", "version"},
		)
	
	gauge = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: "golang",
				Name:      "my_gauge",
				Help:      "This is my counter",
			},
			[]string{"code", "method", "path", "version"},
			)

	histogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace:   "golang",
			Name:      "my_histogram",
			Help:        "How long it took to process the request",
			// ConstLabels: prometheus.Labels{"service": name},
			Buckets:     buckets,
		},
			[]string{"code", "method", "path", "version"},
		)

	summary = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: "golang",
			Name:      "my_summary",
			Help:      "This is my summary",
		},
        []string{"code", "method", "path", "version"},
    )
)


func World2() string {
    return " yyyy world"
}


func RegisterCounter(){
    prometheus.MustRegister(counter)
}

func RegisterHistogram(){
    prometheus.MustRegister(histogram)
}

func RegisterGauge(){
    prometheus.MustRegister(gauge)
}

func RegisterSummary(){
    prometheus.MustRegister(summary)
}


func GetCounter()(*prometheus.CounterVec){
    return counter
}

func GetHistogram()(*prometheus.HistogramVec){
    return histogram
}
func GetGauge()(*prometheus.GaugeVec){
    return gauge
}
func GetSummary()(*prometheus.SummaryVec){
    return summary
}