package sys

import (
	"math"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/wailsapp/wails"
)

// Stats .
type Stats struct {
	log *wails.CustomLogger
}

// CPUUsage .
type CPUUsage struct {
	Average int `json:"avg"`
}

// WailsInit .
func (s *Stats) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Stats")
	return nil
}

// GetCPUUsage .
func (s *Stats) GetCPUUsage() *CPUUsage {
	percent, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		s.log.Errorf("unable to get cpu stats: %s", err.Error())
		return nil
	}
	return &CPUUsage{
		Average: int(math.Round(percent[0])),
	}
}
