package helpers

import (
	"os"
	"os/exec"
)

func ClearCmd() {
	// a := 0
	// fmt.Printf("Loading")
	// for {
	// 	fmt.Printf(".")
	// 	time.Sleep(30 * time.Millisecond)
	// 	a += 20

	// 	if a == 1500 {
	// 		break
	// 	}
	// }
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
