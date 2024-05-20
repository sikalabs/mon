package collect

import (
	"fmt"

	"github.com/sikalabs/mon/pkg/collect/disk"
)

type HostInfo struct {
	Disk disk.DataInfo
}

func GetHostInfo() (HostInfo, error) {
	diskInfo, err := disk.GetDiskInfo()
	if err != nil {
		return HostInfo{}, err
	}

	return HostInfo{
		Disk: diskInfo,
	}, nil
}

func PrintHostInfo() error {
	info, err := GetHostInfo()
	if err != nil {
		return err
	}
	fmt.Printf("disk_root_free_gb=%.2f\n", info.Disk.RootFreeGB)
	fmt.Printf("disk_root_free_percent=%.2f\n", info.Disk.RootFreePercent)
	return nil
}
