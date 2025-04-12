package writer

import (
	"dns-tracker/model"
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type JSONWriter struct {
	logDir string
}

func NewJSONWriter(logDir string) *JSONWriter {
	return &JSONWriter{logDir: logDir}
}

func (w *JSONWriter) Write(log model.DNSLog) error {
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
