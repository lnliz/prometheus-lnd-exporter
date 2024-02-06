package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func getEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}

var (
	// Set during go build
	version   string
	gitCommit string

	// maxMsgRecvSize is the largest message our client will receive. We
	// set this to ~50Mb atm.
	maxMsgRecvSize = grpc.MaxCallRecvMsgSize(1 * 1024 * 1024 * 50)
)

func main() {

	// Defaults values
	var (
		defaultNamespace     = getEnv("NAMESPACE", "lnd")
		defaultListenAddress = getEnv("LISTEN_ADDRESS", ":9113")
		defaultMetricsPath   = getEnv("TELEMETRY_PATH", "/metrics")
		defaultRpcAddr       = getEnv("RPC_ADDR", "localhost:10009")
		defaultTLSCertPath   = getEnv("TLS_CERT_PATH", "/root/.lnd")
		defaultMacaroonPath  = getEnv("MACAROON_PATH", "")
		defaultGoMetrics, _  = strconv.ParseBool(getEnv("GO_METRICS", "false"))
	)

	// Command-line flags
	var (
		namespace = flag.String("namespace", defaultNamespace,
			"The namespace or prefix to use in the exported metrics. The default value can be overwritten by NAMESPACE environment variable.")
		listenAddr = flag.String("web.listen-address", defaultListenAddress,
			"An address to listen on for web interface and telemetry. The default value can be overwritten by LISTEN_ADDRESS environment variable.")
		metricsPath = flag.String("web.telemetry-path", defaultMetricsPath,
			"A path under which to expose metrics. The default value can be overwritten by TELEMETRY_PATH environment variable.")
		rpcAddr = flag.String("rpc.addr", defaultRpcAddr,
			"Lightning node RPC host. The default value can be overwritten by RPC_HOST environment variable.")
		tlsCertPath = flag.String("lnd.tls-cert-path", defaultTLSCertPath,
			"The path to the tls certificate. The default value can be overwritten by TLS_CERT_PATH environment variable.")
		macaroonPath = flag.String("lnd.macaroon-path", defaultMacaroonPath,
			"The path to the read only macaroon. The default value can be overwritten by MACAROON_PATH environment variable.")
		goMetrics = flag.Bool("go-metrics", defaultGoMetrics,
			"Enable process and go metrics from go client library. The default value can be overwritten by GO_METRICS environmental variable.")
	)

	flag.Parse()
	log.Printf("Lightning Prometheus Exporter Version=%v GitCommit=%v", version, gitCommit)

	defaultTimeout := 15 * time.Second

	registry := prometheus.NewRegistry()
	registry.MustRegister(
		NewLightningExporter(
			*namespace,
			*rpcAddr,
			*tlsCertPath, *macaroonPath,
			defaultTimeout, true,
		))

	if *goMetrics {
		registry.MustRegister(collectors.NewGoCollector())
		registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	}

	http.Handle(*metricsPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Lightning Exporter</title></head>
			<body>
			<h1>Lightning Exporter</h1>
			<p><a href='/metrics'>Metrics</a></p>
			</body>
			</html>`))
	})

	log.Printf("ListenAndServe %s \n", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
