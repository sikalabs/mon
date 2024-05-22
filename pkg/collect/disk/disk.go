package disk

import "syscall"

type DataInfo struct {
	RootFreeGB      float64
	RootFreePercent float64
}

func GetDiskInfo() (DataInfo, error) {
	freeSpaceGB, freeSpacePercentage, err := getFreeSpaceGB("/")
	if err != nil {
		return DataInfo{}, err
	}

	return DataInfo{
		RootFreeGB:      freeSpaceGB,
		RootFreePercent: freeSpacePercentage,
	}, nil
}

func getFreeSpaceGB(path string) (float64, float64, error) {
	var stat syscall.Statfs_t

	// Get the filesystem statistics
	err := syscall.Statfs(path, &stat)
	if err != nil {
		return 0, 0, err
	}

	// Calculate the free space in bytes
	freeSpaceBytes := stat.Bavail * uint64(stat.Bsize)

	// Calculate the total space in bytes
	totalSpaceBytes := stat.Blocks * uint64(stat.Bsize)

	// Calculate the used space in bytes
	usedSpaceBytes := totalSpaceBytes - (stat.Bfree * uint64(stat.Bsize))

	// Convert the free space to gigabytes
	const bytesInGB = 1024 * 1024 * 1024
	freeSpaceGB := float64(freeSpaceBytes) / bytesInGB

	// Calculate the used free percentage
	freeSpacePercentage := (1 - (float64(usedSpaceBytes) / float64(totalSpaceBytes))) * 100

	return freeSpaceGB, freeSpacePercentage, nil
}
