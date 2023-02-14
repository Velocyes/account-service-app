package helpers

import (
	"os"
	"os/exec"
)

func ClearCmd() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
