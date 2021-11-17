package prom

import (
    "github.com/prometheus/client_golang/prometheus"
	//"github.com/prometheus/client_golang/prometheus/promhttp"

)

var buckets = []float64{100, 300, 500, 700, 900, 1200}

//definisiin tipe tipe metrics nya
var (
	counter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "deoxys",
			Name:      "my_counter",
			Help:      "This is my counter",
		},
		[]string{"code", "method", "path", "client"},
		)
	
	gauge = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: "deoxys",
				Name:      "my_gauge",
				Help:      "This is my counter",
			},
			[]string{"code", "method", "path", "client"},
			)
	gauge_simple = prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Namespace: "deoxys",
					Name:      "version_app",
					Help:      "App Version",
				},
				[]string{"version"},
				)

	histogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace:   "deoxys",
			Name:      "my_histogram",
			Help:        "How long it took to process the request",
			// ConstLabels: prometheus.Labels{"service": name},
			Buckets:     buckets,
		},
			[]string{"code", "method", "path", "client"},
		)

	summary = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: "deoxys",
			Name:      "my_summary",
			Help:      "This is my summary",
		},
        []string{"code", "method", "path", "client"},
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

func RegisterGaugeSimple(){
    prometheus.MustRegister(gauge_simple)
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

func GetGaugeSimple()(*prometheus.GaugeVec){
    return gauge_simple
}

func GetSummary()(*prometheus.SummaryVec){
    return summary
}
