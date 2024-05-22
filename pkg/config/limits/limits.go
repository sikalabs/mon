package limits

type Limits struct {
	DiskRootFreeGB      float64
	DiskRootFreePercent float64
}

func SikaLabsLimits() Limits {
	return Limits{
		DiskRootFreeGB:      5,
		DiskRootFreePercent: 10,
	}
}
