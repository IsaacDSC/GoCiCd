package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type SettingsFile struct {
	AppName               string `json:"appName"`
	Version               string `json:"version"`
	ContainerName         string `json:"containerName"`
	ShTest                string `json:"shTest"`
	UserRegisterContainer string `json:"userRegisterContainer"`
	Hcl                   string `json:"hcl"`
}

func (this_settings *SettingsFile) GetSettings() {
	buff, err := os.ReadFile("./.gocid.json")
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = json.Unmarshal(buff, this_settings)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("\n[ * ] Settings File: %+v", this_settings)
}

func (this_settings *SettingsFile) Validate() {
	if len(this_settings.AppName) == 0 {
		log.Fatalf("[ * ] Required parameter in settings [ AppName ]")
	}
	if len(this_settings.ContainerName) == 0 {
		log.Fatalf("[ * ] Required parameter in settings [ ContainerName ]")
	}
	if len(this_settings.Version) == 0 {
		log.Fatalf("[ * ] Required parameter in settings [ Version ]")
	}
	if len(this_settings.Hcl) == 0 {
		log.Fatalf("[ * ] Required parameter in settings [ Hcl ]")
	}
}

// BUILD APP
func (this_settings *SettingsFile) BuildApp() {
	fmt.Println("\n[ * ] Execute build app")
	build_image := fmt.Sprintf(
		"docker build -t %s/%s:%s .",
		this_settings.UserRegisterContainer,
		this_settings.AppName,
		this_settings.Version,
	)
	fmt.Println(build_image)
	out, err := exec.Command("/bin/sh", "-c", build_image).Output()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("%s", out)
}

// REGISTER CONTAINER
func (this_settings *SettingsFile) RegisterContainer() {
	fmt.Println("\n[ * ] Execute register container")
	cmd_push_image := fmt.Sprintf("docker push %s/%s:%s",
		this_settings.UserRegisterContainer,
		this_settings.AppName,
		this_settings.Version,
	)
	registerBASH, err := exec.Command("/bin/bash", "-c", cmd_push_image).Output()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("%s", registerBASH)
}

// DEPLOY
func (this_settings *SettingsFile) DeployToApp() {
	fmt.Println("\n[ * ] Execute deploy to nomad")
	cmd_deploy := fmt.Sprintf(
		"nomad job run -var image=%s/%s:%s %s",
		this_settings.UserRegisterContainer,
		this_settings.AppName,
		this_settings.Version,
		this_settings.Hcl,
	)
	deployBASH, err := exec.Command("/bin/bash", "-c", cmd_deploy).Output()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("%s", deployBASH)
}

func main() {
	settings := SettingsFile{}
	settings.GetSettings()
	settings.Validate()
	settings.BuildApp()
	settings.RegisterContainer()
	settings.DeployToApp()
}
