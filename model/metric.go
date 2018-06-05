package modelMetric

type MetricItem struct {
	Endpoint    string      `json:"endpoint"` // hostname or IP addr
	Metric      string      `json:"metric"`
	CounterType string      `json:"counterType"` // "GAUGE" for keep original value, "COUNTER" for diff of values(ps)
	Value       interface{} `json:"value"`
	Tags        string      `json:"tags"`      // Example "port=3306,k=v"
	Timestamp   int64       `json:"timestamp"` // Unix timestamp
	Step        int64       `json:"step"`
}
