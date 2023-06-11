package containers

import (
	"errors"
	"fmt"
	"os/exec"
)

func RegisterContainer(containerName string, version string) (err error) {
	if len(containerName) == 0 || len(version) == 0 {
		err = errors.New("NOT-FOUND_CONTAINER_NAME OR VERSION")
		return
	}
	build_image := fmt.Sprintf("docker build -t %s:%s .", containerName, version)
	cmd, err := exec.Command(build_image).Output()
	if err != nil {
		return
	}
	output := string(cmd)
	fmt.Println(output)
	push_image := fmt.Sprintf("docker push %s%s", containerName, version)
	cm2, err := exec.Command(push_image).Output()
	if err != nil {
		return
	}
	output2 := string(cm2)
	fmt.Println(output2)
	return
}
