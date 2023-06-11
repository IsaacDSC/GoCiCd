package gocicd

import (
	"fmt"
	"log"
	"os"

	"github.com/IsaacDSC/GoCiCd/cicd_test"
	"github.com/IsaacDSC/GoCiCd/containers"
	"github.com/IsaacDSC/GoCiCd/deploy"
	"github.com/IsaacDSC/GoCiCd/versioning"
)

func Started_cli() {
	args := os.Args
	versioning := versioning.Versioning{}
	versioning.GetVersion()

	if len(args) == 2 {
		switch args[1] {
		case "deploy":
			err := cicd_test.Running_all_tests(versioning.SH_Test)
			if err != nil {
				log.Fatal("[ * ] Error in test...\n")
				return
			}
			err = containers.RegisterContainer(versioning.ContainerName, versioning.Version)
			if err != nil {
				log.Fatal("[ * ] Error: ", err.Error())
				return
			}
			err = deploy.DeployNomad(versioning.ContainerName, versioning.Version)
			if err != nil {
				log.Fatal("[ * ] Error-Deploy-Nomad: ", err.Error())
				return
			}
			fmt.Println("[ * ] Deployed in Nomad")
			versioning.SetVersionSoftware()
			fmt.Println("[ * ] Set Version Deploy:", versioning.Version)
			return
		case "tests":
			cicd_test.Running_all_tests(versioning.SH_Test)
			return
		default:
			log.Fatal("not-found-parameter")
		}
	}
}
