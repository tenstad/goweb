# goweb

sum(rate(process_cpu_seconds_total{app="goweb"}[10s])) by (kubernetes_pod_name)

sum(rate(http_requests_processed{app="goweb"}[10s])) by (kubernetes_pod_name)

10000
50
