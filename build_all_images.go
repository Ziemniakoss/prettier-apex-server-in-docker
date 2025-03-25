package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const PRETTIER_PACKAGE_NAME = "prettier-plugin-apex"
const FILE_WITH_BUILT_VERSIONS = "built_versions.txt"
const IMAGE_NAME = "ziemniakoss/prettier-apex-server"

func main() {
	fmt.Println(">>> Reading package versions")
	avalibleVersions := getPrettierVersionNumbers()
	buildResultsChannel := make(chan buildResult)
	for _, version := range avalibleVersions {
		go buildImage(version, buildResultsChannel)
	}

	wasSuccess := true
	for i := 0; i < len(avalibleVersions); i++ {
		result := <-buildResultsChannel
		message := "build successfully"
		if !result.isSuccess {
			wasSuccess = false
			message = "build failed:"
			println("Failed!")
			println(result.output)
		}
		println("[", i+1, "|", len(avalibleVersions), "]", result.version, message)
		wasSuccess = wasSuccess && result.isSuccess
	}
	if !wasSuccess {
		fmt.Println("Failed")
		os.Exit(1)
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
		println("=====\n" + outAsStr + "\n=====")
		println(err)
		panic("Error while deconding json happened")
	}
	return versions
}

func buildImage(pacakgeVersion string, c chan buildResult) {
	command := exec.Command(
		"docker",
		"build",
		"--build-arg",
		"PLUGIN_VERSION="+pacakgeVersion,
		"--tag",
		IMAGE_NAME+":"+pacakgeVersion,
		"--no-cache",
		".",
	)
	command.Wait()
	out, err := command.CombinedOutput()
	c <- buildResult{
		version:   pacakgeVersion,
		output:    string(out),
		isSuccess: err == nil,
	}
}

type buildResult struct {
	version   string
	output    string
	isSuccess bool
}
