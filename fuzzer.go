package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Represents a Fuzzer and its state, as reported in the "fuzzer_stats" file
type Fuzzer struct {
	fuzzer_pid        prometheus.Gauge
	cycles_done       prometheus.Gauge
	execs_done        prometheus.Gauge
	execs_per_sec     prometheus.Gauge
	paths_total       prometheus.Gauge
	paths_favored     prometheus.Gauge
	paths_found       prometheus.Gauge
	paths_imported    prometheus.Gauge
	max_depth         prometheus.Gauge
	cur_path          prometheus.Gauge
	pending_favs      prometheus.Gauge
	pending_total     prometheus.Gauge
	variable_paths    prometheus.Gauge
	stability         prometheus.Gauge
	bitmap_cvg        prometheus.Gauge
	unique_crashes    prometheus.Gauge
	unique_hangs      prometheus.Gauge
	last_path         prometheus.Gauge
	last_crash        prometheus.Gauge
	last_hang         prometheus.Gauge
	execs_since_crash prometheus.Gauge
	exec_timeout      prometheus.Gauge
	slowest_exec_ms   prometheus.Gauge
	peak_rss_mb       prometheus.Gauge
	afl_banner        string
	afl_directory     string
	// missing start_time, last_update, afl_version, target_mode and command_line, and to be honest, I don't see a reason to export these values via prometheus
}

// Initialises all gauges. Needs afl_banner to be set properly before (and doesn't check that on its own)
func (f *Fuzzer) initGauges() {
	log.Printf("%s %s", f.afl_directory, f.afl_banner)
	f.fuzzer_pid = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_fuzzer_pid"})
	f.cycles_done = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_cycles_done"})
	f.execs_done = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_execs_done"})
	f.execs_per_sec = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_execs_per_sec"})
	f.paths_total = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_paths_total"})
	f.paths_favored = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_paths_favored"})
	f.paths_found = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_paths_found"})
	f.paths_imported = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_paths_imported"})
	f.max_depth = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_max_depth"})
	f.cur_path = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_cur_path"})
	f.pending_favs = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_pending_favs"})
	f.pending_total = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_pending_total"})
	f.variable_paths = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_variable_paths"})
	f.stability = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_stability"})
	f.bitmap_cvg = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_bitmap_cvg"})
	f.unique_crashes = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_unique_crashes"})
	f.unique_hangs = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_unique_hangs"})
	f.last_path = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_last_path"})
	f.last_crash = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_last_crash"})
	f.last_hang = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_last_hang"})
	f.execs_since_crash = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_execs_since_crash"})
	f.exec_timeout = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_exec_timeout"})
	f.slowest_exec_ms = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_slowest_exec_ms"})
	f.peak_rss_mb = promauto.NewGauge(prometheus.GaugeOpts{Name: f.afl_banner + "_peak_rss_mb"})
}

// Parses the "fuzzer_stats" file present in the given directory, and updates the gauges of the Fuzzer instance accordingly
func (f *Fuzzer) ParseStatsFile() error {
	// Open file and read it
	path := fmt.Sprintf("%s%cfuzzer_stats", f.afl_directory, os.PathSeparator)
	fileBytes, fileReadErr := ioutil.ReadFile(path)
	if fileReadErr != nil {
		return fileReadErr
	}

	// Split file contents into lines
	fileLines := strings.Split(string(fileBytes), "\n")

	// So we want to walk over the file two times.
	//  1) If it is unset: read "afl_banner", set it, and run initGauges()
	//  2) Set every other value to its corresponding gauge.
	bannerSet := f.afl_banner != "" // 1)
	gaugesSet := false // 2)
	for !(bannerSet && gaugesSet) {
		// Iterate over every line
		for _, l := range fileLines {
			// Skip empty lines
			if l == "" {
				continue
			}

			// Convert line to key and value
			parts := strings.SplitN(l, ":", 2)
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])
			
			// 1)
			if !bannerSet {
				if key == "afl_banner" {
					f.afl_banner = val
					bannerSet = true
					f.initGauges()
					break
				} else {
					continue
				}
			}
			
			// 2)
			switch key {
			case "fuzzer_pid":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.fuzzer_pid.Set(float64(convVal))
			case "cycles_done":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.cycles_done.Set(float64(convVal))
			case "execs_done":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.execs_done.Set(float64(convVal))
			case "execs_per_sec":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.execs_per_sec.Set(float64(convVal))
			case "paths_total":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.paths_total.Set(float64(convVal))
			case "paths_favored":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.paths_favored.Set(float64(convVal))
			case "paths_found":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.paths_found.Set(float64(convVal))
			case "paths_imported":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.paths_imported.Set(float64(convVal))
			case "max_depth":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.max_depth.Set(float64(convVal))
			case "cur_path":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.cur_path.Set(float64(convVal))
			case "pending_favs":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.pending_favs.Set(float64(convVal))
			case "pending_total":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.pending_total.Set(float64(convVal))
			case "variable_paths":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.variable_paths.Set(float64(convVal))
			case "stability":
				val = strings.Replace(val, "%", "", 1)
				convVal, _ := strconv.ParseFloat(val, 64)
				f.stability.Set(convVal)
			case "bitmap_cvg":
				val = strings.Replace(val, "%", "", 1)
				convVal, _ := strconv.ParseFloat(val, 64)
				f.bitmap_cvg.Set(convVal)
			case "unique_crashes":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.unique_crashes.Set(float64(convVal))
			case "unique_hangs":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.unique_hangs.Set(float64(convVal))
			case "last_path":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.last_path.Set(float64(convVal))
			case "last_crash":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.last_crash.Set(float64(convVal))
			case "last_hang":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.last_hang.Set(float64(convVal))
			case "execs_since_crash":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.execs_since_crash.Set(float64(convVal))
			case "exec_timeout":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.exec_timeout.Set(float64(convVal))
			case "slowest_exec_ms":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.slowest_exec_ms.Set(float64(convVal))
			case "peak_rss_mb":
				convVal, _ := strconv.ParseInt(val, 10, 64)
				f.peak_rss_mb.Set(float64(convVal))
			}

			gaugesSet = true
		}
	}

	// [+] Done parsing. Have a nice day.
	return nil
}
