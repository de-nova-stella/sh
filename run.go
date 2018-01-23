package sh

import (
	"fmt"
	"os/exec"
	"strings"
)

func Run(cmd *exec.Cmd) (string, *exec.ExitError) {
	fmt.Printf("[INFO] running %s...\n", cmd.Args)
	stdout, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			fmt.Println("[ERROR] stderr:", strings.Trim(string(exitErr.Stderr), "\n"))
			return string(stdout), exitErr
		}
		fmt.Println("[ERROR] command failed without exit error")
		return string(stdout), nil
	}
	return string(stdout), nil
}
