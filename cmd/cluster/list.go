/*
Copyright © 2024 Gustavo V. Goulart <gvillela7@gmail.com>
*/
package cluster

import (
	"encoding/json"
	"fmt"
	"github.com/alexeyco/simpletable"
	"github.com/gvillela7/snake-docker-cli/cmd/cluster/models"
	"github.com/gvillela7/snake-docker-cli/cmd/cluster/service"
	"github.com/gvillela7/snake-docker-cli/cmd/internal"
	"github.com/spf13/cobra"
	"log"
)

// listCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all nodes",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		list := service.DockerNodes()
		var nodes []models.Nodes
		if err := json.Unmarshal(list, &nodes); err != nil {
			log.Fatal(err)
		}
		table := simpletable.New()
		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: internal.Blue("Id")},
				{Align: simpletable.AlignCenter, Text: internal.Blue("Hostname")},
				{Align: simpletable.AlignCenter, Text: internal.Blue("Availability")},
				{Align: simpletable.AlignCenter, Text: internal.Blue("State")},
				{Align: simpletable.AlignCenter, Text: internal.Blue("Status")},
				{Align: simpletable.AlignCenter, Text: internal.Blue("Addr")},
			},
		}
		var cells [][]*simpletable.Cell
		for _, node := range nodes {
			cells = append(cells, *&[]*simpletable.Cell{
				{Text: fmt.Sprintf("%s", node.Id[0:8])},
				{Text: fmt.Sprintf("%s", node.Hostname.Hostname)},
				{Text: fmt.Sprintf("%s", node.Spec.Availability)},
				{Text: internal.Green(fmt.Sprintf("\u2705 %s", node.State))},
				{Text: fmt.Sprintf("%s", node.Status.State)},
				{Text: fmt.Sprintf("%s", node.Status.Addr)},
			})
		}
		table.Body = &simpletable.Body{Cells: cells}
		table.SetStyle(simpletable.StyleRounded)
		table.Println()
	},
}

func init() {

}
