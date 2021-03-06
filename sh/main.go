package main

import (
	"fmt"
	"os/exec"
)

func main() {
	command := `./pre.sh`
	cmd := exec.Command("/bin/sh", "-c", command)

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}
