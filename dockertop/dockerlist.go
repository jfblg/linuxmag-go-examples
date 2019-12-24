package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func dockerList(error, []string, map[string]types.Container) {
	item := []string{}
	containerMap := make(map[string]types.Continer)

	opt := client.WithAPIVersionNegotiation()
	cli, err := client.NewClientWithOpts(opt)
	if err != nil {
		return err, nil, nil
	}
	defer cli.Close()

	containers, err := cli.ContainerList(
		context.Background(),
		types.ContainerListOptions{}
	)
	if err != nil {
		return err, nil, nil
	}

	for _, container := range containers {
		name := fmt.Sprintf("%s-%s", container.Image, container.ID[:10])
		itmes = append(items, name)
		containerMap[name] = container
	}
	return nil, items, containerMap
}