package main


import (
	"github.com/prometheus/client_golang/prometheus"			//import prometheus lib
	"github.com/prometheus/client_golang/prometheus/promhttp"	//import promhttp lib 
	"time"
	"net/http"	
	"github.com/labstack/echo"
	"strings"
	"math/rand"
	"fmt"
	//"strconv"
)




//definisiin tipe tipe metrics nya
var (
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
	

)




func init() {

	prometheus.MustRegister(gauge_simple)
     
  }

func SetVersion(version string) {
	 gauge_simple.WithLabelValues(version).Set(1)
 fmt.Printf(time.Now().Format("2006-01-02 15:04:05") +  " telah dijalankan.\n")
}

func main() {
	e := echo.New()

	
	
	




	e.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "Hello this is base path")
	})

	e.GET("/tsel", func(c echo.Context) error {

		
		return c.String(http.StatusOK, "hello this is tsel page")
	})

	e.GET("/indosat", func(c echo.Context) error {



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

	


	// menampilkan metrics
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	e.Logger.Fatal(e.Start(":9002"))
}




// dekstop/golang-prome/main.go
