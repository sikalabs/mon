package alert_v2

import "fmt"

type AlertV2 struct {
	CheckType    string
	CheckName    string
	OK           bool
	ErrorMessage string
}

func SprintAlertsV2(alerts []AlertV2) string {
	var okOrErr string
	out := ""
	for _, alert := range alerts {
		if alert.OK {
			okOrErr = "[OK] "
		} else {
			okOrErr = "[ERR]"
		}

		out += fmt.Sprintf("%s %s/%s\n", okOrErr, alert.CheckType, alert.CheckName)
		if len(alert.ErrorMessage) > 0 {
			out += alert.ErrorMessage
			out += "\n"
		}
		out += "\n"
	}
	return out
}
