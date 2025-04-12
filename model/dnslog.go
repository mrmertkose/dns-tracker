package model

type DNSLog struct {
	Timestamp string `json:"timestamp"`
	SrcIP     string `json:"src_ip"`
	Query     string `json:"query"`
	QType     string `json:"query_type"`
}
