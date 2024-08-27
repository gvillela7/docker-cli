/*
Copyright © 2024 Gustavo V. Goulart <gvillela7@gmail.com>
*/
package services

import (
	"encoding/json"
	"fmt"
	"github.com/gvillela7/snake-docker-cli/cmd/cluster/models"
	"github.com/gvillela7/snake-docker-cli/cmd/cluster/service"
	"github.com/spf13/cobra"
	"log"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all nodes",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		list := service.DockerServices()
		var services []models.Services
		if err := json.Unmarshal(list, &services); err != nil {
			log.Fatal(err)
		}
		for _, serviceDocker := range services {
			fmt.Println(serviceDocker.Id, serviceDocker.Spec.Name, serviceDocker.Spec.Mode.Replicated, serviceDocker.Spec.Name, serviceDocker.Ports.Ports.TargetPort)
		}
	},
}

func init() {

}
