/*
Copyright © 2024 Gustavo V. Goulart <gvillela7@gmail.com>
*/
package container

import (
	"encoding/json"
	"fmt"
	"github.com/alexeyco/simpletable"
	"github.com/gvillela7/snake-docker-cli/cmd/container/models"
	"github.com/gvillela7/snake-docker-cli/cmd/container/service"
	"github.com/gvillela7/snake-docker-cli/cmd/internal"
	"github.com/spf13/cobra"
	"log"
)

// listCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all containers",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		list := service.ContainersList()
		var containers []models.Container
		if err := json.Unmarshal(list, &containers); err != nil {
			log.Fatal(err)
		}
		table := simpletable.New()
		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: internal.Blue("Id")},
				{Align: simpletable.AlignCenter, Text: internal.Blue("Image")},
				{Align: simpletable.AlignCenter, Text: internal.Blue("State")},
				{Align: simpletable.AlignCenter, Text: internal.Blue("Status")},
				{Align: simpletable.AlignCenter, Text: internal.Blue("IPAddress")},
			},
		}
		var cells [][]*simpletable.Cell
		for _, container := range containers {
			cells = append(cells, *&[]*simpletable.Cell{
				{Text: fmt.Sprintf("%s", container.Id[0:8])},
				{Text: fmt.Sprintf("%s", container.Image)},
				{Text: fmt.Sprintf("\u2705 %s", container.State)},
				{Text: internal.Green(fmt.Sprintf("\u2705 %s", container.Status))},
				{Text: fmt.Sprintf("%s", container.NetworkSettings.Networks.Bridge.IPAddress)},
			})
		}
		table.Body = &simpletable.Body{Cells: cells}
		table.SetStyle(simpletable.StyleRounded)
		table.Println()
	},
}

func init() {

}
