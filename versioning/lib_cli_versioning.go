package versioning

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const config_path = "./lib/cli/versioning.json"

type Versioning struct {
	AppName       string `json:"appName"`
	Version       string `json:"version"`
	ContainerName string `json:"containerName"`
	SH_Test       string `json:"shTest"`
}

func (this_version *Versioning) GetVersion() {
	jsonFile, err := os.ReadFile(config_path)
	if err != nil {
		log.Fatal(err)
	}
	var api Versioning
	json.Unmarshal(jsonFile, &api)
	this_version.AppName = api.AppName
	this_version.Version = api.Version
	this_version.ContainerName = api.ContainerName
	this_version.SH_Test = api.SH_Test
	this_version.validateVersionSoftware()
	fmt.Println("FROM VERSION:", api.Version)
	fmt.Println("TO VERSION:", this_version.Version)

}

func (this_version *Versioning) validateVersionSoftware() {
	var version string
	separateVersion := strings.Split(this_version.Version, ".")
	v1, err := strconv.Atoi(separateVersion[0])
	if err != nil {
		panic(err)
	}
	v2, err := strconv.Atoi(separateVersion[1])
	if err != nil {
		panic(err)
	}
	v3, err := strconv.Atoi(separateVersion[2])
	if err != nil {
		panic(err)
	}
	v4, err := strconv.Atoi(separateVersion[3])
	if err != nil {
		panic(err)
	}
	if v4 < 9 {
		separateVersion[3] = fmt.Sprintf("%d", v4+1)
		version = fmt.Sprintf("%s.%s.%s.%s", separateVersion[0], separateVersion[1], separateVersion[2], separateVersion[3])
		this_version.Version = version
		return
	}
	if v3 < 9 {
		separateVersion[2] = fmt.Sprintf("%d", v3+1)
		version = fmt.Sprintf("%s.%s.%s.%s", separateVersion[0], separateVersion[1], separateVersion[2], "0")
		this_version.Version = version
		return
	}
	if v2 < 9 {
		separateVersion[1] = fmt.Sprintf("%d", v2+1)
		version = fmt.Sprintf("%s.%s.%s.%s", separateVersion[0], separateVersion[1], "0", "0")
		this_version.Version = version
		return
	}
	if v1 < 9 {
		separateVersion[0] = fmt.Sprintf("%d", v1+1)
		version = fmt.Sprintf("%s.%s.%s.%s", separateVersion[0], "0", "0", "0")
		this_version.Version = version
		return
	}
	if v1 == 9 && v2 == 9 && v3 == 9 && v4 == 9 {
		log.Fatal("CONGRATULATIONS MAX VERSION IS CONFIGURED")
		return
	}
}

func (this_version *Versioning) SetVersionSoftware() {
	file, _ := json.MarshalIndent(this_version, "", " ")
	err := os.WriteFile(config_path, file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
