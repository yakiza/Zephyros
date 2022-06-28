package instrumentation

import "net/http"

func ExposeMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2112", nil)
	if err != nil {
		return
	}
}
