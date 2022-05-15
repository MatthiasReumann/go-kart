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
	buildContext, err := os.Open(dockerfile)
	if err != nil {
		return err
	}
	buildOptions := types.ImageBuildOptions{Dockerfile: "Dockerfile", Tags: []string{"mini"}}
	_, err = c.cli.ImageBuild(context.Background(), buildContext, buildOptions)
	return err
}

func (c dockerClient) Run() {

}

func (c dockerClient) Remove() {

}

func (c dockerClient) RemoveImage() {

}
