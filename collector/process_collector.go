/*
 * @Description:
 * @version:
 * @Author: MoonKnight
 * @Date: 2021-12-05 21:48:46
 * @LastEditors: MoonKnight
 * @LastEditTime: 2021-12-07 00:36:09
 */

package collector

import (
	"../track"
	"github.com/prometheus/client_golang/prometheus/"
)

type Process_tracker struct {
	process_name string
	statesDesc   *prometheus.Desc
}

// Describe implements prometheus.Collector.
func (p *Process_tracker) Describe(ch chan<- *prometheus.Desc) {
	ch <- p.statesDesc
}

func NewProcess_tracker(worker_process string) *Process_tracker {
	return &Process_tracker{
		process_name: worker_process,
		statesDesc: prometheus.NewDesc(
			"process",
			"Number of process.",
			[]string{"process"},
			prometheus.Labels{"process_name": worker_process},
		)}
}

func (p *Process_tracker) Collect(ch chan<- prometheus.Metric) {
	tmp_statesDesc := track.GetProcessNum()
	for host, source := range tmp_statesDesc {
		ch <- prometheus.MustNewConstMetric(
			p.statesDesc,
			prometheus.GaugeValue,
			source,
			host,
		)
	}
}
