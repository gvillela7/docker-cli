/*
Copyright © 2024 Gustavo V. Goulart <gvillela7@gmail.com>
*/
package cluster

import (
	"fmt"
	"github.com/spf13/cobra"
)

var list bool

// containerCmd represents the container command
var ClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "List nodes for a cluster",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if list {
			ListCmd.Run(cmd, args)
		}
	},
}

func init() {
	ClusterCmd.Flags().BoolVar(&list, "list", false, "List nodes")
	if err := ClusterCmd.MarkFlagRequired("list"); err != nil {
		fmt.Println("usage: --list")
	}
}
