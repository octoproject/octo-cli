package service

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/octoproject/octo-cli/config"
)

const (
	imageDefaultTag = "latest"
)

func BuildFunction(s *config.Service, registryPrefix, imgTag string) error {
	serviceDir := s.ServiceName
	// check if service dir exists
	if !exists(serviceDir) {
		return fmt.Errorf("folder %s is not exists", serviceDir)
	}

	imageTag := imgTag
	//if tag is not provided, default tag will be used
	if len(imageTag) == 0 {
		imageTag = imageDefaultTag
	}

	imageName := fmt.Sprintf("%s/%s:%s", registryPrefix, s.ServiceName, imageTag)
	fmt.Println(imageName, serviceDir)
	cmd := exec.Command("docker", "build", "-t", imageName, serviceDir)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
