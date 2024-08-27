/*
Copyright © 2024 Gustavo V. Goulart <gvillela7@gmail.com>
*/
package services

import (
	"fmt"
	"github.com/spf13/cobra"
)

var list bool
var show string
var ServicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Show all services",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if list {
			ListCmd.Run(cmd, args)
		}
		if show != "" {
			fmt.Println("Show service by Id")
		}
	},
}

func init() {
	ServicesCmd.Flags().BoolVar(&list, "list", false, "List all services")
	ServicesCmd.Flags().StringVarP(&show, "show", "s", "", "Show specific service")
}
