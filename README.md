# afl-prom

## What?

*afl-prom* exposes [AFL](https://aflplus.plus/)'s `fuzzer_stats` files to be collected by [Prometheus](https://prometheus.io/)

## Why?

Monitoring your fuzzers is an important task to stay up-to-date with the progress of your fuzzers - which means: time consumed and money spent.
While many users do this by running *afl-fuzz* in `tmux` or `screen` and attach to them every now and then, I don't think that this is a good monitoring. Neither does it scale well, nor does it allow the creation of histograms or cool graphs.

This is the problem which *afl-prom* tries to solve.
It exposes the stats which are reported on the *afl-fuzz* status screen and written in the `fuzzer_stats` file of each fuzzer.
In combination with [Prometheus](https://prometheus.io/) and [Grafana](https://grafana.com), this allows state-of-the-art monitoring of all of your fuzzers.

## How?

Install [Golang](https://golang.org/), then run

`go get github.com/maride/afl-prom`

After that, you can run `afl-prom`, like this:

`afl-prom --scan-delay 30 -- /path/to/fuzzer1 /path/to/fuzzer2`

This exposes an HTTP server on port `2112`. Have a look at the `/metrics` subpage.
[Set up a Prometheus instance](https://prometheus.io/docs/prometheus/latest/getting_started/) to grab these metrics. See the example configuration below.

```
scrape_configs:
  - job_name: 'afl-prom'

    scrape_interval: 5s

    static_configs:
            - targets: ['127.0.0.1:2112']
```

Then, [set up a Grafana instance](https://grafana.com/get) instance and use Prometheus as a data source.

You're done! Have fun with your new graphs.

