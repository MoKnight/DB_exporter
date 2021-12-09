/*
 * @Description:
 * @version:
 * @Author: MoonKnight
 * @Date: 2021-12-05 19:55:58
 * @LastEditors: MoonKnight
 * @LastEditTime: 2021-12-07 00:25:02
 */

package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(worker_process)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
