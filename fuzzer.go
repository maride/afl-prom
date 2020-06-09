package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

// Represents a Fuzzer and its state, as reported in the "fuzzer_stats" file
type Fuzzer struct {
	afl_banner    string
	afl_directory string
}

// Creates a fuzzer from the given target path
func CreateFuzzer(target string) Fuzzer {
	return Fuzzer{
		afl_banner:    path.Base(target),
		afl_directory: target,
	}
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

	for _, l := range fileLines {
		// Skip unparseable files
		if !strings.Contains(l, ":") {
			continue
		}

		// Convert line to key and value
		parts := strings.SplitN(l, ":", 2)
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		switch key {
		case "fuzzer_pid":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			fuzzer_pid.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "cycles_done":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			cycles_done.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "execs_done":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			execs_done.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "execs_per_sec":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			execs_per_sec.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "paths_total":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			paths_total.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "paths_favored":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			paths_favored.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "paths_found":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			paths_found.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "paths_imported":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			paths_imported.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "max_depth":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			max_depth.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "cur_path":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			cur_path.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "pending_favs":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			pending_favs.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "pending_total":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			pending_total.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "variable_paths":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			variable_paths.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "stability":
			val = strings.Replace(val, "%", "", 1)
			convVal, _ := strconv.ParseFloat(val, 64)
			stability.WithLabelValues(f.afl_banner).Set(convVal)
		case "bitmap_cvg":
			val = strings.Replace(val, "%", "", 1)
			convVal, _ := strconv.ParseFloat(val, 64)
			bitmap_cvg.WithLabelValues(f.afl_banner).Set(convVal)
		case "unique_crashes":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			unique_crashes.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "unique_hangs":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			unique_hangs.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "last_path":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			last_path.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "last_crash":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			last_crash.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "last_hang":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			last_hang.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "execs_since_crash":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			execs_since_crash.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "exec_timeout":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			exec_timeout.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "slowest_exec_ms":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			slowest_exec_ms.WithLabelValues(f.afl_banner).Set(float64(convVal))
		case "peak_rss_mb":
			convVal, _ := strconv.ParseInt(val, 10, 64)
			peak_rss_mb.WithLabelValues(f.afl_banner).Set(float64(convVal))
		}
	}

	// [+] Done parsing. Have a nice day.
	return nil
}
