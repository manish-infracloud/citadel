# Prometheus

## Contents

- [Prometheus](#prometheus)
  - [PromQL](#promql)
  - [Prometheus SDK](#prometheus-SDK)
  - [Grafana](#grafana)
  - [Grafana Dashboards](#grafana-dashboards)
  - [Prometheus operator](#operator)
- [Reference links](#reference-links)


## Prometheus

- Observability & Monitoring
- Metrics
- Timeseries Databases
- Configuring Prometheus
  - Scrape Config
  - Recording Rules
  - Alerting Rules

### Assignments

- [ ] Download and run Prometheus locally or in a Docker container

## PromQL

- Prometheus metrics
- Types
  - Counters
  - Gauges
  - Histograms
  - Summary
- PromQL syntax & operators
- Instant Vector vs Range Vectors
- PromQL functions
- Querying prometheus using the HTTP API


### Assignments

- [ ] Install prometheus using helm & access prometheus using kubectl port-forward.
- [ ] Write a query to find cpu usage
- [ ] Write a query to find memory usage

## Prometheus SDK

- Instrumentation
- Client libraries
  - Go
  - Scala/Java
  - Python
  - Ruby
- PushGateway
- Exporters

### Assignments

- [ ] Write a Go program to expose Prometheus metrics using the Go client

## Grafana

- Running Grafana
- Dashboards
- DataSources
- Notification Channels & alerts
- Configuration
- Provisioning

### Assignments

- [ ] Deploy Grafana and add a Prometheus instance as a datasource
- [ ] Create a Slack webhook and add it as a notification channel in Grafana
- [ ] Provision a datasource on Grafana startup

## Grafana Dashboards

- Creating Dashboards
- Query editor
- Query inspector
- Panel Types (Graphs, Singlestats, ...)
- Template variables
- Exporting and importing dashboards

### Assignments

- [ ] Create a dashboard with multiple panels showing memory, cpu usage
- [ ] Export a dashboard as json and import it into another grafana instance
- [ ] Provision a dashboard on grafana startup

## Reference Links
- [Prometheus docs](https://prometheus.io/docs/introduction/overview/)
- [Grafana docs](https://grafana.com/docs/)
- [Understanding CPU usage](https://www.robustperception.io/understanding-machine-cpu-usage)
- [Infracloud blog - Monitoring Kubernetes with Prometheus operator](https://www.infracloud.io/monitoring-kubernetes-prometheus/) [(git repo)](https://github.com/kanuahs/prometheus-operator-demo)
- [Using the Prometheus Operator Helm chart with a values.yaml override](https://github.com/kanuahs/sock-shop-prometheus-operator)
