/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// refreshCmd represents the refresh command
var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Refreshes client-extension workload resources in localdev server",
	Run: func(cmd *cobra.Command, args []string) {
		dockerClient := InitDocker()
		runLocaldevRefresh("localdev-server", dockerClient)
	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// refreshCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// refreshCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runLocaldevRefresh(imageTag string, dockerClient *client.Client) error {
	ctx := context.Background()

	networkConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{},
	}
	networkConfig.EndpointsConfig[viper.GetString(Const.dockerNetwork)] =
		&network.EndpointSettings{}

	resp, err := dockerClient.ContainerCreate(
		ctx,
		&container.Config{
			Image:        imageTag,
			Cmd:          []string{"tilt", "trigger", "(Tiltfile)", "--host", "host.docker.internal"},
			AttachStdout: true,
			AttachStderr: true,
			Tty:          true,
		},
		&container.HostConfig{
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: "/var/run/docker.sock",
					Target: "/var/run/docker.sock",
				},
			},
			AutoRemove: true,
		},
		networkConfig,
		nil,
		"localdev-server-refresh")

	if err != nil {
		log.Fatalf("Failed to create container %s: %s", imageTag, err)
		return err
	}

	err = dockerClient.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})

	if err != nil {
		log.Fatalf("Failed to start container %s: %s", imageTag, err)
	}

	hijacked, err := dockerClient.ContainerAttach(ctx, resp.ID, types.ContainerAttachOptions{
		Stderr: true,
		Stdout: true,
		Stream: true,
	})

	if err != nil {
		log.Fatalf("Failed to attach to container %s", resp.ID)
	}

	go io.Copy(os.Stdout, hijacked.Reader)
	go io.Copy(os.Stderr, hijacked.Reader)

	statusCh, errCh := dockerClient.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}
	return nil
}
