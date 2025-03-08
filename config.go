package xginx

import (
	"os"

	"github.com/xSaCh/xginx/pkg/schedulers"
	"gopkg.in/yaml.v3"
)

type Config struct {
	LoadBalancer LoadBalancerConfig `yaml:"xginx"`
}

type LoadBalancerConfig struct {
	Name           string                        `yaml:"name"`
	Host           string                        `yaml:"host"`
	Port           int                           `yaml:"port"`
	Scheduler      schedulers.SchedulerAlgorithm `yaml:"scheduler"`
	HealthCheck    HealthCheckConfig             `yaml:"health_check"`
	BackendServers []BackendServer               `yaml:"backend_servers"`
	Security       SecurityConfig                `yaml:"security"`
	Logging        LoggingConfig                 `yaml:"logging"`
	StickySessions StickySessionsConfig          `yaml:"sticky_sessions"`
}

type HealthCheckConfig struct {
	Enabled  bool `yaml:"enabled"`
	Interval int  `yaml:"interval"`
	Timeout  int  `yaml:"timeout"`
	Retries  int  `yaml:"retries"`
}

type BackendServer struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
	Weight  int    `yaml:"weight"`
}

type SecurityConfig struct {
	TLS          TLSConfig          `yaml:"tls"`
	RateLimiting RateLimitingConfig `yaml:"rate_limiting"`
}

type TLSConfig struct {
	Enabled     bool   `yaml:"enabled"`
	Certificate string `yaml:"certificate"`
	PrivateKey  string `yaml:"private_key"`
}

type RateLimitingConfig struct {
	Enabled           bool `yaml:"enabled"`
	RequestsPerSecond int  `yaml:"requests_per_second"`
}

type LoggingConfig struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

type StickySessionsConfig struct {
	Enabled bool   `yaml:"enabled"`
	Type    string `yaml:"type"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	setDefaults(&config)
	return &config, nil
}

func setDefaults(config *Config) {
	if config.LoadBalancer.Name == "" {
		config.LoadBalancer.Name = "xginx"
	}
	if config.LoadBalancer.Host == "" {
		config.LoadBalancer.Host = "0.0.0.0"
	}
	if config.LoadBalancer.Port == 0 {
		config.LoadBalancer.Port = 8080
	}
	if config.LoadBalancer.Scheduler == "" {
		config.LoadBalancer.Scheduler = schedulers.SCHEDULER_ROUND_ROBIN
	}
	if config.LoadBalancer.HealthCheck.Interval == 0 {
		config.LoadBalancer.HealthCheck.Interval = 5
	}
	if config.LoadBalancer.HealthCheck.Timeout == 0 {
		config.LoadBalancer.HealthCheck.Timeout = 2
	}
	if config.LoadBalancer.HealthCheck.Retries == 0 {
		config.LoadBalancer.HealthCheck.Retries = 3
	}

	if config.LoadBalancer.Logging.Level == "" {
		config.LoadBalancer.Logging.Level = "debug"
	}
	if config.LoadBalancer.Logging.File == "" {
		config.LoadBalancer.Logging.File = "xginx.log"
	}
}
