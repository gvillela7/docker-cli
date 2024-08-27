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
	"log"
)

func ShowContainerByID(id string) {
	inspect := service.GetContainerByID(id)
	var containerInspect models.Inspect
	if err := json.Unmarshal(inspect, &containerInspect); err != nil {
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
			{Align: simpletable.AlignCenter, Text: internal.Blue("Memory")},
			{Align: simpletable.AlignCenter, Text: internal.Blue("ShmSize")},
			{Align: simpletable.AlignCenter, Text: internal.Blue("CpuPeriod")},
			{Align: simpletable.AlignCenter, Text: internal.Blue("CpuQuota")},
			{Align: simpletable.AlignCenter, Text: internal.Blue("CpuPercent")},
			{Align: simpletable.AlignCenter, Text: internal.Blue("CpuShares")},
		},
	}
	var cells [][]*simpletable.Cell
	cells = append(cells, *&[]*simpletable.Cell{
		{Text: fmt.Sprintf("%s", containerInspect.Id[0:8])},
		{Text: fmt.Sprintf("%s", containerInspect.Config.Image)},
		{Text: fmt.Sprintf("\u2705 %v", containerInspect.State.Running)},
		{Text: internal.Green(fmt.Sprintf("\u2705 %s", containerInspect.State.Status))},
		{Text: fmt.Sprintf("%s", containerInspect.NetworkSettings.Networks.Bridge.IPAddress)},
		{Text: fmt.Sprintf("%d", containerInspect.HostConfig.Memory)},
		{Text: fmt.Sprintf("%d", containerInspect.HostConfig.ShmSize)},
		{Text: fmt.Sprintf("%d", containerInspect.HostConfig.CPUPeriod)},
		{Text: fmt.Sprintf("%d", containerInspect.HostConfig.CPUQuota)},
		{Text: fmt.Sprintf("%d", containerInspect.HostConfig.CPUPercent)},
		{Text: fmt.Sprintf("%d", containerInspect.HostConfig.CPUShares)},
	})
	table.Body = &simpletable.Body{Cells: cells}
	table.SetStyle(simpletable.StyleRounded)
	table.Println()
}
