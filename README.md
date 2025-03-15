# XGINX Load Balancer

A high-performance load balancer built from scratch in Go. This project delivers efficient HTTP routing with robust health checks and easy YAML configuration, all powered by a simple yet effective round robin scheduling algorithm.

## Features
- **HTTP Routing:** Seamlessly directs incoming HTTP requests.
- **YAML Configurable:** Customize settings with an easy-to-edit YAML file.
- **Health Checkup:** Continuously monitor backend servers to ensure uptime.
- **Round Robin Scheduling:** Evenly distribute load across multiple backends.

## Quick Start

```bash
git clone https://github.com/xSaCh/xginx.git
cd xginx
go run cmd/main.go
```

## Configuration
Sample configuration for xginx
```yaml
xginx:
  name: "Xginx" # Name of the load balancer
  host: "0.0.0.0" # Host address to bind
  port: 8080 # Port to listen on

  scheduler: "round_robin" # Load balancing algorithm: "round_robin", "least_connections", "ip_hash"

  health_check:
    enabled: true
    interval: 5 # Interval (in seconds) for health checks
    timeout: 2 # Timeout (in seconds) for a response
    retries: 3 # Number of retries before marking a server as down

  backend_servers:
    - name: "Server1"
      address: "localhost"
      port: 8081
      # weight: 1 # Optional: Load distribution weight
    - name: "Server2"
      address: "localhost"
      port: 8082
      # weight: 2 # Higher weight means more requests

  security:
    tls:
      enabled: false
      certificate: "/path/to/cert.pem"
      private_key: "/path/to/key.pem"
    rate_limiting:
      enabled: false
      requests_per_second: 100

  logging:
    level: "info" # Log levels: "debug", "info", "warn", "error"
    file: "/var/log/load_balancer.log"
```

