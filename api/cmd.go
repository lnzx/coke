package api

import (
	"bytes"
	"os/exec"
)

func run(command string) string {
	cmd := exec.Command("sh", "-c", command)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return ""
	}
	return out.String()
}
