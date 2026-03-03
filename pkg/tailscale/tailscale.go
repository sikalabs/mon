package tailscale

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/sikalabs/mon/pkg/alert_v2"
	"github.com/sikalabs/mon/pkg/config"
)

type tailscaleStatus struct {
	BackendState string   `json:"BackendState"`
	AuthURL      string   `json:"AuthURL"`
	Health       []string `json:"Health"`
}

func RunTailscaleChecks(c config.Config) (error, []alert_v2.AlertV2) {
	if !c.Tailscale.Enabled {
		return nil, nil
	}

	// Capture stdout separately — tailscale exits non-zero when not Running,
	// but still writes valid JSON to stdout.
	cmd := exec.Command("tailscale", "status", "--json")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Run()

	var status tailscaleStatus
	if err := json.Unmarshal(stdout.Bytes(), &status); err != nil {
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
		if status.AuthURL != "" {
			alert.ErrorMessage += fmt.Sprintf("\nLogin URL: %s", status.AuthURL)
		}
		if len(status.Health) > 0 {
			alert.ErrorMessage += fmt.Sprintf("\nHealth: %s", strings.Join(status.Health, "; "))
		}
	}

	return nil, []alert_v2.AlertV2{alert}
}
