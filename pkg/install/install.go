package install

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/sikalabs/slu/pkg/utils/error_utils"
)

const systemdUnitPath = "/etc/systemd/system/mon.service"

func Install() {
	binaryPath, err := os.Executable()
	error_utils.HandleError(err)

	unit := fmt.Sprintf(`[Unit]
Description=mon monitoring server
After=network.target

[Service]
ExecStart=%s server
Restart=always

[Install]
WantedBy=multi-user.target
`, binaryPath)

	err = os.WriteFile(systemdUnitPath, []byte(unit), 0644)
	error_utils.HandleError(err)
	fmt.Println("written", systemdUnitPath)

	for _, args := range [][]string{
		{"systemctl", "daemon-reload"},
		{"systemctl", "enable", "mon"},
		{"systemctl", "start", "mon"},
	} {
		fmt.Println("running", args)
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		error_utils.HandleError(cmd.Run())
	}

	fmt.Println("mon service installed and started")
}
