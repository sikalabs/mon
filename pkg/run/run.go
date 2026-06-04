package run

import (
	"fmt"
	"os"

	"github.com/sikalabs/mon/pkg/alert"
	"github.com/sikalabs/mon/pkg/alert_v2"
	"github.com/sikalabs/mon/pkg/config"
	"github.com/sikalabs/mon/pkg/http_checks"
	"github.com/sikalabs/mon/pkg/notify"
	"github.com/sikalabs/mon/pkg/tailscale"
	"github.com/sikalabs/slu/pkg/utils/error_utils"
)

func RunOrDie() {
	config := config.LoadConfig()
	RunWithConfigOrDie(config)
}

func RunWithConfigOrDie(config config.Config) {
	hostname, err := os.Hostname()
	error_utils.HandleError(err)

	alerts, err := alert.GetAlert(config)
	error_utils.HandleError(err)

	// Alerts (v1) to AlertsV2
	var alertsV2 []alert_v2.AlertV2
	for _, alert := range alerts {
		alertsV2 = append(alertsV2, alert_v2.AlertV2{
			CheckType: "host-info",
			CheckName: alert.Name,
			OK:        alert.CurrentValueFloat64 > alert.LimitValueFloat64,
			ErrorMessage: fmt.Sprintf(
				"current=%.2f is lower than limit=%.2f",
				alert.CurrentValueFloat64,
				alert.LimitValueFloat64,
			),
		})
	}

	_, http_alerts := http_checks.RunHttpChecks(config)
	alertsV2 = append(alertsV2, http_alerts...)

	_, tailscale_alerts := tailscale.RunTailscaleChecks(config)
	alertsV2 = append(alertsV2, tailscale_alerts...)

	body := alert_v2.SprintAlertsV2(alertsV2)

	fmt.Print(body)

	isErrorInAlerts := false
	for _, alert := range alertsV2 {
		if !alert.OK {
			isErrorInAlerts = true
			continue
		}
	}

	if isErrorInAlerts == false && config.Notifications.SendOK == false {
		return
	}

	err = notify.SendEmailNotification(config, hostname, body)
	error_utils.HandleError(err)

	err = notify.SendTelegramNotification(config, hostname, body)
	error_utils.HandleError(err)
}
