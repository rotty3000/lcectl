/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package runtime

import (
	"fmt"
	"io"

	"github.com/docker/docker/api/types/container"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"liferay.com/lcectl/constants"
	lcectldocker "liferay.com/lcectl/docker"
	"liferay.com/lcectl/prereq"
	"liferay.com/lcectl/spinner"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates the runtime environment for Liferay Client Extension development",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		prereq.Prereq(Verbose)

		config := container.Config{
			Image: "localdev-server",
			Cmd:   []string{"/repo/scripts/runtime/create.sh"},
		}
		host := container.HostConfig{
			Binds: []string{
				fmt.Sprintf("%s:%s", viper.GetString(constants.Const.RepoDir), "/repo"),
				"/var/run/docker.sock:/var/run/docker.sock",
			},
			NetworkMode: container.NetworkMode(viper.GetString(constants.Const.DockerNetwork)),
		}

		spinner.Spin(
			"Creating", "Created", Verbose,
			func(fior func(io.ReadCloser, bool)) int {
				return lcectldocker.InvokeCommandInLocaldev("localdev-create", config, host, true, Verbose, fior)
			})
	},
}

func init() {
	createCmd.Flags().BoolVarP(&Verbose, "verbose", "v", false, "enable verbose output")
	runtimeCmd.AddCommand(createCmd)
}