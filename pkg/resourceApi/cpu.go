package resourceapi

import (
	"math"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func GetCpuUtilization() (int, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return int(math.Round(percent[0])), nil

}
