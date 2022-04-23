/*
 * @Description:
 * @version:
 * @Author: MoonKnight
 * @Date: 2022-02-16 23:06:02
 * @LastEditors: MoonKnight
 * @LastEditTime: 2022-02-21 16:10:40
 */

package main

import (
	"db_exporter/exporter"
	"flag"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

var (
	showVersion   = flag.Bool("version", false, "Print version information.")
	listenAddress = flag.String("web.listen-address", ":9399", "Address to listen on for web interface and telemetry.")
	metricsPath   = flag.String("web.metrics-path", "/metrics", "Path under which to expose metrics.")
	configFile    = flag.String("config.file", "db_exporter.yml", "SQL Exporter configuration file name.")
)

func main() {
	flag.Parse()

	db_exporter, err := exporter.NewExporter(*configFile)
	if err != nil {
		log.Fatalf("Error creating exporter: %s", err)
	}

	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(db_exporter)

	gatherers := prometheus.Gatherers{reg}

	// 给采集器提供一个http的访问方式
	h := promhttp.HandlerFor(gatherers,
		promhttp.HandlerOpts{
			ErrorLog:      log.NewErrorLogger(),
			ErrorHandling: promhttp.ContinueOnError,
		})

	// 访问路径也可以是其他的
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
	log.Infoln("Start server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Errorf("Error occur when start server %v", err)
		os.Exit(1)
	}
}
