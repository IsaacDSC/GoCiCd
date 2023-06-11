package cicd_test

import (
	"fmt"
	"os/exec"
)

func Running_all_tests(shCommand string) (err error) {
	fmt.Println(shCommand)
	cmd, err := exec.Command("/bin/sh", shCommand).Output()
	if err != nil {
		return
	}
	output := string(cmd)
	fmt.Println(output)
	return
}
