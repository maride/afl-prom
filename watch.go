package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	registeredFuzzers []Fuzzer
)

// Returns the path to every fuzzer to watch
func getFuzzersToWatch() ([]string, error) {
	for i, a := range os.Args {
		if a == "--" {
			// Choose arguments after that one as the target fuzzer directories
			return os.Args[i+1:], nil
		}
	}

	// Wrong usage - construct a helpful error message
	return []string{}, fmt.Errorf("Please give at least one fuzzer directory to watch.\n%s [options...] -- /path/to/fuzzer1 /path/to/fuzzer2", os.Args[0])
}

// Registers the given paths as fuzzer directories which should be monitored
func registerFuzzers(targets []string) {
	// First, create fuzzer instances based on the directory
	for _, f := range targets {
		tmpFuzzer := CreateFuzzer(f)
		registeredFuzzers = append(registeredFuzzers, tmpFuzzer)
	}

	// Create gauges
	InitializeGauges()
}

// Watch over the fuzzer(s)
func watchFuzzers() {
	// Loop forever
	for {
		// Loop over every registered fuzzer
		for _, f := range registeredFuzzers {
			parseErr := f.ParseStatsFile()

			if parseErr != nil {
				log.Printf("Encountered error while parsing %s: %s", f, parseErr.Error())
			}
		}

		// and sleep
		time.Sleep(30 * time.Second)
	}
}
