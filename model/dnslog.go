package model

import "time"

type DNSLog struct {
	Timestamp time.Time `json:"timestamp"`
	SrcIP     string    `json:"src_ip"`
	Query     string    `json:"query"`
	QType     string    `json:"query_type"`
}
