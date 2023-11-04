package resourceapi

import (
	"fmt"
	"strings"

	"github.com/shirou/gopsutil/disk"
)

func getDiskUsage() ([]int, error) {
	disks, err := disk.Partitions(false)
	if err != nil {
		fmt.Printf("get disk partitions failed, err: %v\n", err)
		return nil, err
	}
	var diskUsage []int
	for _, d := range disks {
		usage, err := disk.Usage(d.Mountpoint)
		if err != nil {
			fmt.Printf("get disk usage failed, err: %v\n", err)
			continue
		}
		if strings.Contains(d.Device, "loop") || usage.Total == 0 {
			continue
		}
		if strings.Contains(d.Device, "dev") {

			// fmt.Printf("Disk: %s\n", d.Device)

			// fmt.Printf("Mountpoint: %s\n", d.Mountpoint)
			// fmt.Printf("op: %s\n", d.Opts)

			// fmt.Printf("Total: %d B\n", usage.Total)
			// fmt.Printf("Used: %d B\n", usage.Used)
			// fmt.Printf("Free: %d B\n", usage.Free)
			// fmt.Printf("UsedPercent: %.2f%%\n", usage.UsedPercent)
			// fmt.Println("-----------------------------------------")
			diskUsage = append(diskUsage, int(usage.UsedPercent))
		}
	}
	return diskUsage, nil
}

func CheckIfAnyDeviceReachingQouta(quota int) bool {
	// if quota == nil {
	// 	return false
	// }
	diskUsage, err := getDiskUsage()
	if err != nil {
		return false
	}
	for _, d := range diskUsage {
		if d > quota {
			return true
		}
	}
	return false
}
