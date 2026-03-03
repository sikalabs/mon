package tailscale

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/sikalabs/mon/pkg/alert_v2"
	"github.com/sikalabs/mon/pkg/config"
)

type tailscaleStatus struct {
	BackendState string `json:"BackendState"`
}

func RunTailscaleChecks(c config.Config) (error, []alert_v2.AlertV2) {
	if !c.Tailscale.Enabled {
		return nil, nil
	}

	out, err := exec.Command("tailscale", "status", "--json").Output()
	if err != nil {
		return nil, []alert_v2.AlertV2{{
			CheckType:    "tailscale",
			CheckName:    "status",
			OK:           false,
			ErrorMessage: fmt.Sprintf("failed to run tailscale status: %s", err),
		}}
	}

	var status tailscaleStatus
	if err := json.Unmarshal(out, &status); err != nil {
		return nil, []alert_v2.AlertV2{{
			CheckType:    "tailscale",
			CheckName:    "status",
			OK:           false,
			ErrorMessage: fmt.Sprintf("failed to parse tailscale status: %s", err),
		}}
	}

	ok := status.BackendState == "Running"
	alert := alert_v2.AlertV2{
		CheckType: "tailscale",
		CheckName: "status",
		OK:        ok,
	}
	if !ok {
		alert.ErrorMessage = fmt.Sprintf("BackendState is %q, expected \"Running\"", status.BackendState)
	}

	return nil, []alert_v2.AlertV2{alert}
}
