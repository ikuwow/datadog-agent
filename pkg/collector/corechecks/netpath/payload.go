package netpath

type TracerouteHop struct {
	TTL       int     `json:"ttl"`
	IpAddress string  `json:"ip_address"`
	Host      string  `json:"host"`
	Duration  float64 `json:"duration"`
	Success   bool    `json:"success"`
}

type Traceroute struct {
	TracerouteSource string                   `json:"traceroute_source"`
	Timestamp        int64                    `json:"timestamp"`
	AgentHost        string                   `json:"agent_host"`
	DestinationHost  string                   `json:"destination_host"`
	Hops             []TracerouteHop          `json:"hops"`
	HopsByIpAddress  map[string]TracerouteHop `json:"hops_by_ip_address"`
}
