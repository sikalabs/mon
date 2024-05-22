package alert

import (
	"fmt"

	"github.com/sikalabs/mon/pkg/collect"
	"github.com/sikalabs/mon/pkg/config/limits"
)

type Alert struct {
	Name                string
	LimitValueFloat64   float64
	CurrentValueFloat64 float64
}

func GetAlert() ([]Alert, error) {
	alerts := []Alert{}

	limits := limits.SikaLabsLimits()
	info, err := collect.GetHostInfo()
	if err != nil {
		return alerts, err
	}

	if limits.DiskRootFreeGB > info.Disk.RootFreeGB {
		alerts = append(alerts, Alert{
			Name:                "DiskRootFreeGB",
			LimitValueFloat64:   limits.DiskRootFreeGB,
			CurrentValueFloat64: info.Disk.RootFreeGB,
		})
	}

	if limits.DiskRootFreePercent > info.Disk.RootFreePercent {
		alerts = append(alerts, Alert{
			Name:                "DiskRootFreePercent",
			LimitValueFloat64:   limits.DiskRootFreePercent,
			CurrentValueFloat64: info.Disk.RootFreePercent,
		})
	}

	return alerts, nil
}

func SprintAlerts(alerts []Alert) string {
	out := ""
	for _, alert := range alerts {
		out += fmt.Sprintf("%s\n", alert.Name)
		out += fmt.Sprintf(
			"limit=%.2f current=%.2f\n",
			alert.LimitValueFloat64,
			alert.CurrentValueFloat64,
		)
		out += "\n"
	}
	return out
}
