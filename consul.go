//+build mage

package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/gruntwork-io/terratest/modules/logger"

	//http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	consul "github.com/hashicorp/consul/api"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Consul3 runs 3 Consul Docker containers and exercises them
func Consul3() error {
	return RunGoTests([]testing.InternalTest{
		{Name: "TestConsul3", F: TestConsul3},
	})
}

func TestConsul3(t *testing.T) {
	type Unit struct {
		docker_id      string
		consul_address string
		join_address   string
		docker_info    *docker.ContainerInspect
		consul         *consul.Client
	}

	var units []*Unit

	t.Run("Launch 3 Consul agents, each as their own server", func(t *testing.T) {
		opts := &docker.RunOptions{
			EnvironmentVariables: []string{`CONSUL_LOCAL_CONFIG={"skip_leave_on_interrupt": true}`},
			Command:              []string{"agent", "-server", "-bootstrap-expect=1", "-client=0.0.0.0"},
			Detach:               true,
			Remove:               false,
			OtherOptions:         []string{"--publish=8500"}, // publish all ports to host
		}

		units = []*Unit{
			{docker_id: docker.RunAndGetID(t, "consul:1.2.3", opts)},
			{docker_id: docker.RunAndGetID(t, "consul:1.2.3", opts)},
			{docker_id: docker.RunAndGetID(t, "consul:1.2.3", opts)},
		}

		for _, u := range units {
			u.docker_info = docker.Inspect(t, u.docker_id)
			u.consul_address = fmt.Sprintf("%s:%d", "localhost", u.docker_info.Ports[0].HostPort)
			u.join_address = "" // TODO: populate from internal docker IP

			logger.Log(t, u.consul_address)
			c, err := consul.NewClient(&consul.Config{Address: u.consul_address})
			u.consul = c

			assert.NoError(t, err)
		}
		logger.Log(t, "Setup complete", units)
	})

	defer docker.Stop(t, []string{units[0].docker_id, units[1].docker_id, units[2].docker_id}, &docker.StopOptions{})

	t.Run("Write a different value to each of their KV stores", func(t *testing.T) {
		for i, u := range units {
			assert.Eventually(t, func() bool {
				val := fmt.Sprintf("unit%d", i)
				logger.Logf(t, "Consul PUT hello = %s", val)
				_, err := u.consul.KV().Put(&consul.KVPair{Key: "hello", Value: []byte(val)}, &consul.WriteOptions{})
				return assert.NoError(t, err)
			}, 20*time.Second, 100*time.Millisecond)
		}
	})

	t.Run("Read back the KV values", func(t *testing.T) {
		for i, u := range units {
			val := fmt.Sprintf("unit%d", i)
			pair, _, err := u.consul.KV().Get("hello", nil)
			logger.Logf(t, "Consul GET hello = %s", pair.Value)
			assert.NoError(t, err)
			assert.Equal(t, val, string(pair.Value))
		}
	})

	t.Run("Join them to form a 3-node cluster", func(t *testing.T) {
		err := units[1].consul.Agent().Join(units[0].consul_address, false)
		assert.NoError(t, err)

		err = units[2].consul.Agent().Join(units[0].consul_address, false)
		assert.NoError(t, err)

		members, err := units[0].consul.Agent().Members(false)
		assert.NoError(t, err)
		require.Len(t, members, 3)

		logger.Log(t, members[0].Name, members[0].Addr)
		logger.Log(t, members[1].Name, members[1].Addr)
		logger.Log(t, members[2].Name, members[2].Addr)
	})

	t.Run("Read KV values from the 3 nodes with and without stale", func(t *testing.T) {

	})
}
