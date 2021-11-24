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
	"fmt"
	//"strconv"
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
	 gauge_simple.WithLabelValues(version).Set(1)
 fmt.Printf(time.Now().Format("2006-01-02 15:04:05") +  " telah dijalankan.\n")
}

func main() {
	e := echo.New()

	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
    scheduler := cron.New(cron.WithLocation(jakartaTime))

    // stop scheduler tepat sebelum fungsi berakhir
    defer scheduler.Stop()

    // t := time.Now().Local()
    // s := t.Format("2006-01-02")
	 
    // ns := strings.Replace(s, "-", ".", -1)
		t := time.Now().Local()
		t = t.AddDate(0, 0, 1)
		s := t.Format("2006-01-02-15:00")
		 
		ns := strings.Replace(s, "-", ".", -1)
		ns = strings.Replace(ns, ":", ".", -1)


    scheduler.AddFunc("0 */4 * * *", func() { SetVersion(ns) })

	// start scheduler
    go scheduler.Start()
	
	


	//SetVersion(ns)


	e.GET("/", func(c echo.Context) error {

		// t := time.Now().Local()
		// s := t.Format("2006-01-02-15:00")
		 
		// ns := strings.Replace(s, "-", ".", -1)
		// ns = strings.Replace(ns, ":", ".", -1)
		// fmt.Printf(ns)
		//add counter untuk path /
		counter.WithLabelValues("200", "GET", "/", "index").Add(1)
		//add histogram untuk path /
		//elapsed_time  := timeTrack(time.Now(),rand.Intn(5))
		rand_num := rand.Intn(1000)
		//histogram.WithLabelValues("200", "GET", "/", "index").Observe(float64(elapsed_time) / 1000000) //milisecons
		histogram.WithLabelValues("200", "GET", "/", "index").Observe(float64(rand_num)) //milisecons
			


		return c.String(http.StatusOK, "Hello this is base path")
	})

	e.GET("/tsel", func(c echo.Context) error {


		//add counter untuk path /satu
		counter.WithLabelValues("200", "GET", "/tsel", "telkomsel").Add(1)
		//add histogram untuk path /satu
		//elapsed_time  := timeTrack(time.Now(),rand.Intn(5))
		rand_num := rand.Intn(900)
		histogram.WithLabelValues("200", "GET", "/tsel", "telkomsel").Observe(float64(rand_num)) //milisecons
			

		
		return c.String(http.StatusOK, "hello this is tsel page")
	})

	e.GET("/indosat", func(c echo.Context) error {

		//add counter untuk path /dua
		counter.WithLabelValues("200", "GET", "/indosat", "indosat").Add(1)
		//add histogram untuk path /dua
		//elapsed_time  := timeTrack(time.Now(),rand.Intn(5))
		rand_num := rand.Intn(2000)
		//histogram.WithLabelValues("200", "GET", "/indosat", "indosat").Observe(float64(elapsed_time) / 1000000) //milisecons
		histogram.WithLabelValues("200", "GET", "/indosat", "indosat").Observe(float64(rand_num)) //milisecons
			


		return c.String(http.StatusOK, "Hello this is indosat")
	})


	e.GET("/setversion", func(c echo.Context) error {

		day := rand.Intn(100)

		t := time.Now().Local()
		t = t.AddDate(0, 0, day)
		s := t.Format("2006-01-02-15:00")
		 
		ns := strings.Replace(s, "-", ".", -1)
		ns = strings.Replace(ns, ":", ".", -1)

		SetVersion(ns)

		return c.String(http.StatusOK, ns)
	})

	e.GET("/frontend_metrics", func(c echo.Context) error {

		
		
		data :="# HELP deoxys_version_app App Version \n" +"# TYPE deoxys_version_app gauge \n"+ "frontend_deoxys_version_app{version='2021.11.24.13.00'} 1"
   		 return c.String(http.StatusOK, data)

		
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
