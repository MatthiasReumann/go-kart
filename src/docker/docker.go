package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"os"
)

type DockerClient interface {
	Build(string) error
	Run()
	Remove()
	RemoveImage()
}

type dockerClient struct {
	cli *client.Client
}

func NewDockerClient() (DockerClient, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return dockerClient{cli}, nil
}

func (c dockerClient) Build(dockerfile string) error {
	ctx := context.Background()
	f, err := os.Open(dockerfile)
	if err != nil {
		return err
	}
	_, err = c.cli.ImageBuild(ctx, f, types.ImageBuildOptions{})
	return err
}

func (c dockerClient) Run() {

}

func (c dockerClient) Remove() {

}

func (c dockerClient) RemoveImage() {

}
