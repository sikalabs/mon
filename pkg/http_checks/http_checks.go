package http_checks

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/sikalabs/mon/pkg/alert_v2"
	"github.com/sikalabs/mon/pkg/config"
)

func RunHttpChecks(
	config config.Config,
) (error, []alert_v2.AlertV2) {
	var alerts []alert_v2.AlertV2
	for _, check := range config.HTTPChecks {
		parsedURL, err := url.Parse(check.URL)
		checkName := check.URL
		if err == nil {
			checkName = parsedURL.Host
		}
		resp, err := http.Get(check.URL)
		if err != nil {
			alerts = append(alerts, alert_v2.AlertV2{
				CheckType:    "http",
				CheckName:    checkName,
				OK:           false,
				ErrorMessage: fmt.Sprintf("ERROR http check %s: %s", check.URL, err),
			})
			continue
		}
		resp.Body.Close()
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			alerts = append(alerts, alert_v2.AlertV2{
				CheckType:    "http",
				CheckName:    checkName,
				OK:           false,
				ErrorMessage: fmt.Sprintf("ERROR http check %s: status %d\n", check.URL, resp.StatusCode),
			})
			continue
		}
		alerts = append(alerts, alert_v2.AlertV2{
			CheckType:    "http",
			CheckName:    checkName,
			OK:           true,
			ErrorMessage: "",
		})
	}
	return nil, alerts
}
