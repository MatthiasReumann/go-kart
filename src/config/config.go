package config

type Endpoint struct {
	Path          string
	DockerfileTar string
	Tags          []string
}

func NewEndpoint(path, dockerfileTar string, tags []string) Endpoint {
	return Endpoint{
		Path:          path,
		DockerfileTar: dockerfileTar,
		Tags:          tags,
	}
}
