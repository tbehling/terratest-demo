//+build mage

package main

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
)

func withDockerContainer(t *testing.T, tag string, f func(*docker.ContainerInspect)) {
	opts := &docker.RunOptions{Detach: true, Remove: true, OtherOptions: []string{"-P"}}
	docker_id := docker.RunAndGetID(t, tag, opts)
	defer docker.Stop(t, []string{docker_id}, &docker.StopOptions{})
	container_info := docker.Inspect(t, docker_id)

	f(container_info)
}
