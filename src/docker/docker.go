package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"io"
	"log"
	"os"
)

type DockerClient interface {
	Build(string, []string) error
	Run(string) error
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

func (c dockerClient) Build(dockerfile string, tags []string) error {
	buildContext, err := os.Open(dockerfile)
	if err != nil {
		return err
	}
	defer buildContext.Close()

	buildOptions := types.ImageBuildOptions{
		Dockerfile: "Dockerfile",
		Context:    buildContext,
		Tags:       tags,
		NoCache:    true,
		Remove:     true,
	}
	buildResponse, err := c.cli.ImageBuild(context.Background(), buildContext, buildOptions)
	if err != nil {
		return err
	}
	defer buildResponse.Body.Close()

	_, err = io.Copy(io.Discard, buildResponse.Body)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (c dockerClient) Run(imageName string) error {
	ctx := context.Background()
	resp, err := c.cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
	}, nil, nil, nil, "")
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	if err := c.cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	statusCh, errCh := c.cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-statusCh:
	}

	out, err := c.cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return err
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	return nil
}

func (c dockerClient) Remove() {

}

func (c dockerClient) RemoveImage() {

}
