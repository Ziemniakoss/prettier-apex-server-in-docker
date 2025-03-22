package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

const PRETTIER_PACKAGE_NAME = "prettier-plugin-apex"
const FILE_WITH_BUILT_VERSIONS = "built_versions.txt"
const START_VERSION = "1.0.0"
const IMAGE_NAME = "ziemniakoss/prettier-apex-server"

func main() {
	fmt.Println(">>> Reading package versions")
	avalibleVersions := getPrettierVersionNumbers()
	alreadyBuiltVersions := getBuiltVersions()
	for _, version := range avalibleVersions {
		if !shouldBeBuilt(version, alreadyBuiltVersions) {
			continue
		}
		println(">>> Building image for package", version)
		buildImage(version)
	}
}

func getPrettierVersionNumbers() []string {
	command := exec.Command("npm", "view", PRETTIER_PACKAGE_NAME, "versions")
	command.Wait()
	out, err := command.Output()
	if err != nil {
		panic("Error listing npm package versions")

	}

	outAsStr := string(out)
	outAsStr = strings.ReplaceAll(outAsStr, "'", "\"") // Yeah, output is not parsable JSON

	var versions []string
	err = json.Unmarshal([]byte(outAsStr), &versions)
	if err != nil {
		fmt.Println("=====\n" + outAsStr + "\n=====")
		fmt.Println(err)
		panic("Error while deconding json happened")
	}
	return versions
}

func getBuiltVersions() []string {
	// TODO
	return []string{}
}

func buildImage(pacakgeVersion string) {
	command := exec.Command(
		"docker",
		"build",
		"--build-arg="+pacakgeVersion,
		"--tag",
		IMAGE_NAME+":"+pacakgeVersion,
		".",
	)
	command.Wait()
	out, err := command.CombinedOutput()
	if err != nil {
		println(out)
		panic(err)
	}
}

func shouldBeBuilt(version string, builtVersions []string) bool {
	if START_VERSION > version {
		return false
	}
	for _, v := range builtVersions {
		if version == v {
			return false
		}
	}
	return true
}
