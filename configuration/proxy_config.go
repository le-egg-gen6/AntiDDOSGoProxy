package configuration

import "time"

type ProxyConfig struct {
	//Server setting
	Listen          string        `yaml:"listen"`
	TargetURL       string        `yaml:"target_url"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	IdleTimeout     time.Duration `yaml:"idle_timeout"`
	LogDir          string        `yaml:"log_dir"`
	EnableProfiling bool          `yaml:"enable_profiling"`

	//Rate Limiting
	RateLimit int `yaml:"rate_limit"`
	RateBurst int `yaml:"rate_burst"`

	//Connection Management
	MaxConnsPerIP     int           `yaml:"max_conns_per_ip"`
	MaxConnsTotal     int           `yaml:"max_conns_total"`
	ConnectionTimeout time.Duration `yaml:"connection_timeout"`

	// Request limits
	MaxBodySize   int64 `yaml:"max_body_size"`
	MaxHeaderSize int   `yaml:"max_header_size"`

	// Client tracking
	ClientTimeout   time.Duration `yaml:"client_timeout"`
	CleanupInterval time.Duration `yaml:"cleanup_interval"`
	ReqCooldown     time.Duration `yaml:"req_cooldown"`

	// Blocking policy
	BlockThreshold int           `yaml:"block_threshold"`
	BlockDuration  time.Duration `yaml:"block_duration"`

	// Layer 3/4 protection
	TCPSynCookies   bool `yaml:"tcp_syn_cookies"`
	ICMPRateLimit   int  `yaml:"icmp_rate_limit"`
	EnableConntrack bool `yaml:"enable_conntrack"`

	// SSL/TLS
	EnableTLS   bool   `yaml:"enable_tls"`
	TLSCertFile string `yaml:"tls_cert_file"`
	TLSKeyFile  string `yaml:"tls_key_file"`

	// IP filtering
	AllowedIPs       []string `yaml:"allowed_ips"`
	BlockedIPs       []string `yaml:"blocked_ips"`
	BlockedCountries []string `yaml:"blocked_countries"`

	// Traffic analysis
	EnableHeuristicDDoS bool    `yaml:"enable_heuristic_ddos"`
	AnomalyThreshold    float64 `yaml:"anomaly_threshold"`
	TrafficSamplingRate int     `yaml:"traffic_sampling_rate"`

	// HTTP protection
	BlockInvalidHostHeaders bool     `yaml:"block_invalid_host_headers"`
	ValidateHTTPMethods     bool     `yaml:"validate_http_methods"`
	AllowedHTTPMethods      []string `yaml:"allowed_http_methods"`

	// Proxy behavior
	HideServerHeader   bool `yaml:"hide_server_header"`
	AddSecurityHeaders bool `yaml:"add_security_headers"`

	// Advanced features
	EnableIPReputation bool   `yaml:"enable_ip_reputation"`
	IPReputationDB     string `yaml:"ip_reputation_db"`
}
