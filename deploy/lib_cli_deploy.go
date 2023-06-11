package deploy

import (
	"fmt"
	"os/exec"
)

const deploy_path = "./lib/cli/scripts/deployment/deploy.sh"

func DeployNomad(containerImage string, containerVersion string) (err error) {
	fmt.Println("\n[ * ] RUNNING DEPLOY WITH NOMAD, AWAIT RESPONSE ....")
	docker_image := fmt.Sprintf("%s:%s", containerImage, containerVersion)
	fmt.Println("\n[ * ] IMAGE UPLOADING", docker_image)
	cmd, err := exec.Command("/bin/sh", deploy_path, docker_image).Output()
	if err != nil {
		return
	}
	output := string(cmd)
	fmt.Println(output)
	return
}
