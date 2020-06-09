package main

import "github.com/prometheus/client_golang/prometheus"

var (
	fuzzer_pid        *prometheus.GaugeVec
	cycles_done       *prometheus.GaugeVec
	execs_done        *prometheus.GaugeVec
	execs_per_sec     *prometheus.GaugeVec
	paths_total       *prometheus.GaugeVec
	paths_favored     *prometheus.GaugeVec
	paths_found       *prometheus.GaugeVec
	paths_imported    *prometheus.GaugeVec
	max_depth         *prometheus.GaugeVec
	cur_path          *prometheus.GaugeVec
	pending_favs      *prometheus.GaugeVec
	pending_total     *prometheus.GaugeVec
	variable_paths    *prometheus.GaugeVec
	stability         *prometheus.GaugeVec
	bitmap_cvg        *prometheus.GaugeVec
	unique_crashes    *prometheus.GaugeVec
	unique_hangs      *prometheus.GaugeVec
	last_path         *prometheus.GaugeVec
	last_crash        *prometheus.GaugeVec
	last_hang         *prometheus.GaugeVec
	execs_since_crash *prometheus.GaugeVec
	exec_timeout      *prometheus.GaugeVec
	slowest_exec_ms   *prometheus.GaugeVec
	peak_rss_mb       *prometheus.GaugeVec
	// missing start_time, last_update, afl_version, target_mode and command_line, and to be honest, I don't see a reason to export these values via prometheus
)

// Initializes all gauges and register them
func InitializeGauges() {
	// Set up gauges
	fuzzer_pid = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "fuzzer_pid"}, []string{"name"})
	cycles_done = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "cycles_done"}, []string{"name"})
	execs_done = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "execs_done"}, []string{"name"})
	execs_per_sec = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "execs_per_sec"}, []string{"name"})
	paths_total = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "paths_total"}, []string{"name"})
	paths_favored = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "paths_favored"}, []string{"name"})
	paths_found = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "paths_found"}, []string{"name"})
	paths_imported = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "paths_imported"}, []string{"name"})
	max_depth = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "max_depth"}, []string{"name"})
	cur_path = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "cur_path"}, []string{"name"})
	pending_favs = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "pending_favs"}, []string{"name"})
	pending_total = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "pending_total"}, []string{"name"})
	variable_paths = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "variable_paths"}, []string{"name"})
	stability = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "stability"}, []string{"name"})
	bitmap_cvg = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "bitmap_cvg"}, []string{"name"})
	unique_crashes = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "unique_crashes"}, []string{"name"})
	unique_hangs = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "unique_hangs"}, []string{"name"})
	last_path = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "last_path"}, []string{"name"})
	last_crash = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "last_crash"}, []string{"name"})
	last_hang = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "last_hang"}, []string{"name"})
	execs_since_crash = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "execs_since_crash"}, []string{"name"})
	exec_timeout = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "exec_timeout"}, []string{"name"})
	slowest_exec_ms = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "slowest_exec_ms"}, []string{"name"})
	peak_rss_mb = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "peak_rss_mb"}, []string{"name"})

	// Register gauges
	prometheus.MustRegister(fuzzer_pid)
	prometheus.MustRegister(cycles_done)
	prometheus.MustRegister(execs_done)
	prometheus.MustRegister(execs_per_sec)
	prometheus.MustRegister(paths_total)
	prometheus.MustRegister(paths_favored)
	prometheus.MustRegister(paths_found)
	prometheus.MustRegister(paths_imported)
	prometheus.MustRegister(max_depth)
	prometheus.MustRegister(cur_path)
	prometheus.MustRegister(pending_favs)
	prometheus.MustRegister(pending_total)
	prometheus.MustRegister(variable_paths)
	prometheus.MustRegister(stability)
	prometheus.MustRegister(bitmap_cvg)
	prometheus.MustRegister(unique_crashes)
	prometheus.MustRegister(unique_hangs)
	prometheus.MustRegister(last_path)
	prometheus.MustRegister(last_crash)
	prometheus.MustRegister(last_hang)
	prometheus.MustRegister(execs_since_crash)
	prometheus.MustRegister(exec_timeout)
	prometheus.MustRegister(slowest_exec_ms)
	prometheus.MustRegister(peak_rss_mb)
}
