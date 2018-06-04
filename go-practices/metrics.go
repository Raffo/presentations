type handler struct {
	dur metrics.Histogram
}
func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func(begin time.Time) { h.dur.Observe(time.Since(begin).Seconds()) }(time.Now())
	w.Write([]byte("Hello world!"))
}
func serveMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8181", nil)
}
func main() {
	var dur metrics.Histogram = prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "myservice",
		Subsystem: "api",
		Name:      "request_duration_seconds",
		Help:      "Total time spent serving requests.",
	}, []string{})

	go serveMetrics()
	log.Printf("Server started.\n")
	http.ListenAndServe(":8080", &handler{dur})
}
