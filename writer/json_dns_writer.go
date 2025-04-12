package writer

import (
	"dns-tracker/model"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type JSONDnsWriter struct {
	logDir string
}

func NewJSONDNSWriter(logDir string) *JSONDnsWriter {
	return &JSONDnsWriter{logDir: logDir}
}

func (w *JSONDnsWriter) DnsWrite(log model.DNSLog) error {
	err := os.MkdirAll(w.logDir, 0755)
	if err != nil {
		return fmt.Errorf("log dir not created: %w", err)
	}

	filename := time.Now().Format("2006-01-02") + ".json"
	fullPath := filepath.Join(w.logDir, filename)

	f, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := json.Marshal(log)
	if err != nil {
		return err
	}

	_, err = f.WriteString(string(data) + "\n")
	return err
}
