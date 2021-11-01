package main


import (
	 "X"											//import custom class
	"github.com/prometheus/client_golang/prometheus"			//import prometheus lib
	"github.com/prometheus/client_golang/prometheus/promhttp"	//import promhttp lib 
	"time"
	"net/http"	
	"github.com/labstack/echo"
	"math/rand"
	"os"
	"strings"
	cron "github.com/robfig/cron/v3"
)


//definisi variable metrics nya
var counter *prometheus.CounterVec
var histogram *prometheus.HistogramVec
var gauge *prometheus.GaugeVec
var gauge_simple *prometheus.GaugeVec

var version_app string




func init() {

	//register tiap tiap metrics
	prom.RegisterCounter()
	counter = prom.GetCounter()

	prom.RegisterHistogram()
	histogram = prom.GetHistogram()

	prom.RegisterGauge()
	gauge = prom.GetGauge()

	prom.RegisterGaugeSimple()
	gauge_simple = prom.GetGaugeSimple()

	
    	version_app = os.Getenv("VERSION")
/*
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta") 
   	scheduler := cron.New(cron.WithLocation(jakartaTime))

  	 // stop scheduler tepat sebelum fungsi berakhir
   	defer scheduler.Stop()
	
	t := time.Now().Local()
	s := t.Format("2006-01-02")
	ns := strings.Replace(s, "-", ".", -1)


	scheduler.AddFunc("0 0 1 1 *", func() { SetVersion(ns) })
*/
	//if (version_app != ""){
	// menambahkan gauge dengan label version

	//version_app = t.In(loc).Format("2021.10.01")	
	//gauge_simple.WithLabelValues(ns).Add(1)

	//}else {
	// menambahkan gauge dengan label version
	//gauge_simple.WithLabelValues("0.0.0").Add(1)
	//}
     
  }

func SetVersion(version string) {
	 gauge_simple.WithLabelValues(version).Add(1)
}

func main() {
	e := echo.New()

	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
        scheduler := cron.New(cron.WithLocation(jakartaTime))

         // stop scheduler tepat sebelum fungsi berakhir
        defer scheduler.Stop()

        t := time.Now().Local()
        s := t.Format("2006-01-02")
        ns := strings.Replace(s, "-", ".", -1)


        scheduler.AddFunc("0 1 * * *", func() { SetVersion(ns) })

	SetVersion(ns)


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

	e.Logger.Fatal(e.Start(":9001"))
}


//fungsi dummy, time track
func timeTrack(start time.Time, sleeptime int) (int64) {
	time.Sleep(time.Duration(sleeptime)*time.Second)
    elapsed := time.Since(start)
    return elapsed.Nanoseconds()
}

// dekstop/golang-prome/main.go
