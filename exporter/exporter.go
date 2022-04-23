/*
 * @Description:
 * @version:
 * @Author: MoonKnight
 * @Date: 2022-02-17 11:15:53
 * @LastEditors: MoonKnight
 * @LastEditTime: 2022-02-21 16:09:57
 */

package exporter

import (
	"strconv"
	"sync"

	"db_exporter/config"

	"github.com/prometheus/client_golang/prometheus"
)

type exporter struct {
	name      string
	config    *config.Config
	queryDesc map[string]*prometheus.Desc
}

// Config implements Exporter.
func (e *exporter) Config() *config.Config {
	return e.config
}

// NewExporter returns a new Exporter with the provided config.
func NewExporter(configFile string) (*exporter, error) {
	c, err := config.Load(configFile)
	if err != nil {
		return nil, err
	}

	tempmap := make(map[string]*prometheus.Desc, len(c.QUERYS()))
	for num, q := range c.QUERYS() {
		tempmap[strconv.Itoa(num)] = prometheus.NewDesc("query"+strconv.Itoa(num), q, []string{"result"}, nil)
	}

	return &exporter{
		name:      "db_exporter",
		config:    c,
		queryDesc: tempmap,
	}, nil
}

// Collect implements Collector.
func (e *exporter) Collect(ch chan<- prometheus.Metric) {
	var wg sync.WaitGroup
	wg.Add(len(e.Config().QUERYS()))
	for num, q := range e.Config().QUERYS() {
		go func(q string) {
			defer wg.Done()
			Exec(e.queryDesc[strconv.Itoa(num)], e.Config().DSN(), q, ch)
		}(q)
	}
	// Only return once all queries have been processed
	wg.Wait()
}

// Describe implements prometheus.Collector.
func (e *exporter) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range e.queryDesc {
		ch <- m
	}
}
