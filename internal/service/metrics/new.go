package metrics

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"word/internal/entities"
)

type DB interface {
	//AddLog should get Type of log, Data of log and Time of log, and returns error
	AddLog(string, string, time.Time) error
	GetLogs(string) ([]entities.Log, error)
}

type MetricsService struct {
	db DB
}

func New(db DB) *MetricsService {
	return &MetricsService{db}
}

func (m MetricsService) Visit(Status int, Path, Method, Duration string) error {
	data := map[string]string{
		"status":   fmt.Sprint(Status),
		"path":     Path,
		"method":   Method,
		"duration": Duration,
	}
	buf, err := json.Marshal(data)
	if err != nil {
		return errors.New("Json marshaling failed")
	}
	return m.db.AddLog("visit", string(buf), time.Now())
}

func (m MetricsService) VisitLogs() ([]entities.Log, error) {
	return m.db.GetLogs("visit")
}
