package controller

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

func ExecutePy(path string) string {

	var out bytes.Buffer
	var stderr bytes.Buffer

	dir, _ := os.Getwd()

	cmd := exec.Command("python3", dir+path[1:])
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
	cmd := exec.Command("g++", file, "-o", outPath)

	output := fmt.Sprintf("components/" + "cpp/" + "output/" + jobID)

	var out bytes.Buffer
	var stderr bytes.Buffer
	var build_out bytes.Buffer
	var build_stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	build := exec.Command("./" + output)
	build.Stdout = &build_out
	build.Stderr = &build_stderr
	build_err := build.Run()

	if err != nil || build_err != nil {
		return fmt.Sprint(err) + ": " + stderr.String()
	} else {
		return build_out.String()
	}
}
