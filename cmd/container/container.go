/*
Copyright © 2024 Gustavo V. Goulart <gvillela7@gmail.com>
*/
package container

import (
	"github.com/spf13/cobra"
)

var (
	list bool
	show string
)

// containerCmd represents the container command
var ContainerCmd = &cobra.Command{
	Use:   "container",
	Short: "List containers, view containers",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if list {
			ListCmd.Run(cmd, args)
		}
		if show != "" {
			ShowContainerByID(show)
		}
	},
}

func init() {
	ContainerCmd.Flags().BoolVar(&list, "list", false, "List all containers")
	ContainerCmd.Flags().StringVarP(&show, "show", "s", "", "Show specific container id")
}
