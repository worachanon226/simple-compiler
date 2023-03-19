package controller

import (
	"bytes"
	"fmt"
	"os/exec"
	"path"
	"strings"
)

func ExecutePy(path string) string {

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("python", path)
	fmt.Println(cmd)

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Sprint(err) + ": " + stderr.String()
	} else {
		return out.String()
	}
}

func ExecuteCpp(file string) string {
	jobID := strings.Split(path.Base(file), ".")[0]
	outPath := path.Join("components", "cpp", "output", jobID)
	output := fmt.Sprintf("components/" + "cpp/" + "output/" + jobID + ".exe")
	cmd := exec.Command("g++", file, "-o", outPath)

	build := exec.Command(output)

	var out bytes.Buffer
	var stderr bytes.Buffer
	var build_out bytes.Buffer
	var build_stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	build.Stdout = &build_out
	build.Stderr = &build_stderr
	build_err := build.Run()

	if err != nil || build_err != nil {
		return fmt.Sprint(err) + ": " + stderr.String()
	} else {
		return build_out.String()
	}
}
