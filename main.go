package main


import (
	prom "./prom" 												//import custom class
	"github.com/prometheus/client_golang/prometheus"			//import prometheus lib
	"github.com/prometheus/client_golang/prometheus/promhttp"	//import promhttp lib 
	"time"
	"net/http"	
	"github.com/labstack/echo"
	"math/rand"
)


//definisi variable metrics nya
var counter *prometheus.CounterVec
var histogram *prometheus.HistogramVec
var gauge *prometheus.GaugeVec




func init() {

	//register tiap tiap metrics
	prom.RegisterCounter()
	counter = prom.GetCounter()

	prom.RegisterHistogram()
	histogram = prom.GetHistogram()

	prom.RegisterGauge()
	gauge = prom.GetGauge()
	
    // menambahkan gauge dengan label version
	gauge.WithLabelValues("200", "GET", "gauge_version", "v.0.1.0").Add(1) 
  }

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {


		//add counter untuk path /
		counter.WithLabelValues("200", "GET", "/", "v.0.1.0").Add(1)
		//add histogram untuk path /
		elapsed_time  := timeTrack(time.Now(),rand.Intn(5))
		histogram.WithLabelValues("200", "GET", "/", "v.0.1.0").Observe(float64(elapsed_time) / 1000000) //milisecons
			


		return c.String(http.StatusOK, "Hello this is base path")
	})

	e.GET("/satu", func(c echo.Context) error {


		//add counter untuk path /satu
		counter.WithLabelValues("200", "GET", "/satu", "v.0.1.0").Add(1)
		//add histogram untuk path /satu
		elapsed_time  := timeTrack(time.Now(),rand.Intn(5))
		histogram.WithLabelValues("200", "GET", "/satu", "v.0.1.0").Observe(float64(elapsed_time) / 1000000) //milisecons
			

		
		return c.String(http.StatusOK, "hello this is satu page")
	})

	e.GET("/dua", func(c echo.Context) error {

		//add counter untuk path /dua
		counter.WithLabelValues("200", "GET", "/dua", "v.0.1.0").Add(1)
		//add histogram untuk path /dua
		elapsed_time  := timeTrack(time.Now(),rand.Intn(5))
		histogram.WithLabelValues("200", "GET", "/dua", "v.0.1.0").Observe(float64(elapsed_time) / 1000000) //milisecons
			


		return c.String(http.StatusOK, "hello this is dua page")
	})


	// menampilkan metrics
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	e.Logger.Fatal(e.Start(":1323"))
}


//fungsi dummy, time track
func timeTrack(start time.Time, sleeptime int) (int64) {
	time.Sleep(time.Duration(sleeptime)*time.Second)
    elapsed := time.Since(start)
    return elapsed.Nanoseconds()
}

// dekstop/golang-prome/main.go