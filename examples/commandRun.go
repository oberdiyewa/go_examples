package examples

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/sh", "-c", "rm -rf testdir; rm *.txt;")

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}
